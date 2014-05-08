package models

import (
	"cms/utils"
	"fmt"
	"strings"
)

type Current struct {
	Id   int64
	Name string
	Role string
}

//是否系统管理员
func IsSystemAdmin(role string) bool {
	role = fmt.Sprintf(",%s,", role)

	return strings.Contains(role, fmt.Sprintf(",%d,", utils.RoleSuper))
}

//是否域管理员
func IsPublisher(role string) bool {
	role = fmt.Sprintf(",%s,", role)

	return strings.Contains(role, fmt.Sprintf(",%d,", utils.RolePublisher))
}

//读取登录用户的Cookie信息
func GetCurrentUser(cookie string) (currentuser *Current) {
	currentuser = new(Current)

	cookie = utils.CookieDecode(cookie)

	//拆分cookie
	curr := strings.Split(cookie, "|")
	if len(curr) > 0 {
		currentuser.Id, _ = utils.Str2int64(curr[0]) //strconv.ParseInt(curr[0], 10, 0)
	}
	if len(curr) > 1 {
		currentuser.Name = curr[1]
	}
	if len(curr) > 2 {
		currentuser.Role = curr[2]
	}
	return
}
