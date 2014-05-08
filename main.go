package main

import (
	"cms/controllers"
	"cms/models"
	"cms/utils"
	"path"

	"github.com/astaxie/beego"
)

func initconfig() {
	beego.AppPath = utils.AppRoot
	beego.AppConfigPath = path.Join(beego.AppPath, "conf", "app.conf")
	beego.ParseConfig()
	beego.ViewsPath = utils.MergePath(beego.ViewsPath)
}

func main() {
	initconfig()

	var displayName = beego.AppConfig.String("ServiceDisplayName")

	doWork(displayName, true)

}

/*
启动服务
name：显示名称
service：service为true，desktop为false
*/
func doWork(name string, service bool) {
	//桌面服务日志
	beego.BeeLogger.SetLogger("file", `{"filename":"`+utils.Sqlite3Path(beego.AppConfig.String("LogPath"))+`"}`)
	beego.SetLevel(beego.LevelWarning)
	//日志
	beego.Error(name + ` 服务启动`)

	//模板函数
	templateFunc()
	//路由
	router()
	//主要是初始化 i18n 文件
	utils.I18n()
	//初始化数据库注册
	models.Init()

	beego.Run()

}

//模板函数
func templateFunc() {
	//将自1970-1-1开始的毫秒数转为时间
	beego.AddFuncMap("Msec2Time", utils.Msec2Time)
	//树形列表缩进
	beego.AddFuncMap("Indent", utils.Indent)
	//狗血的相加函数
	beego.AddFuncMap("Plus", func(a, b int) int {
		return a + b
	})
	//狗血的相乘函数
	beego.AddFuncMap("Multiply", func(a, b int) int {
		return a * b
	})
	//生成导航条函数
	beego.AddFuncMap("Navibar", controllers.Navibar)
	//生成文章频道列表函数

	//生成文章内容函数
}

//路由
func router() {
	beego.DirectoryIndex = true
	beego.SetStaticPath("/test", utils.MergePath("test"))

	beego.AutoRender = false
	beego.SetStaticPath("/static", utils.MergePath("static"))
	beego.SetStaticPath("/upload", utils.MergePath(beego.AppConfig.String("UploadPath")))
	//首页
	home := &controllers.Home{}
	beego.AutoRouter(home)
	beego.Router("/", home, "*:Cn")
	//登录
	beego.Router("/home/login", home, "get,post:Login")
	beego.Router("/home/logout", home, "get,post:Logout")
	//后台管理
	admin := &controllers.Admin{}
	beego.Router("/admin", admin, "*:Index")
	beego.Router("/admin/index", admin, "*:Index")
	beego.Router("/admin/upload", admin, "post:Upload")
	//个人信息
	profile := &controllers.Profile{}
	beego.Router("/admin/profile", profile, "get:Index")
	beego.Router("/admin/profile/password", profile, "get,post:UpdatePassword")
	beego.Router("/admin/profile/edit", profile, "get,post:UpdateProfile")
	//账户管理
	account := &controllers.Account{}
	beego.Router("/admin/account", account, "get:Index")
	beego.Router("/admin/account/create", account, "*:Create")
	beego.Router("/admin/account/edit", account, "post:Edit")
	beego.Router("/admin/account/edit/:id:int", account, "get:Edit")
	beego.Router("/admin/account/delete/:id:int", account, "*:Delete")
	beego.Router("/admin/account/reset", account, "*:Reset")
	//文章管理
	article := &controllers.Article{}
	beego.Router("/admin/article", article, "get:Index")
	beego.Router("/admin/article/create", article, "*:Create")
	beego.Router("/admin/article/edit", article, "post:Edit")
	beego.Router("/admin/article/edit/:id:int", article, "get:Edit")
	beego.Router("/admin/article/delete/:id:int", article, "*:Delete")
	beego.Router("/admin/article/sequence", article, "*:Sequence")
	beego.Router("/admin/article/reset", article, "*:Reset")
	beego.Router("/admin/article/getall", article, "*:GetAll")
	//图片管理
	image := &controllers.Image{}
	beego.Router("/admin/image/index/:id:int", image, "get:Index")
	beego.Router("/admin/image/create", image, "post:Create")
	beego.Router("/admin/image/edit", image, "post:Edit")
	beego.Router("/admin/image/create/:id:int", image, "get:Create")
	beego.Router("/admin/image/edit/:id:int", image, "get:Edit")
	beego.Router("/admin/image/delete/:id:int", image, "*:Delete")
	beego.Router("/admin/image/sequence", image, "*:Sequence")
	beego.Router("/admin/image/reset", image, "*:Reset")
	//频道管理
	ch := &controllers.Channel{}
	beego.Router("/admin/channel", ch, "get:Index")
	beego.Router("/admin/channel/create", ch, "*:Create")
	beego.Router("/admin/channel/edit", ch, "post:Edit")
	beego.Router("/admin/channel/edit/:id:int", ch, "get:Edit")
	beego.Router("/admin/channel/delete/:id:int", ch, "*:Delete")
	beego.Router("/admin/channel/sequence", ch, "*:Sequence")
	beego.Router("/admin/channel/children", ch, "*:Children")
	beego.Router("/admin/channel/reset", ch, "*:Reset")
	beego.Router("/admin/channel/getall", ch, "*:GetAll")
}
