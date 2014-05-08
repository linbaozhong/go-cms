package models

import (
	//"errors"
	//"fmt"
	"cms/utils"
	"strings"

	"github.com/coocood/qbs"
)

type Articles struct {
	Id          int64
	Channelid   int64
	Title       string `valid:"Required;MaxSize(150)"`
	Subtitle    string `valid:"MaxSize(150)"`
	Intro       string `valid:"MaxSize(255)"`
	Content     string `valid:"Required;"`
	Keywords    string `valid:"MaxSize(255)"`
	Description string `valid:"MaxSize(255)"`
	Author      string `valid:"MaxSize(100)"`
	Sequence    int
	Status      int8
	Deleted     int8
	Publisher   int64
	Published   int64
	Creator     int64
	Created     int64
	Updator     int64
	Updated     int64
	Ip          string
}

/*
新增频道
*/
func (this *Articles) Add(m *Articles) (int64, error) {
	defer db.Close()
	return db.Save(m)
}

/*
修改频道
*/
func (this *Articles) Update(m *Articles) (int64, error) {
	defer db.Close()
	type Articles struct {
		Channelid   int64
		Title       string `valid:"Required;MaxSize(150)"`
		Subtitle    string `valid:"Required;MaxSize(150)"`
		Intro       string `valid:"Required;MaxSize(255)"`
		Content     string `valid:"Required;"`
		Keywords    string `valid:"MaxSize(255)"`
		Description string `valid:"MaxSize(255)"`
		Author      string `valid:"MaxSize(100)"`
		Published   int64
		Publisher   int64
		Status      int8
		Updator     int64
		Updated     int64
		Ip          string
	}
	c := new(Articles)
	c.Channelid = m.Channelid
	c.Title = m.Title
	c.Subtitle = m.Subtitle
	c.Intro = m.Intro
	c.Content = m.Content
	c.Keywords = m.Keywords
	c.Description = m.Description
	c.Author = m.Author
	c.Status = m.Status
	c.Published = m.Published
	c.Publisher = m.Updator
	c.Updated = m.Updated
	c.Updator = m.Updator
	c.Ip = m.Ip

	return db.WhereEqual("id", m.Id).Update(c)
}

//获取一个可用文章，屏蔽禁用或已删除的文章
func (this *Articles) Get(id int64) (*Articles, error) {
	defer db.Close()
	m := new(Articles)

	condition := qbs.NewEqualCondition("deleted", utils.DelNormal).AndEqual("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//获取一个文章
func (this *Articles) GetEx(id int64) (*Articles, error) {
	defer db.Close()
	m := new(Articles)

	condition := qbs.NewEqualCondition("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//获取符合条件的记录数
func (this *Articles) Count(channelid int64, page *Pagination) {

}

//分页获取文章
func (this *Articles) GetArticles(channelid int64, page *Pagination) (us []*Articles, err error) {
	defer db.Close()

	condition := qbs.NewEqualCondition("Deleted", utils.DelNormal).AndEqual("Status", utils.StatEnabled)
	//文章频道
	if channelid > 0 {
		condition = condition.AndEqual("channelid", channelid)
	}
	//符合条件的记录数
	page.Count = int(db.Condition(condition).Count("articles"))

	err = db.Condition(condition).Limit(int(page.Size)).Offset(int((page.Index-1)*page.Size)).
		OmitFields("Keywords", "Description", "Author", "Intro", "Content", "Created", "Creator", "Updated", "Updator", "Ip").
		OrderByDesc("sequence").OrderByDesc("Updated").FindAll(&us)
	return
}

//分页获取文章
func (this *Articles) GetAll(channelid int64, page *Pagination) (us []*Articles, err error) {
	defer db.Close()

	condition := qbs.NewEqualCondition("Deleted", utils.DelNormal)
	//文章频道
	if channelid > 0 {
		condition = condition.AndEqual("channelid", channelid)
	}
	//符合条件的记录数
	page.Count = int(db.Condition(condition).Count("articles"))

	err = db.Condition(condition).Limit(int(page.Size)).Offset(int((page.Index-1)*page.Size)).
		OmitFields("Keywords", "Description", "Author", "Intro", "Content", "Created", "Creator", "Updated", "Updator", "Ip").
		OrderByDesc("sequence").OrderByDesc("Updated").FindAll(&us)

	return
}

//分页获取文章
func (this *Articles) GetAllEx() (us []*Articles, err error) {
	defer db.Close()

	err = db.OmitFields("Keywords", "Description", "Author", "Intro", "Content", "Created", "Creator", "Updated", "Updator", "Ip").FindAll(&us)
	return
}

/*
重置文章
禁用、启用文章
*/
func (this *Articles) Reset(id int64, f *Field) (int64, error) {
	defer db.Close()

	m, err := this.GetEx(id)
	if err != nil {
		return 0, err
	}
	//
	var status int8 = 0
	if m.Status == 0 {
		status = 1
	}
	type Articles struct {
		Status  int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(Articles)
	u.Status = status
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", m.Id).Update(u)
}

/*
删除文章
*/
func (this *Articles) Delete(ids []string, f *Field) (int64, error) {
	defer db.Close()
	type Articles struct {
		Deleted int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(Articles)
	u.Deleted = utils.DelDeleted
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereIn("id", qbs.StringsToInterfaces(strings.Join(ids, ","))).Update(u)
}

//文章顺序
func (this *Articles) SetSequence(id int64, f *Field) (int64, error) {
	defer db.Close()

	type articles struct {
		Sequence int
		Updator  int64
		Updated  int64
		Ip       string
	}

	u := new(articles)
	u.Sequence = f.Sequence
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", id).Update(u)
}
