package models

import (
	"cms/utils"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/coocood/qbs"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

//页面公共信息
type Page struct {
	SiteName    string //网站名称
	Title       string //页面标题
	Company     string //公司名称
	Domain      string //域名
	Copyright   string //版权
	Keywords    string //Seo关键词
	Description string //Seo描述
	Author      string //作者
}

//公共字段
type Field struct {
	Sequence int
	Status   int8
	Deleted  int8
	Updator  int64
	Updated  int64
	Ip       string
	Name     string
	Role     string
}

//分页
type Pagination struct {
	Count int
	Prev  int
	Index int
	Next  int
	Size  int
}

//列表选项
type SelectItem struct {
	Key      string
	Value    string
	Selected bool
}

//上传文件
type UploadFile struct {
	Name string
	Ext  string
	Path string
	Size int64
}

var (
	db *qbs.Qbs
	////内存cache
	//bm      *cache.MemoryCache
	//expired int64
)

func Init() {
	//qbs.RegisterSqlite3("e:/mygo/src/cms/data/orange.db")
	qbs.RegisterSqlite3(utils.Sqlite3Path(beego.AppConfig.String("DatabasePath")))
	db, _ = qbs.GetQbs()
	// //cache期限
	// var err error
	// if expired, err = beego.AppConfig.Int64("CacheExpired"); err != nil {
	// 	expired = 60
	// }
	// bm = cache.NewMemoryCache()
}

func CreateTable(structPtr interface{}) error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}
	defer migration.Close()
	return migration.CreateTableIfNotExists(structPtr)
}

//DDL错误字符串
func GetDbError(err error) string {
	switch err {
	case sql.ErrNoRows:
		return "notFound"
	case sql.ErrTxDone:
		return "errTxDone"
	default:
		return "unknown"
	}
}

//扩展structptr
func Extend(dst interface{}, f *Field) {
	d := reflect.Indirect(reflect.ValueOf(dst))
	if v := d.FieldByName("Updator"); v.IsValid() && v.Int() == 0 {
		v.SetInt(f.Updator)
	}
	if v := d.FieldByName("Updated"); v.IsValid() && v.Int() == 0 {
		v.SetInt(f.Updated)
	}
	if v := d.FieldByName("Creator"); v.IsValid() && v.Int() == 0 {
		v.SetInt(f.Updator)
	}
	if v := d.FieldByName("Created"); v.IsValid() && v.Int() == 0 {
		v.SetInt(f.Updated)
	}
	if v := d.FieldByName("Ip"); v.IsValid() && v.String() == "" {
		v.SetString(f.Ip)
	}
}

//
func StringsToPtrs(s []string) (ids []interface{}) {
	for _, v := range s {
		ids = append(ids, v)
	}
	fmt.Println(ids)
	return
}
