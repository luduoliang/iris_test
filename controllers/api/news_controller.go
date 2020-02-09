package api

import (
	"github.com/kataras/iris"
	"strconv"
	"iris_test/models"
	"iris_test/common"
)

//资讯列表，参数：page,per_page,category_id(可传)
func NewsList(c iris.Context) {
	page_temp := c.FormValue("page")
	page, _ := strconv.Atoi(page_temp)
	per_page_temp := c.FormValue("per_page")
	per_page, _ := strconv.Atoi(per_page_temp)
	category_id_temp := c.FormValue("category_id")
	category_id, _ := strconv.Atoi(category_id_temp)

	list := models.GetNewsList(category_id, page, per_page)
	c.JSON(common.JsonData(true, list, "操作成功"))
}

//资讯详情，参数：id
func NewsDetail(c iris.Context) {
	id_temp := c.FormValue("id")
	id, _ := strconv.Atoi(id_temp)

	info := models.GetNewsPreload(id)
	c.JSON(common.JsonData(true, info, "操作成功"))
}