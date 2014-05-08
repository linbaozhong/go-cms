package controllers

import (
	"bytes"
	"cms/models"
	"cms/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	//"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type base struct {
	beego.Controller
	page      models.Page
	xm        *models.Field //用于在controller和model之间传递数据的小蜜
	methodGet bool          //是否GET请求
}

//
func (this *base) init() {
	this.methodGet = strings.ToUpper(this.Ctx.Request.Method) == "GET"
	this.initPage()
}

//终止服务
func (this *base) end() {
	this.Layout = ""
	this.TplNames = ""

	this.StopRun()
}

//错误页处理
func (this *base) errorHandle(msg ...interface{}) {
	n := len(msg)

	if n > 0 {
		this.Data["msg"] = msg[0]
	}
	this.TplNames = getTplFileName("", "error")
	this.Render()
	this.end()
}

//获取当前语言
func (this *base) lang(k string) string {
	return utils.Lang(k, this.Ctx.Request.Header.Get("Accept-Language"))
}

//获取URL参数
func (this *base) getParamsInt64(key string) (int64, error) {
	i64, err := utils.Str2int64(this.getParamsString(key)) //strconv.ParseInt(this.getParamsString(key), 10, 0)
	return i64, err
}

func (this *base) getParamsInt(key string) (int, error) {
	i64, err := this.getParamsInt64(key)
	return int(i64), err
}

func (this *base) getParamsString(key string) string {
	//fmt.Println(this.Ctx.Params[fmt.Sprintf("%v", key)])
	return this.Ctx.Input.Param(key)
}

//验证合法用户
func (this *base) validUser() (*models.Current, bool) {
	coo := this.Ctx.GetCookie(beego.AppConfig.String("CookieName"))

	if coo == "" {
		return nil, false
	}

	currentuser := models.GetCurrentUser(coo)
	if currentuser.Id == 0 {
		return nil, false
	}
	return currentuser, true
}

//允许新的请求，数据通用字段初始信息，附带验证用户是否合法(err)，
func (this *base) allowRequest() (ok bool) {
	currentuser, ok := this.validUser()
	field := new(models.Field)

	if ok {
		field.Status = 0
		field.Deleted = 0
		field.Updator = currentuser.Id
		field.Updated = utils.Millisecond(time.Now())
		field.Ip = utils.GetIp(this.Ctx.Request.RemoteAddr)
		field.Role = currentuser.Role
		field.Name = currentuser.Name
	}
	this.xm = field

	return
}

//是否外链
func (this *base) isOutLink() bool {
	host, err := url.Parse(this.Ctx.Request.Referer())
	if err != nil {
		return true
	}
	return this.Ctx.Request.Host != host.Host
}

//文件服务
func (this *base) serverFile(file string) {
	file = filepath.Join(".", file)
	this.Ctx.ResponseWriter.Header().Set("Content-Description", "File Transfer")
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/octet-stream;charset=UTF-8")
	this.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment; filename="+utils.UrlEncode(filepath.Base(file)))
	this.Ctx.ResponseWriter.Header().Set("Content-Transfer-Encoding", "binary")
	this.Ctx.ResponseWriter.Header().Set("Expires", "0")
	this.Ctx.ResponseWriter.Header().Set("Cache-Control", "must-revalidate")
	this.Ctx.ResponseWriter.Header().Set("Pragma", "public")

	http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, file)
	this.end()
}

//设置签名，防止重复提交
func (this *base) token() string {
	s := utils.MD5(time.Now().String())
	this.SetSession("token", string(s))

	return s
}

//验证签名是否无效
func (this *base) invalidToken() bool {
	sess := this.GetSession("token")
	if sess == nil {
		return true
	}
	if token := this.GetString("token"); token != sess.(string) {
		return true
	}
	return false
}

//响应签名丢失错误
func (this *base) renderLoseToken() {
	data := utils.JsonMessage(false, "invalidFormToken", this.lang("invalidFormToken"))
	this.renderJson(data)
}

//删除签名
func (this *base) deleteToken() {
	this.DelSession("token")
}

//页面公共信息
func (this *base) initPage() {
	this.page.SiteName = this.conf("sitename")
	this.page.Title = this.conf("title")
	this.page.Company = this.conf("company")
	this.page.Copyright = this.conf("copyright")
	this.page.Domain = this.conf("domain")
	this.page.Keywords = this.conf("keywords")
	this.page.Description = this.conf("description")
	this.page.Author = this.conf("author")
}

//读取配置
func (this *base) conf(key string) string {
	return beego.AppConfig.String(key)
}

//
func (this *base) setJsonData(data interface{}) {
	//操作成功，清除token
	if resp := reflect.Indirect(reflect.ValueOf(data)); resp.FieldByName("Ok").Bool() {
		this.deleteToken()
	}

	this.Data["json"] = data
}

//返回json响应格式
func (this *base) renderJson(data interface{}) {
	this.setJsonData(data)
	this.ServeJson()
}

//返回jsonp响应
func (this *base) renderJsonp(data interface{}) {
	this.setJsonData(data)
	this.ServeJsonp()
}

//渲染html响应
func (this *base) renderHtml(args ...interface{}) {

}

/*
数据对象合法性验证
args：要检验的元素，长度为0，则检验全部元素
*/
func (this *base) invalidModel(m interface{}, args ...interface{}) (data interface{}, invalid bool) {

	valid := validation.Validation{}

	b, err := valid.Valid(m)
	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
		invalid = true
		return
	}
	if !b {
		var errstr string
		//检验元素
		if n := len(args); n > 0 {
			for _, err := range valid.Errors {
				if utils.ListContains(args, err.Key[0:strings.Index(err.Key, ".")]) {
					errstr += fmt.Sprintf("%s %s;", err.Key, err.Message)
				}
			}
		} else {
			for _, err := range valid.Errors {
				errstr += fmt.Sprintf("%s %s;", err.Key, err.Message)
			}
		}

		if errstr == "" {
			invalid = false
		} else {
			data = utils.JsonMessage(false, "", errstr)
			invalid = true
		}
		return
	}
	return
}

//文件上传
func (this *base) upload(key string) (files []*models.UploadFile, err error) {
	//处理上传文件
	var header *multipart.FileHeader
	var file multipart.File
	var f *os.File

	//根据年月选择文件夹
	t := time.Now().Format(time.RFC3339)
	//文件夹是否存在或创建文件夹
	UploadPath := beego.AppConfig.String("UploadPath")
	folder := utils.MergePath(UploadPath)
	err = utils.GetDir(folder)
	if err != nil {
		return
	}
	// //用户文件夹是否存在或创建文件夹
	// folder = filepath.Join(folder, strconv.Itoa(int(this.xm.Updator)))
	// err = utils.GetDir(folder)
	// if err != nil {
	// 	return
	// }
	//文件夹是否存在或创建文件夹
	UploadPath = path.Join(UploadPath, beego.Substr(t, 0, 7))
	folder = path.Join(folder, beego.Substr(t, 0, 7))
	err = utils.GetDir(folder)
	if err != nil {
		return
	}

	fs := this.Ctx.Request.MultipartForm.File[key]

	n := len(fs)
	if n == 0 {
		err = errors.New("files not found")
		return
	}

	for i := 0; i < n; i++ {
		header = fs[i]
		file, err = fs[i].Open()

		if err != nil {
			return
		}

		defer file.Close()

		//提取原始文件信息
		disposition := strings.Split(header.Header.Get("Content-Disposition"), ";")

		var key, value string
		for _, v := range disposition {

			pos := strings.Index(v, "=")
			if pos > -1 {
				key = v[:pos]

				if strings.TrimSpace(key) == "filename" {
					value = strings.Replace(v[pos+1:], "\"", "", -1)
					break
				}
			}
		}
		//
		filename := filepath.Base(value)

		//新建文件
		UploadPath = path.Join("/", UploadPath, fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(filename)))
		f, err = os.OpenFile(path.Join(folder, fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(filename))), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return
		}

		defer f.Close()

		io.Copy(f, file)

		upf := new(models.UploadFile)
		upf.Name = filename
		upf.Ext = filepath.Ext(filename)
		upf.Path = UploadPath
		fi, _ := f.Stat()
		upf.Size = fi.Size()

		files = append(files, upf)
	}
	return
}

//获取模板
func getTplFileName(p, s string) string {
	if p == "" {
		return fmt.Sprintf("%s.%s", s, beego.AppConfig.String("TemplateFileSuffix"))
	} else {
		return fmt.Sprintf("%s/%s.%s", p, s, beego.AppConfig.String("TemplateFileSuffix"))
	}
}

//
func Navibar(currentChannelEnName string, level int) template.HTML {
	if len(currentChannelEnName) == 0 {
		return ""
	}
	pid := channels.GetParentId(currentChannelEnName)
	navs, err := channels.GetAll(utils.ChNavigation, pid, utils.StatEnabled, level)
	if err != nil {
		return ""
	}

	data := make(map[string]interface{})
	data["navs"] = navs
	data["action"] = currentChannelEnName

	ibytes := bytes.NewBufferString("")
	t, err := template.ParseFiles(fmt.Sprintf("%s/%s", beego.ViewsPath, getTplFileName("home", "navibar")))
	if err == nil {
		t.Execute(ibytes, data)
		icontent, _ := ioutil.ReadAll(ibytes)
		return template.HTML(string(icontent))
	}
	return template.HTML("")
}
