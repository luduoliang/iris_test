package api

import (
	"github.com/kataras/iris"
	"iris_test/models"
	"iris_test/common"
)

//资讯分类列表，参数：
func CategoryList(c iris.Context) {
	list := models.GetAllCategoryList()
	c.JSON(common.JsonData(true, list, "操作成功"))
}