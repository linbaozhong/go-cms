package controllers

import (
	"cms/models"
	"cms/utils"
)

type Profile struct {
	Admin
}

//首页
func (this *Profile) Index() {
	this.TplNames = this.getTplFileName("index")
	this.Render()
}

//修改密码
func (this *Profile) UpdatePassword() {

	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()
		this.TplNames = this.getTplFileName("password")
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
	var data interface{}

	pass := new(models.Password)
	pass.OldPassword, pass.NewPassword, pass.RePassword = this.GetString("oldpassword"), this.GetString("newpassword"), this.GetString("repassword")

	//数据合法性检验
	if data, inv := this.invalidModel(pass); inv {
		this.renderJson(data)
		return
	}

	//两次输入的新密码是否一致
	if pass.NewPassword != pass.RePassword {
		//返回错误
		data = utils.JsonMessage(false, "inconsistent", this.lang("inconsistent"))
		return
	}
	//变更密码
	err := users.UpdatePassword(this.xm, pass)

	if err == nil {
		data = utils.JsonMessage(true, "success", this.lang("success"))
	} else {
		data = utils.JsonMessage(false, "updatePasswordFail", this.lang("updatePasswordFail"))
	}
	this.renderJson(data)
}

//修改账户信息
func (this *Profile) UpdateProfile() {

	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()
		this.Data["user"], _ = users.Get(this.xm.Updator)
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

	u.Id = this.xm.Updator
	//u.Loginname = this.GetString("loginname")
	u.Relname = this.GetString("relname")

	//数据合法性检验
	if data, inv := this.invalidModel(u); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	var data interface{}
	_, err := users.UpdateProfile(u)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonMessage(true, "", "")
	}
	this.renderJson(data)
}

func (this *Profile) getTplFileName(s string) string {
	return getTplFileName("admin/profile", s)
}
