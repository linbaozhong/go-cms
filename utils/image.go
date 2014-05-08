/*
使用方法：
	im := new(utils.Image)
	//缩略图
	im.ToThumbnail("upload/images/0/1.jpg")
	//展示图
	im.ToView("upload/images/0/1.jpg")
*/
package utils

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/disintegration/imaging"
	"image"
	"path/filepath"
	"regexp"
)

type Image struct {
	ImagePath      string
	UserPath       string
	ThumbnailWidth int
	ViewWidth      int
	filename       string
	source         image.Image
}

/*
缩略图
filename:源文件名
path:用户文件夹，一般是用户id
*/
func (this *Image) ToThumbnail(filename, path string) (string, error) {
	this.UserPath = this.repath(path)
	return this.resize(filename, this.thumbnailWidth(), 0, 1)
}

/*
展示图
filename:源文件名
path:用户文件夹，一般是用户id
*/
func (this *Image) ToView(filename, path string) (string, error) {
	this.UserPath = this.repath(path)
	return this.resize(filename, this.viewWidth(), 0, 2)
}

/*
清理路径中的非法字符
*/
func (this *Image) repath(path string) string {
	r := regexp.MustCompile("[\\s]+")
	return r.ReplaceAllString(path, "")
}

//打开图片文件
func (this *Image) open(filename string) error {

	src, err := imaging.Open(filename)
	if err != nil {
		return err
	}
	this.filename = filepath.Base(filename)
	this.source = src
	return nil
}

//图片库路径
func (this *Image) imagePath() string {
	if this.ImagePath == "" {
		this.ImagePath = beego.AppConfig.String("UploadPath")
	}
	return this.ImagePath
}

//缩略图宽度
func (this *Image) thumbnailWidth() int {
	if this.ThumbnailWidth == 0 {
		w, _ := beego.AppConfig.Int("ThumbnailWidth")
		this.ThumbnailWidth = w
	}
	return this.ThumbnailWidth
}

//展示图宽度
func (this *Image) viewWidth() int {
	if this.ViewWidth == 0 {
		w, _ := beego.AppConfig.Int("ViewWidth")
		this.ViewWidth = w
	}
	return this.ViewWidth
}

//目标路径
func (this *Image) dstFilename(t int) (string, error) {
	//用户路径是否存在
	if this.UserPath == "" {
		this.UserPath = "0"
	}
	//文件夹
	path := filepath.Join(this.imagePath(), Int2str(t))

	if err := GetDir(path); err != nil {
		return "", err
	}
	//返回目标文件全路径名
	return filepath.Join(path, this.filename), nil
}

//缩放
func (this *Image) resize(filename string, w, h, s int) (string, error) {
	if this.source == nil {
		if err := this.open(filename); err != nil {
			return "", err
		}
	}

	var dst *image.NRGBA
	dst = imaging.Resize(this.source, w, h, imaging.CatmullRom)
	//取目标文件名
	dstfile, err := this.dstFilename(s)
	if err != nil {
		return "", err
	}
	//保存文件
	return dstfile, imaging.Save(dst, dstfile)
}
