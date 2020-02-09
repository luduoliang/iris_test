package middleware

import (
	"github.com/kataras/iris"
	"iris_test/common"
)

//验证用户是否登录
func CheckAdminLogin() iris.Handler{
	return func(c iris.Context) {
		admin_info := common.GetSession(c, "admin_info")
		if admin_info == nil {
			c.Redirect("/admin/login", iris.StatusFound)
			return
		}

		//登录信息分享到所有视图中
		c.ViewData("admin_info", admin_info)
		c.Next()
	}
}

