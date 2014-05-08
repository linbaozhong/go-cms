package models

import (
	//"errors"
	"cms/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/coocood/qbs"
)

type Channels struct {
	Id       int64
	Pid      int64
	Name     string `valid:"Required;MaxSize(20)"`
	Enname   string `valid:"MaxSize(20)"`
	Type     int8
	Children int
	Sequence int
	Level    int8
	Path     string
	Status   int8
	Deleted  int8
	Creator  int64
	Created  int64
	Updator  int64
	Updated  int64
	Ip       string
}

/*
新增频道
*/
func (this *Channels) Add(m *Channels) (int64, error) {
	defer db.Close()

	if err := this.perfectModel(m); err != nil {
		return 0, err
	}

	id, err := db.Save(m)
	if err != nil {
		return id, err
	}
	if err := this.perfectModel(m); err != nil {
		return id, err
	}
	return db.Save(m)
}

/*
修改频道
*/
func (this *Channels) Update(m *Channels) (int64, error) {
	defer db.Close()
	type Channels struct {
		Pid      int64
		Name     string `valid:"Required;MaxSize(20)"`
		Enname   string `valid:"MaxSize(20)"`
		Type     int8
		Children int
		Level    int8
		Path     string
		Status   int8
		Updator  int64
		Updated  int64
		Ip       string
	}

	c := new(Channels)
	c.Pid = m.Pid
	c.Name = m.Name
	c.Enname = m.Enname
	c.Type = m.Type
	c.Children = m.Children
	c.Status = m.Status
	c.Updated = m.Updated
	c.Updator = m.Updator
	c.Ip = m.Ip

	if err := this.perfectModel(m); err != nil {
		return 0, err
	}
	c.Level = m.Level
	//c.Path = m.Path

	return db.WhereEqual("id", m.Id).Update(c)
}

/*
完善数据对象
新建或修改的频道需要重新计算 频道的层级和路径
*/
func (this *Channels) perfectModel(m *Channels) error {
	if m.Pid == 0 {
		m.Level = 0
		m.Path = fmt.Sprintf("%d", m.Id)
	} else {
		ch, err := this.GetEx(m.Pid)
		if err != nil {
			return err
		}
		m.Level = ch.Level + 1
		m.Path = fmt.Sprintf("%s/%d", ch.Path, m.Id)
	}
	return nil
}

//获取一个可用频道，屏蔽禁用或已删除的频道
func (this *Channels) Get(id int64) (*Channels, error) {
	defer db.Close()
	m := new(Channels)

	condition := qbs.NewEqualCondition("Deleted", utils.DelNormal).AndEqual("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//
func (this *Channels) GetByName(enname string) (*Channels, error) {
	defer db.Close()
	m := new(Channels)

	condition := qbs.NewEqualCondition("Deleted", utils.DelNormal).AndEqual("enname", enname)
	err := db.Condition(condition).Find(m)
	return m, err
}

//获取一个频道
func (this *Channels) GetEx(id int64) (*Channels, error) {
	defer db.Close()
	m := new(Channels)

	condition := qbs.NewEqualCondition("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//分页获取频道
func (this *Channels) GetAll(args ...interface{}) ([]*Channels, error) {
	var t int = -1      //类型
	var pid int64 = -1  //父id
	var status int = -1 //状态
	var level int = -1  //层级

	if n := len(args); n > 0 {
		if args[0].(int) > -1 {
			t = args[0].(int)
		}

		if n > 1 {
			if args[1].(int64) > -1 {
				pid = args[1].(int64)

			}
			if n > 2 {
				status = args[2].(int)
				if n > 3 {
					level = args[3].(int)
				}
			}
		}

	}

	us := make([]*Channels, 0)
	return this.getChannels(t, pid, status, utils.DelNormal, us, level)
}

//递归
func (this *Channels) getChannels(t int, pid int64, status, deleted int, us []*Channels, level int) ([]*Channels, error) {
	defer db.Close()
	//如果层级=0，终止
	if level == 0 {
		return us, nil
	} else {
		level--
	}

	condition := qbs.NewCondition("id>?", 0)
	if t > -1 {
		condition.AndEqual("type", t)
	}
	if pid > -1 {
		condition.AndEqual("pid", pid)
	}
	if deleted > -1 {
		condition.AndEqual("deleted", deleted)
	}
	if status > -1 {
		condition.AndEqual("status", status)
	}

	//fmt.Println(t, pid)
	var tmp []*Channels
	err := db.Condition(condition).OmitFields("Created", "Creator", "Updated", "Updator", "Ip").OrderByDesc("sequence").FindAll(&tmp)

	if len(tmp) > 0 {
		for _, c := range tmp {
			us = append(us, c)
			us, err = this.getChannels(t, c.Id, status, deleted, us, level)
		}
	}

	return us, err
}

//分页获取频道
func (this *Channels) GetAllEx(args ...interface{}) ([]*Channels, error) {
	var t int = -1      //类型
	var pid int64 = -1  //父id
	var status int = -1 //状态
	var level int = -1  //层级

	if n := len(args); n > 0 {
		if args[0].(int) > -1 {
			t = args[0].(int)
		}

		if n > 1 {
			if args[1].(int64) > -1 {
				pid = args[1].(int64)

			}
			if n > 2 {
				status = args[2].(int)
				if n > 3 {
					level = args[3].(int)
				}
			}
		}

	}

	var us []*Channels
	return this.getChannels(t, pid, status, -1, us, level)
}

/*
重置频道
禁用、启用频道
*/
func (this *Channels) Reset(id int64, f *Field) (int64, error) {
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
	type channels struct {
		Status  int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(channels)
	u.Status = status
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", m.Id).Update(u)
}

//频道顺序
func (this *Channels) SetSequence(id int64, f *Field) (int64, error) {
	defer db.Close()

	type channels struct {
		Sequence int
		Updator  int64
		Updated  int64
		Ip       string
	}

	u := new(channels)
	u.Sequence = f.Sequence
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", id).Update(u)
}

//子项限制
func (this *Channels) SetChildren(id int64, f *Field) (int64, error) {
	defer db.Close()

	type channels struct {
		Children int
		Updator  int64
		Updated  int64
		Ip       string
	}

	u := new(channels)
	u.Children = f.Sequence
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", id).Update(u)
}

/*
删除频道
*/
func (this *Channels) Delete(ids []string, f *Field) (int64, error) {
	defer db.Close()
	type channels struct {
		Deleted int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(channels)
	u.Deleted = utils.DelDeleted
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereIn("id", qbs.StringsToInterfaces(strings.Join(ids, ","))).Update(u)
}

//所属频道SelectItems
func (this *Channels) GetChannelSelectItems(ids ...int64) (s []*SelectItem) {
	//
	var id, selectid int64 = 0, 0
	var od int = -1

	n := len(ids)
	if n > 0 {
		id = ids[0]
		if n > 1 {
			od = int(ids[1])
			if n > 2 {
				selectid = ids[2]
			}
		}
	}

	us, err := this.GetAll(od, id)

	s = append(s, &SelectItem{Key: "--选择频道--", Value: "0"})
	//var data interface{}
	if err == nil {
		n := len(us)

		for i := 0; i < n; i++ {
			si := new(SelectItem)
			si.Key = utils.Indent(us[i].Name, us[i].Level)
			si.Value = strconv.Itoa(int(us[i].Id))
			si.Selected = selectid == us[i].Id

			s = append(s, si)
		}

		//data = utils.JsonResult(true, "", s)
	} else {
		//data = utils.JsonMessage(false, "", err.Error())
	}

	//this.renderJson(data)
	return
}

//所属频道类型SelectItems
func (this *Channels) GetTypeSelectItems(id int8) (s []*SelectItem) {

	s = append(s, &SelectItem{Key: "--选择分类--", Value: "-1"})
	s = append(s, &SelectItem{Key: "导航", Value: "0", Selected: id == 0})
	s = append(s, &SelectItem{Key: "文章", Value: "1", Selected: id == 1})

	return
}

//取指定频道名称的父id
func (this *Channels) GetParentId(enname string) int64 {
	db.Close()
	if enname == "" {
		return 0
	}
	ch, _ := this.GetByName(enname)
	return ch.Pid
}
