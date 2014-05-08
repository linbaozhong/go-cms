package controllers

import (
	"cms/models"
	"cms/utils"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type Article struct {
	Admin
}

var articles = &models.Articles{}

//首页
func (this *Article) Index() {
	// as, _ := articles.GetAll()

	// this.Data["articles"] = as
	//所属频道选项
	this.Data["chs"] = channels.GetChannelSelectItems(0 /*, utils.ChNews*/)

	this.TplNames = this.getTplFileName("index")
	this.Render()
}

//频道列表
func (this *Article) GetAll() {
	var data interface{}
	//
	page := new(models.Pagination)
	if v, err := this.GetInt("index"); err == nil {
		page.Index = int(v)
	}
	if v, err := this.GetInt("size"); err == nil {
		page.Size = int(v)
	}

	//参数错误
	if page.Index <= 0 {
		page.Index = 1
	}
	if page.Size <= 0 {
		//data = utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams"))
		page.Size = 15
	}
	channelid, _ := this.GetInt("channelid")

	//
	us, err := articles.GetAll(channelid, page)
	if err == nil {
		data = utils.JsonResult(true, fmt.Sprintf("%d", page.Count), us)
	} else {
		data = utils.JsonMessage(false, "", "")
	}

	this.renderJson(data)
}

//新增频道
func (this *Article) Create() {

	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()
		//所属频道选项
		this.Data["chs"] = channels.GetChannelSelectItems(0 /*, utils.ChNews*/)

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
	m := new(models.Articles)

	models.Extend(m, this.xm)

	m.Channelid, _ = this.GetInt("channelid")
	m.Title = this.GetString("title")
	m.Subtitle = this.GetString("subtitle")
	m.Intro = this.GetString("intro")
	m.Content = this.GetString("content")
	m.Keywords = this.GetString("keywords")
	m.Description = this.GetString("description")
	m.Author = this.GetString("author")

	if this.GetString("status") == "on" {
		m.Status = 1
	} else {
		m.Status = 0
	}

	if t, err := beego.DateParse(this.GetString("published"), "Y-n-j H:i:s"); err == nil {
		m.Published = utils.Millisecond(t)
	} else {
		m.Published = utils.Millisecond(time.Now())
	}

	//数据合法性检验
	if data, inv := this.invalidModel(m); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	var data interface{}
	_, err := articles.Add(m)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonMessage(true, "", "")
	}
	this.renderJson(data)
}

//修改账户信息
func (this *Article) Edit() {

	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()

		id, err := this.getParamsInt64(":id")
		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams")))
			return
		}

		c, err := articles.Get(id)

		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "", err.Error()))
			return
		}
		this.Data["article"] = c
		//所属频道选项
		//this.Data["chs"] = channels.GetChannelSelectItems(-1, utils.ChNews, c.Channelid)
		this.Data["chs"] = channels.GetChannelSelectItems(0, -1, c.Channelid)

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
	//提交DDL
	var data interface{}

	id, err := this.GetInt("id")
	if err != nil || id == 0 {
		this.errorHandle(utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams")))
		return
	}

	//获取原始数据模型
	m, err := articles.Get(id)
	if err != nil {
		this.errorHandle(utils.JsonMessage(false, "", err.Error()))
		return
	}
	//赋值
	m.Channelid, _ = this.GetInt("channelid")
	m.Title = this.GetString("title")
	m.Subtitle = this.GetString("subtitle")
	m.Intro = this.GetString("intro")
	m.Content = this.GetString("content")
	m.Keywords = this.GetString("keywords")
	m.Description = this.GetString("description")
	m.Author = this.GetString("author")
	m.Updated = this.xm.Updated
	m.Updator = this.xm.Updator
	m.Ip = this.xm.Ip

	if this.GetString("status") == "on" {
		m.Status = 1
	} else {
		m.Status = 0
	}

	if t, err := beego.DateParse(this.GetString("published"), "Y-n-j"); err == nil {
		m.Published = utils.Millisecond(t)
	} else {
		m.Published = utils.Millisecond(time.Now())
	}

	//数据合法性检验
	if data, inv := this.invalidModel(m); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	_, err = articles.Update(m)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonMessage(true, "", "")
	}
	this.renderJson(data)
}

/*
重置文章
method：post
params：id
*/
func (this *Article) Reset() {
	var data interface{}
	//
	id, err := this.GetInt("id")
	if err != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err.Error())
	} else {
		//
		_, err = articles.Reset(id, this.xm)
		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "resetFail", err.Error())
		}
	}
	this.renderJson(data)
}

/*
删除文章
method：post
params：ids
*/
func (this *Article) Delete() {
	var data interface{}
	//
	id := this.getParamsString(":id")
	//
	ids := this.GetStrings("ids")

	if id != "" {
		ids = append(ids, id)
	}
	//
	var err error
	//
	_, err = articles.Delete(ids, this.xm)
	if err == nil {
		data = utils.JsonMessage(true, "", "")
	} else {
		data = utils.JsonMessage(false, "deleteFail", this.lang("deleteFail"))
	}
	this.renderJson(data)
}

//重新排列文章顺序
func (this *Article) Sequence() {
	var data interface{}
	//
	id, err1 := this.GetInt("id")
	sq, err2 := this.GetInt("sq")

	if err1 != nil || err2 != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err1.Error()+"\n"+err2.Error())
	} else {
		this.xm.Sequence = int(sq)
		_, err := articles.SetSequence(id, this.xm)
		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "resetFail", err.Error())
		}
	}
	this.renderJson(data)
}

//
func (this *Article) getTplFileName(s string) string {
	return getTplFileName("admin/article", s)
}
