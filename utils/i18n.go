package utils

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var jsons config.ConfigContainer

func I18n() {
	c := config.JsonConfig{}
	var err error
	jsons, err = c.Parse(MergePath(beego.AppConfig.String("LangPath")))
	if err != nil {
		panic(err)
	}
}

func Tr(key string, args ...string) string {
	//语言
	local := "zh"
	if len(args) > 0 {
		local = args[0]
	}
	//按语言读取
	lang, _ := jsons.DIY(local)

	m := lang.(map[string]interface{})
	for k, v := range m {
		if k == key {
			return v.(string)
		}
	}
	return ""
}
