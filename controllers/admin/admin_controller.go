package admin

import (
	"github.com/kataras/iris"
	"strconv"
	"iris_test/models"
	"iris_test/library"
	"iris_test/common"
)

//列表
func AdminList(c iris.Context){
	page_temp := c.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetAdminListPaginator(page)
	paginator := library.NewPaginator(c.Request(), models.PAGE_LIMIT, nums)

	c.ViewData("list", list)
	c.ViewData("Page", paginator)
	c.View("admin/admin/list.html")
}

//添加
func AdminAdd(c iris.Context) {
	if(c.Method() == "POST"){
		username := c.FormValue("username")
		password := c.FormValue("password")

		info := models.Admin{}
		info.Username = username
		info.Password = password
		info = models.CreateAdmin(info)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}else{
		c.View("admin/admin/add.html")
	}

}

//编辑
func AdminEdit(c iris.Context) {
	id_temp := c.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetAdmin(id)
	if(c.Method() == "POST"){
		username := c.FormValue("username")
		password := c.FormValue("password")

		info.Username = username
		info.Password = password
		info = models.UpdateAdmin(info)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}else{
		c.ViewData("info", info)
		c.View("admin/admin/edit.html")
	}
}


//删除
func AdminDelete(c iris.Context) {
	if(c.Method() == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)

		info := models.DeleteAdmin(id)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}
}

//禁用/恢复
func AdminForbid(c iris.Context) {
	if(c.Method() == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)
		info := models.GetAdmin(id)

		if(info.Status == 1){
			info.Status = 2
		}else if(info.Status == 2){
			info.Status = 1
		}

		info = models.UpdateAdminInfo(info)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}
}

