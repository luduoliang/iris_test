package admin

import (
	"github.com/kataras/iris"
	"strconv"
	"time"
	"os"
	"io"
	"iris_test/common"
)

//首页
func Index(c iris.Context){
	c.View("admin/index/index.html")
}

//首页信息
func Info(c iris.Context) {
	c.View("admin/index/info.html")
}

//上传文件
func UploadFile(c iris.Context) {
	file, info, form_err := c.FormFile("pic")
	defer file.Close()
	if form_err != nil {
		c.JSON(common.JsonData(false, "", form_err.Error()))
		return
	}

	fname := info.Filename

	year := strconv.Itoa(time.Now().Year())
	month := time.Now().Month().String()
	day := strconv.Itoa(time.Now().Day())
	//创建目录
	dirpath := "static/upload/"+year+"/"+month+"/"+day
	dir_err := os.MkdirAll(dirpath, os.ModePerm)
	if dir_err != nil {
		c.JSON(common.JsonData(false, "", dir_err.Error()))
		return
	}

	//上传文件相对路劲
	file_path := dirpath + "/" + fname

	out, open_err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()

	if open_err != nil {
		c.JSON(common.JsonData(false, "", open_err.Error()))
		return
	}


	_, copy_err := io.Copy(out, file)
	if(copy_err != nil){
		c.JSON(common.JsonData(false, "", copy_err.Error()))
		return
	}
	scheme := "http://"
	if c.Request().TLS != nil {
		scheme = "https://"
	}

	all_file_path := scheme + c.Host() + "/" + file_path
	upload_info := make(map[string]string)
	upload_info["pic_url"] = all_file_path

	c.JSON(common.JsonData(true, upload_info, "上传成功"))
}


