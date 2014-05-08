package controllers

import (
	"cms/models"
	"cms/utils"
	"strings"
)

type Home struct {
	Front
}

/*
读取文章
参数：
	第一个：导航频道英文名称
	第二个：下属文章频道id
	第三个：文章id
	第四个：页码
*/
func (this *Home) Cn() {
	//读取记录数限制
	p := new(models.Pagination)
	p.Index = 1

	var action string
	var childChannelid, articleid int64

	args := this.Ctx.Input.Params
	n := len(args)
	//取传入的参数
	if n > 0 {
		action = args["0"]
		if n > 1 {
			childChannelid, _ = utils.Str2int64(args["1"]) //strconv.ParseInt(args["1"], 10, 64)
			if n > 2 {
				articleid, _ = utils.Str2int64(args["2"]) //strconv.ParseInt(args["2"], 10, 64)
				if n > 3 {
					p.Index, _ = utils.Str2int(args["3"])
					if p.Index == 0 {
						p.Index += 1
					}
				}
			}
		}
	}
	////取导航
	//nvas, err := channels.GetAll(utils.ChNavigation, int64(0), utils.StatEnabled)
	//if err != nil {
	//	return
	//}
	//this.Data["navs"] = nvas
	//fmt.Println(nvas[0])
	//
	if strings.TrimSpace(action) == "" {
		action = "index"
	}
	this.Data["action"] = action
	if action == "index" {
		this.TplNames = this.getTplFileName(action)
		return
	}

	//取当前频道channelid
	//if channelid == 0 {
	// for _, v := range nvas {
	// 	if v.Enname == action {
	// 		channelid = v.Id
	// 		this.page.Title += " - " + v.Name
	// 		break
	// 	}
	// }
	//}

	ch, err := channels.GetByName(action)
	if err != nil {
		return
	}
	//this.Data["channelId"] = ch.Id
	this.Data["channel"] = ch
	this.page.Title += " - " + ch.Name

	//取当前频道的子频道
	chs, _ := channels.GetAll( /*utils.ChNews*/ -1, ch.Id, utils.StatEnabled)
	//如果channelid为空，取子频道第一个channel的Id
	if len(chs) > 0 {
		if childChannelid == 0 {
			childChannelid = chs[0].Id
		}
		//分页记录数
		for i, v := range chs {
			if v.Id == childChannelid {
				p.Size = int(chs[i].Children)
				break
			}
		}
	}

	this.Data["cases"] = chs
	this.Data["currentChannelId"] = childChannelid

	//取频道下的全部文章列表,如果p.Size==0,p.Size==1
	if p.Size < 1 {
		p.Size = 1
	}
	if childChannelid > 0 && p.Size > 0 {
		as, _ := articles.GetArticles(childChannelid, p)
		if articleid == 0 && len(as) > 0 {
			articleid = as[0].Id
		}
		this.Data["articles"] = as
	} else {
		this.Data["articles"] = make([]interface{}, 0)
	}

	this.Data["currentArticleId"] = articleid

	//取文章和相应图片
	article := new(models.Articles)
	if articleid > 0 {
		//文章
		article, _ = articles.Get(articleid)
		//图片
		imgs, _ := images.GetImages(articleid)
		this.Data["images"] = imgs
	} else {
		this.Data["images"] = make([]interface{}, 0)
	}
	this.Data["article"] = article
	//分页对象
	p.Prev = p.Index - 1
	//检查是否到末尾
	if p.Index*p.Size >= p.Count {
		p.Next = 0
	} else {
		p.Next = p.Index + 1
	}
	this.Data["pagination"] = p

	//页面公共信息
	this.page.Description += "," + article.Description
	this.page.Keywords += "," + article.Keywords
	this.Data["page"] = this.page
	//模板
	this.TplNames = this.getTplFileName("case")
}

//
func (this *Home) getTplFileName(s string) string {
	return getTplFileName("home", s)
}
