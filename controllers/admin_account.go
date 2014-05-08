package controllers

import (
	"cms/models"
	"cms/utils"
	"fmt"
	"github.com/astaxie/beego"
	//"strconv"
)

type Account struct {
	Admin
}

var users = &models.Users{}

//首页
func (this *Account) Index() {
	us, _ := users.GetAll()
	this.Data["accounts"] = us

	this.TplNames = this.getTplFileName("index")

	this.Render()
}

//新增账户
func (this *Account) Create() {
	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()
		this.Data["password"] = beego.AppConfig.String("DefaultPassword")
		this.TplNames = this.getTplFileName("create")
		this.Render()
		return
	}
	//Post方法
	//签名错误，返回重复提交错误
	if this.invalidToken() {
		this.renderLoseToken()
		return
	}
	//数据模型
	u := new(models.Users)

	models.Extend(u, this.xm)

	u.Loginname = this.GetString("loginname")
	u.Password = this.GetString("password")
	u.Relname = this.GetString("relname")
	r, _ := this.GetInt("role")
	u.Role = int8(r)

	//数据合法性检验
	if data, inv := this.invalidModel(u); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	var data interface{}
	_, err := users.Add(u)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonMessage(true, "", "")
	}
	this.renderJson(data)
}

//修改账户信息
func (this *Account) Edit() {
	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()

		id, err := this.getParamsInt64(":id")
		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams")))
			return
		}
		u, err := users.Get(id)

		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "", err.Error()))
			return
		}
		this.Data["user"] = u

		this.TplNames = this.getTplFileName("edit")
		this.Render()
		return
	}
	//Post方法
	//签名错误，返回重复提交错误
	if this.invalidToken() {
		this.renderLoseToken()
		return
	}
	//数据模型
	u := new(models.Users)

	models.Extend(u, this.xm)

	u.Id, _ = this.GetInt("id")
	u.Loginname = this.GetString("loginname")
	u.Relname = this.GetString("relname")

	//数据合法性检验
	// valid := validation.Validation
	// valid.Required(u.Loginname, "loginname")
	// valid.Required(u.Relname, "relname")

	if data, inv := this.invalidModel(u, "Loginname", "Relname"); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	var data interface{}
	_, err := users.Update(u)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonMessage(true, "", "")
	}
	this.renderJson(data)
}

//账户列表
func (this *Account) GetAll() {
	us, err := users.GetAll()

	var data interface{}
	if err == nil {
		data = utils.JsonResult(true, "", us)
	} else {
		data = utils.JsonMessage(false, "", "")
	}

	this.renderJson(data)
}

/*
是否存在重名用户
*/
func (this *Account) Exist() {
	var data interface{}
	name := this.GetString("loginname")

	if len(name) > 0 {

		if ok := users.Exist(name); ok {
			data = utils.JsonMessage(true, "sameNameAccount", this.lang("sameNameAccount"))
		} else {
			data = utils.JsonMessage(false, "", "")
		}

	} else {
		data = utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams"))
	}
	this.renderJson(data)
}

/*
重置用户
method：post
params：id
*/
func (this *Account) Reset() {
	var data interface{}
	//
	id, err := this.GetInt("id")
	if err != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err.Error())
	} else {
		//不能自我重置
		if id == this.xm.Updator {
			data = utils.JsonMessage(false, "denyOneself", this.lang("denyOneself"))
		} else {
			_, err = users.Reset(id, this.xm)
			if err == nil {
				data = utils.JsonMessage(true, "", "")
			} else {
				data = utils.JsonMessage(false, "resetFail", err.Error())
			}
		}
	}
	this.renderJson(data)
}

/*
删除用户
method：post
params：ids
*/
func (this *Account) Delete() {
	var data interface{}
	id := this.getParamsString(":id")
	//
	ids := this.GetStrings("ids")

	if id != "" {
		ids = append(ids, id)
	}
	//不能自我重置
	if utils.StringsContains(ids, fmt.Sprintf("%d", this.xm.Updator)) {
		data = utils.JsonMessage(false, "denyOneself", this.lang("denyOneself"))
	} else {
		var err error

		_, err = users.Delete(ids, this.xm)

		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "deleteFail", this.lang("deleteFail"))
		}
	}
	this.renderJson(data)
}

//取模板文件
func (this *Account) getTplFileName(s string) string {
	return getTplFileName("admin/account", s)
}
