package models

import (
	"errors"
	//"fmt"
	"cms/utils"
	"github.com/coocood/qbs"
	"strings"
)

type Users struct {
	Id        int64
	Loginname string `valid:"Required;MinSize(6);MaxSize(100)"`
	Password  string `valid:"Required;MinSize(6);MaxSize(20)"`
	Relname   string `valid:"Required;MinSize(6);MaxSize(20)"`
	Role      int8
	Status    int8
	Deleted   int8
	Updator   int64
	Updated   int64
	Ip        string
}

type Password struct {
	OldPassword string
	NewPassword string `valid:"Required;MinSize(6);MaxSize(20)"`
	RePassword  string
}

//登录
func (this *Users) Login(loginName, password string, f *Field) (*Users, error) {
	defer db.Close()

	u := new(Users)
	condition := qbs.NewEqualCondition("Loginname", loginName).AndEqual("Deleted", utils.DelNormal)
	err := db.Condition(condition).Find(u)

	if err != nil {
		return nil, errors.New("accoundNotFound")
	} else {
		//是否被锁定
		if u.Status == utils.StatDisabled {
			return nil, errors.New("accountLocked")
		}
		//校验密码
		if u.Password == utils.MD5(password) {
			//db.Save(u)
			return u, nil
		} else {
			return nil, errors.New("invalidPassword")
		}
	}

}

//修改密码
func (this *Users) UpdatePassword(f *Field, pass *Password) error {
	defer db.Close()
	//用户密码是否正确
	u, err := this.ValidPassword(f.Updator, pass.OldPassword)
	if err != nil {
		return err
	}
	//更新数据
	_, err = db.Exec("update users set password=?,updated=?,updator=?,ip=? where id=?",
		utils.MD5(pass.NewPassword), f.Updated, f.Updator, f.Ip, u.Id)

	return err
}

//验证用户密码是否合法
func (this *Users) ValidPassword(id int64, password string) (*Users, error) {
	defer db.Close()

	u := new(Users)

	err := db.WhereEqual("Id", id).Find(u)

	if err != nil {
		return nil, errors.New("accoundNotFound")
	} else {
		//校验密码
		if u.Password == utils.MD5(password) {
			return u, nil
		} else {
			return nil, errors.New("invalidPassword")
		}
	}
}

/*
新增账户
1、新增账户
2、增加账户基本信息
*/
func (this *Users) Add(m *Users) (int64, error) {
	defer db.Close()
	m.Password = utils.MD5(m.Password)
	return db.Save(m)
}

//更新
func (this *Users) Update(m *Users) (int64, error) {
	defer db.Close()
	type users struct {
		Id        int64
		Loginname string
		Relname   string
		Updator   int64
		Updated   int64
		Ip        string
	}
	u := new(users)
	u.Id = m.Id
	u.Loginname = m.Loginname
	u.Relname = m.Relname
	u.Updated = m.Updated
	u.Updator = m.Updator
	u.Ip = m.Ip

	return db.WhereEqual("id", u.Id).Update(u)
}

//更新profile
func (this *Users) UpdateProfile(m *Users) (int64, error) {
	defer db.Close()
	type users struct {
		Relname string
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(users)
	u.Relname = m.Relname
	u.Updated = m.Updated
	u.Updator = m.Updator
	u.Ip = m.Ip

	return db.WhereEqual("id", m.Id).Update(u)
}

// //缓存测试
// func (this *Users) GetByCache(key string) (val interface{}) {
// 	if val = bm.Get(key); val != nil {
// 		return
// 	}

// 	val = time.Now().Second()
// 	bm.Put(key, val, expired)
// 	return
// }

//获取一个可用账户，屏蔽禁用或已删除的账户
func (this *Users) Get(id int64) (*Users, error) {
	defer db.Close()
	m := new(Users)

	condition := qbs.NewEqualCondition("deleted", utils.DelNormal).AndEqual("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//获取一个账户
func (this *Users) GetEx(id int64) (*Users, error) {
	defer db.Close()
	m := new(Users)
	//err := db.WhereEqual("id", id).Find(m)
	condition := qbs.NewEqualCondition("id", id)
	err := db.Condition(condition).Find(m)
	return m, err
}

//分页获取账户
func (this *Users) GetAll() (us []*Users, err error) {
	defer db.Close()

	condition := qbs.NewEqualCondition("Deleted", utils.DelNormal)
	err = db.Condition(condition).OmitFields("Password", "Updator", "Ip").FindAll(&us)

	return
}

//分页获取账户
func (this *Users) GetAllEx() (us []*Users, err error) {
	defer db.Close()

	err = db.OmitFields("Password", "Updator", "Ip").FindAll(&us)
	return
}

/*
重置账户
禁用、启用账户
*/
func (this *Users) Reset(id int64, f *Field) (int64, error) {
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
	type users struct {
		Status  int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(users)
	u.Status = status
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereEqual("id", m.Id).Update(u)
}

/*
删除账户
*/
func (this *Users) Delete(ids []string, f *Field) (int64, error) {
	defer db.Close()
	type users struct {
		Deleted int8
		Updator int64
		Updated int64
		Ip      string
	}
	u := new(users)
	u.Deleted = utils.DelDeleted
	u.Updated = f.Updated
	u.Updator = f.Updator
	u.Ip = f.Ip

	return db.WhereIn("id", qbs.StringsToInterfaces(strings.Join(ids, ","))).Update(u)
}

/*
是否存在同名账户
*/
func (this *Users) Exist(name string) bool {
	defer db.Close()
	condition := qbs.NewEqualCondition("loginname", name)
	num := db.Condition(condition).Count("users")
	return num > 0
}

//
func (this *Users) XXX() string {
	return "id"
}
