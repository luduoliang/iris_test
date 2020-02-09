package admin

import (
	"github.com/kataras/iris"
	"strconv"
	"iris_test/models"
	"iris_test/library"
	"iris_test/common"
)

//列表
func NewsList(c iris.Context){
	page_temp := c.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetNewsListPaginator(page)
	paginator := library.NewPaginator(c.Request(), models.PAGE_LIMIT, nums)
	c.ViewData("list", list)
	c.ViewData("Page", paginator)
	c.View("admin/news/list.html")
}

//添加
func NewsAdd(c iris.Context) {
	if(c.Method() == "POST"){
		category_id_temp := c.FormValue("category_id")
		category_id, _ := strconv.Atoi(category_id_temp)
		title := c.FormValue("title")
		description := c.FormValue("description")
		pic := c.FormValue("pic")
		content := c.FormValue("content")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)


		info := models.News{}
		info.CategoryId = category_id
		info.Title = title
		info.Description = description
		info.Pic = pic
		info.Content = content
		info.Weight = weight

		info = models.CreateNews(info)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}else{
		category := models.GetAllCategoryList()
		c.ViewData("category", category)
		c.View("admin/news/add.html")
	}

}

//编辑
func NewsEdit(c iris.Context) {
	id_temp := c.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetNews(id)
	if(c.Method() == "POST"){
		category_id_temp := c.FormValue("category_id")
		category_id, _ := strconv.Atoi(category_id_temp)
		title := c.FormValue("title")
		description := c.FormValue("description")
		pic := c.FormValue("pic")
		content := c.FormValue("content")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info.CategoryId = category_id
		info.Title = title
		info.Description = description
		info.Pic = pic
		info.Content = content
		info.Weight = weight

		info = models.UpdateNews(info)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}else{
		category := models.GetAllCategoryList()
		c.ViewData("category", category)
		c.ViewData("info", info)
		c.View("admin/news/edit.html")
	}
}


//删除
func NewsDelete(c iris.Context) {
	if(c.Method() == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)

		info := models.DeleteNews(id)
		c.JSON(common.JsonData(true, info, "操作成功"))
	}
}
