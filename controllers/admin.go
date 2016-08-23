package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	//"cms/models"
	"strings"

	"github.com/linbaozhong/go-cms/utils"
)

type Admin struct {
	base
}

func (this *Admin) Prepare() {
	this.init()
	this.Layout = this.getTplFileName("layout_admin")

	controller, action := this.Controller.GetControllerAndAction()
	this.Data["page"] = this.page
	this.Data["path"] = strings.ToLower(controller)

	//返回地址
	returnurl := controller
	if controller != "admin" {
		returnurl = fmt.Sprintf("admin/%s", strings.ToLower(action))
	}

	//如果不是合法用户
	if !this.allowRequest() {
		this.Redirect(fmt.Sprintf("%s?returnurl=/%s", beego.AppConfig.String("LoginPath"), returnurl), 302)
	} else {
		this.Data["isSuperAdmin"] = this.isSuperAdmin()
	}
	//当前用户信息
	this.Data["current"] = this.xm
}

//文件上传，配合ueditor.js
func (this *Admin) Upload() {
	var url, state string

	files, err := this.upload("upfile")
	if err == nil {
		state = "SUCCESS"
		url = files[0].Path
	} else {
		state = err.Error()
	}
	this.Ctx.WriteString("<script>parent.UM.getEditor('" + this.GetString("editorid") + "').getWidgetCallback('image')('" + url + "','" + state + "')</script>")
	this.end()
}

func (this *Admin) Index() {
	this.TplName = this.getTplFileName("index")
	this.Render()
}
func (this *Admin) isSuperAdmin() bool {
	return this.xm.Role == fmt.Sprintf("%d", utils.RoleSuper)
}
func (this *Admin) getTplFileName(s string) string {
	return getTplFileName("admin", s)
}
