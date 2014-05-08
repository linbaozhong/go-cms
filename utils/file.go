package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//当前应用的绝对路径
func GetAppRoot() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return filepath.Dir(p)
}

//合并应用文件路径
func MergePath(args ...string) string {
	for i, e := range args {
		if e != "" {
			return filepath.Join(AppRoot, filepath.Clean(strings.Join(args[i:], string(filepath.Separator))))
		}
	}
	return AppRoot
}

//Sqlite3路径
func Sqlite3Path(p string) string {
	s := MergePath(p)
	return strings.Replace(s, "\\", "/", -1)
}

//检查文件夹是否存在，如果不存在，则创建一个新文件夹
func GetDir(path string) error {
	//文件夹是否存在
	if DirExists(path) {
		return nil
	} else {
		//创建文件夹
		if err := os.Mkdir(path, os.ModeDir); err != nil {
			return err
		}
		return nil
	}
}

//移动文件或文件夹
func MoveFilePath(dst, src string) error {
	return nil
}

//移除文件或文件夹
