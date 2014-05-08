package controllers

import (
	//"fmt"
	"cms/models"
	"cms/utils"
)

type Channel struct {
	Admin
}

var channels = &models.Channels{}

//首页
func (this *Channel) Index() {
	cs, _ := channels.GetAll(-1, int64(0))

	this.Data["channels"] = cs
	//所属频道选项
	//this.Data["chs"] = channels.GetChannelSelectItems(0)
	//频道类型
	this.Data["types"] = channels.GetTypeSelectItems(-1)

	this.TplNames = this.getTplFileName("index")
	this.Render()
}

//频道列表
func (this *Channel) GetAll() {
	t, err := this.GetInt("type")
	if err != nil {
		t = -1
	}
	us, err := channels.GetAll(int(t))

	var data interface{}
	if err == nil {
		data = utils.JsonResult(true, "", us)
	} else {
		data = utils.JsonMessage(false, "", err.Error())
	}

	this.renderJson(data)
}

//新增频道
func (this *Channel) Create() {
	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()
		//所属频道选项
		this.Data["chs"] = channels.GetChannelSelectItems(0, utils.ChNavigation)
		//频道类型
		this.Data["types"] = channels.GetTypeSelectItems(0)

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
	m := new(models.Channels)

	models.Extend(m, this.xm)

	m.Pid, _ = this.GetInt("pid")
	m.Name = this.GetString("name")
	m.Enname = this.GetString("enname")
	t, _ := this.GetInt("children")
	m.Children = int(t)
	t, _ = this.GetInt("type")
	m.Type = int8(t)

	if this.GetString("status") == "on" {
		m.Status = 1
	} else {
		m.Status = 0
	}

	//数据合法性检验
	if data, inv := this.invalidModel(m); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	var data interface{}
	_, err := channels.Add(m)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonMessage(true, "", "")
	}
	this.renderJson(data)
}

//修改账户信息
func (this *Channel) Edit() {
	//Get方法
	if this.methodGet {
		this.Data["token"] = this.token()

		id, err := this.getParamsInt64(":id")
		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "invalidRequestParams", this.lang("invalidRequestParams")))
			return
		}

		c, err := channels.Get(id)

		if err != nil {
			this.errorHandle(utils.JsonMessage(false, "", err.Error()))
			return
		}
		this.Data["channel"] = c
		//所属频道选项
		this.Data["chs"] = channels.GetChannelSelectItems(-1, utils.ChNavigation, c.Pid)
		//频道类型
		this.Data["types"] = channels.GetTypeSelectItems(c.Type)

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
	m, err := channels.Get(id)
	if err != nil {
		this.errorHandle(utils.JsonMessage(false, "", err.Error()))
		return
	}
	//赋值
	m.Pid, _ = this.GetInt("pid")
	m.Name = this.GetString("name")
	m.Enname = this.GetString("enname")
	t, _ := this.GetInt("children")
	m.Children = int(t)
	t, _ = this.GetInt("type")
	m.Type = int8(t)
	m.Updated = this.xm.Updated
	m.Updator = this.xm.Updator
	m.Ip = this.xm.Ip

	if this.GetString("status") == "on" {
		m.Status = 1
	} else {
		m.Status = 0
	}

	//数据合法性检验
	if data, inv := this.invalidModel(m); inv {
		this.renderJson(data)
		return
	}
	//提交DDL
	_, err = channels.Update(m)

	if err != nil {
		data = utils.JsonMessage(false, "", err.Error())
	} else {
		data = utils.JsonMessage(true, "", "")
	}
	this.renderJson(data)
}

/*
重置频道
method：post
params：id
*/
func (this *Channel) Reset() {
	var data interface{}
	//
	id, err := this.GetInt("id")
	if err != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err.Error())
	} else {

		_, err = channels.Reset(id, this.xm)

		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "resetFail", err.Error())
		}

	}
	this.renderJson(data)
}

/*
删除频道
method：post
params：ids
*/
func (this *Channel) Delete() {
	var data interface{}
	id := this.getParamsString(":id")
	//
	ids := this.GetStrings("ids")

	if id != "" {
		ids = append(ids, id)
	}
	//
	var err error
	//
	_, err = channels.Delete(ids, this.xm)
	if err == nil {
		data = utils.JsonMessage(true, "", "")
	} else {
		data = utils.JsonMessage(false, "deleteFail", this.lang("deleteFail"))
	}
	this.renderJson(data)
}

//重新排列频道顺序
func (this *Channel) Sequence() {
	var data interface{}
	//
	id, err1 := this.GetInt("id")
	sq, err2 := this.GetInt("sq")

	if err1 != nil || err2 != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err1.Error()+"\n"+err2.Error())
	} else {
		this.xm.Sequence = int(sq)
		_, err := channels.SetSequence(id, this.xm)
		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "resetFail", err.Error())
		}
	}
	this.renderJson(data)
}

//子项限制
func (this *Channel) Children() {
	var data interface{}
	//
	id, err1 := this.GetInt("id")
	sq, err2 := this.GetInt("sq")

	if err1 != nil || err2 != nil {
		data = utils.JsonMessage(false, "invalidRequestParams", err1.Error()+"\n"+err2.Error())
	} else {
		this.xm.Sequence = int(sq)
		_, err := channels.SetChildren(id, this.xm)
		if err == nil {
			data = utils.JsonMessage(true, "", "")
		} else {
			data = utils.JsonMessage(false, "resetFail", err.Error())
		}
	}
	this.renderJson(data)
}

//
func (this *Channel) getTplFileName(s string) string {
	return getTplFileName("admin/channel", s)
}
