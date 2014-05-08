package models

import (
	//"errors"
	"cms/utils"
	"fmt"
	"github.com/coocood/qbs"
	"strings"
)

type Images struct {
	Id          int64
	Articleid   int64
	Title       string `valid:"MaxSize(150)"`
	Path        string `valid:"Required;MaxSize(255)"`
	Ext         string `valid:"Required;MaxSize(10)"`
	Size        int64
	Srcfilename string `valid:"Required;MaxSize(255)"`
	Url         string `valid:"MaxSize(255)"`
	Sequence    int
	Status      int8
	Deleted     int8
	Creator     int64
	Created     int64
	Updator     int64
	Updated     int64
	Ip          string
}

/*
新增频道
*/
func (this *Images) Add(m *Images) (int64, error) {
	defer db.Close()
	return db.Save(m)
}

/*
修改频道
*/
func (this *Images) Update(m *Images) (int64, error) {
	defer db.Close()
	type Images struct {
		Articleid   int64
		Title       string `valid:"MaxSize(150)"`
		Path        string `valid:"Required;MaxSize(255)"`
		Ext         string `valid:"Required;MaxSize(10)"`
		Size        int64
		Srcfilename string `valid:"Required;MaxSize(255)"`
		Url         string `valid:"MaxSize(255)"`
		Updator     int64
		Updated     int64
		Ip          string
	}
	c := new(Images)
	c.Articleid = m.Articleid
	c.Title = m.Title
	c.Path = m.Path
	c.Ext = m.Ext
	c.Srcfilename = m.Srcfilename
	c.Url = m.Url
	c.Updated = m.Updated
	c.Updator = m.Updator
	c.Ip = m.Ip

	return db.WhereEqual("id", m.Id).Update(c)
}

//获取一个可用图片，屏蔽禁用或已删除的图片
func (this *Images) Get(id int64) (*Images, error) {
	defer db.Close()
	m := new(Images)

	condition := qbs.NewEqualCondition("deleted", 0).AndEqual("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//获取一个图片
func (this *Images) GetEx(id int64) (*Images, error) {
	defer db.Close()
	m := new(Images)

	condition := qbs.NewEqualCondition("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//分页获取图片
func (this *Images) GetImages(args ...interface{}) (us []*Images, err error) {
	defer db.Close()

	condition := qbs.NewEqualCondition("Deleted", utils.DelNormal).AndEqual("Status", utils.StatEnabled)

	if len(args) > 0 && args[0].(int64) > 0 {
		condition = condition.AndEqual("articleid", args[0].(int64))
	}

	err = db.Condition(condition).OmitFields("Created", "Creator", "Updated", "Updator", "Ip").OrderByDesc("sequence").FindAll(&us)
	return
}

//分页获取图片
func (this *Images) GetAll(args ...interface{}) (us []*Images, err error) {
	defer db.Close()

	condition := qbs.NewEqualCondition("Deleted", utils.DelNormal)

	if len(args) > 0 && args[0].(int64) > 0 {
		condition = condition.AndEqual("articleid", args[0].(int64))
	}

	err = db.Condition(condition).OmitFields("Created", "Creator", "Updated", "Updator", "Ip").OrderByDesc("sequence").FindAll(&us)
	return
}

//分页获取图片
func (this *Images) GetAllEx() (us []*Images, err error) {
	defer db.Close()

	err = db.OmitFields("Created", "Creator", "Updated", "Updator", "Ip").OrderByDesc("sequence").FindAll(&us)
	return
}

//图片顺序
func (this *Images) SetSequence(id int64, f *Field) (int64, error) {
	defer db.Close()

	type images struct {
		Sequence int
		Updator  int64
		Updated  int64
		Ip       string
	}

	u := new(images)
	u.Sequence = f.Sequence
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", id).Update(u)
}

/*
重置图片
禁用、启用图片
*/
func (this *Images) Reset(id int64, f *Field) (int64, error) {
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
	type Images struct {
		Status  int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(Images)
	u.Status = status
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", m.Id).Update(u)
}

/*
删除图片
*/
func (this *Images) Delete(ids []string, f *Field) (int64, error) {
	defer db.Close()
	type Images struct {
		Deleted int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(Images)
	u.Deleted = utils.DelDeleted
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip
	fmt.Println(ids)
	return db.WhereIn("id", qbs.StringsToInterfaces(strings.Join(ids, ","))).Update(u)
}
