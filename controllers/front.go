package controllers

import (
	"cms/models"
	"cms/utils"
	"fmt"

	"github.com/astaxie/beego"
)

type Front struct {
	base
}

func (this *Front) Prepare() {
	this.init()
	this.Data["page"] = this.page
	this.Layout = this.getTplFileName("layout")
}

func (this *Front) Finish() {
	if this.TplNames == "" {
		return
	}

	this.Render()
}

//登录页
func (this *Front) Login() {
	loginName := this.GetString("loginname")
	//登录成功后转向地址
	returnUrl := this.GetString("returnurl")
	//Get
	if this.methodGet {
		this.Layout = ""

		if returnUrl == "" {
			returnUrl = this.Ctx.Request.Referer()
		}
		if returnUrl == "" {
			returnUrl = "/"
		}

		if _, ok := this.validUser(); ok {
			this.Redirect(returnUrl, 302)
			return
		}

		//fmt.Println(returnUrl, this.Ctx.Request.Referer())

		this.Data["loginname"] = loginName
		this.Data["returnurl"] = returnUrl
		this.Data["token"] = this.token()
		this.page.Title = "登录"
		this.TplNames = this.getTplFileName("login")
		return
	}
	//Post
	password := this.GetString("password")
	always := this.GetString("always") == "on"

	//登录成功后转向地址
	if returnUrl == "" {
		returnUrl = "/"
	}

	var users = &models.Users{}
	u, err := users.Login(loginName, password, this.xm)
	if err != nil {
		//this.Redirect(fmt.Sprintf("/home/login?loginname=%s&returnurl=%s", loginName, returnUrl), 302)
		data := utils.JsonMessage(false, err.Error(), this.lang(err.Error()))
		this.renderJson(data)
		return
	}

	//登录状态存入cookie，缺省时间是1年：365*24*60*60
	var cookieDuration interface{}
	if always {
		cookieDuration, _ = beego.AppConfig.Int("CookieDuration")

	} else {
		cookieDuration = ""
	}

	this.Ctx.SetCookie(beego.AppConfig.String("CookieName"),
		utils.CookieEncode(fmt.Sprintf("%d|%s|%d", u.Id, u.Relname, u.Role)),
		cookieDuration, "/")

	this.Redirect(returnUrl, 302)
}

//注销
func (this *Front) Logout() {
	this.Ctx.SetCookie(beego.AppConfig.String("CookieName"), "", 0, "/")
	this.Redirect("/", 302)
}

func (this *Front) getTplFileName(s string) string {
	return getTplFileName("home", s)
}
