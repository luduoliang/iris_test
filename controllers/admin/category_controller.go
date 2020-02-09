package admin

import (
	"github.com/kataras/iris"
	"strconv"
	"iris_test/models"
	"iris_test/library"
	"iris_test/common"
)

//列表
func CategoryList(c iris.Context){
	page_temp := c.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetCategoryListPaginator(page)
	paginator := library.NewPaginator(c.Request(), models.PAGE_LIMIT, nums)

	c.ViewData("list", list)
	c.ViewData("Page", paginator)
	c.View("admin/category/list.html")
}

//添加
func CategoryAdd(c iris.Context) {
	if(c.Method() == "POST"){
		name := c.FormValue("name")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)
		info := models.Category{}
		info.Name = name
		info.Weight = weight
		info = models.CreateCategory(info)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}else{
		c.View("admin/category/add.html")
	}

}

//编辑
func CategoryEdit(c iris.Context) {
	id_temp := c.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetCategory(id)
	if(c.Method() == "POST"){
		name := c.FormValue("name")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info.Name = name
		info.Weight = weight
		info = models.UpdateCategory(info)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}else{
		c.ViewData("info", info)
		c.View("admin/category/edit.html")
	}
}


//删除
func CategoryDelete(c iris.Context) {
	if(c.Method() == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)

		info := models.DeleteCategory(id)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}
}

