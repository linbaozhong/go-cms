package controllers

import (
	//"fmt"
	"cms/models"
	"cms/utils"
)

type Image struct {
	Admin
}

var images = &models.Images{}

//首页
func (this *Image) Index() {
	id, err := this.getParamsInt64(":id")

	if err != nil {
		this.errorHandle(utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams")))
		return
	}

	article, err := articles.Get(id)

	if err == nil {
		this.Data["article"] = article
		is, _ := images.GetAll(id)
		this.Data["images"] = is
	}

	this.TplNames = this.getTplFileName("index")
	this.Render()
}

//图片列表
func (this *Image) GetAll() {
	us, err := images.GetAll()

	var data interface{}
	if err == nil {
		data = utils.JsonResult(true, "", us)
	} else {
		data = utils.JsonMessage(false, "", "")
	}

	this.renderJson(data)
}

//新增图片
func (this *Image) Create() {

	//Get方法
	if this.methodGet {
		id, err := this.getParamsInt64(":id")
		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams")))
			return
		}
		this.Data["articleid"] = id
		this.Data["token"] = this.token()
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
	m := new(models.Images)

	models.Extend(m, this.xm)

	m.Articleid, _ = this.GetInt("articleid")
	m.Title = this.GetString("title")
	m.Url = this.GetString("url")

	if this.GetString("status") == "on" {
		m.Status = 1
	} else {
		m.Status = 0
	}

	files, err := this.upload("file")
	if err == nil {
		m.Path = files[0].Path
		m.Ext = files[0].Ext
		m.Srcfilename = files[0].Name
		m.Size = files[0].Size
	}

	//数据合法性检验
	if data, inv := this.invalidModel(m); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	var data interface{}
	_, err = images.Add(m)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonResult(true, "", files)
	}
	this.renderJson(data)
}

//修改账户信息
func (this *Image) Edit() {

	//Get方法
	if this.methodGet {

		id, err := this.getParamsInt64(":id")
		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams")))
			return
		}

		c, err := images.Get(id)

		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "", err.Error()))
			return
		}
		this.Data["image"] = c

		this.Data["token"] = this.token()
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
	m, err := images.Get(id)
	if err != nil {
		this.errorHandle(utils.JsonMessage(false, "", err.Error()))
		return
	}
	//赋值
	m.Id = id
	m.Articleid, _ = this.GetInt("articleid")
	m.Title = this.GetString("title")
	m.Url = this.GetString("url")
	m.Updated = this.xm.Updated
	m.Updator = this.xm.Updator
	m.Ip = this.xm.Ip
	//
	if this.GetString("status") == "on" {
		m.Status = 1
	} else {
		m.Status = 0
	}

	files, err := this.upload("file")
	if err == nil && len(files) > 0 {
		m.Path = files[0].Path
		m.Ext = files[0].Ext
		m.Srcfilename = files[0].Name
		m.Size = files[0].Size
	}

	//数据合法性检验
	if data, inv := this.invalidModel(m); inv {
		this.renderJson(data)
		return
	}
	//提交
	_, err = images.Update(m)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonResult(true, "", files)
	}

	this.renderJson(data)
}

//重新排列图片顺序
func (this *Image) Sequence() {
	var data interface{}
	//
	id, err1 := this.GetInt("id")
	sq, err2 := this.GetInt("sq")

	if err1 != nil || err2 != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err1.Error()+"\n"+err2.Error())
	} else {
		this.xm.Sequence = int(sq)
		_, err := images.SetSequence(id, this.xm)
		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "resetFail", err.Error())
		}
	}
	this.renderJson(data)
}

/*
重置图片
method：post
params：id
*/
func (this *Image) Reset() {
	var data interface{}
	//
	id, err := this.GetInt("id")
	if err != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err.Error())
	} else {
		//
		_, err = images.Reset(id, this.xm)
		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "resetFail", err.Error())
		}
	}
	this.renderJson(data)
}

/*
删除图片
method：post
params：ids
*/
func (this *Image) Delete() {
	var data interface{}
	//
	id := this.getParamsString(":id")
	//
	ids := this.GetStrings("ids")

	if id != "" {
		ids = append(ids, id)
	}

	_, err := images.Delete(ids, this.xm)
	if err == nil {
		data = utils.JsonMessage(true, "", "")
	} else {
		data = utils.JsonMessage(false, "deleteFail", this.lang("deleteFail"))
	}
	this.renderJson(data)
}

//
func (this *Image) getTplFileName(s string) string {
	return getTplFileName("admin/image", s)
}
