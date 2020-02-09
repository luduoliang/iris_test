package router

import (
	"github.com/kataras/iris"
	"iris_test/controllers/admin"
	"iris_test/middleware"
	"iris_test/controllers/api"
)

func InitRouter(app *iris.Application){
	//后台路由
	//登录后台
	app.Any("/admin/login", admin.Login)
	//验证码
	app.Any("/captcha", admin.Captcha)
	//退出登录
	app.Any("/admin/logout", admin.Logout)
	//后台路由分组，使用检查是否登录CheckAdminLogin中间间
	admin_router := app.Party("/admin", middleware.CheckAdminLogin()).AllowMethods(iris.MethodOptions)
	{
		//后台首页
		admin_router.Any("/", admin.Index)
		//后台首页信息
		admin_router.Any("/info", admin.Info)
		//上传图片
		admin_router.Any("/file/upload", admin.UploadFile)

		//系统用户
		admin_router.Any("/admin/list", admin.AdminList)
		admin_router.Any("/admin/add", admin.AdminAdd)
		admin_router.Any("/admin/edit", admin.AdminEdit)
		admin_router.Any("/admin/delete", admin.AdminDelete)
		admin_router.Any("/admin/forbid", admin.AdminForbid)
		//资讯分类
		admin_router.Any("/category/list", admin.CategoryList)
		admin_router.Any("/category/add", admin.CategoryAdd)
		admin_router.Any("/category/edit", admin.CategoryEdit)
		admin_router.Any("/category/delete", admin.CategoryDelete)
		//资讯管理
		admin_router.Any("/news/list", admin.NewsList)
		admin_router.Any("/news/add", admin.NewsAdd)
		admin_router.Any("/news/edit", admin.NewsEdit)
		admin_router.Any("/news/delete", admin.NewsDelete)



	}

	//接口路由分组，使用跨域中间间
	api_router := app.Party("/api", middleware.Cors()).AllowMethods(iris.MethodOptions)
	{
		//资讯分类列表，参数：
		api_router.Any("/category/list", api.CategoryList)
		//资讯列表，参数：page,per_page,category_id(可传)
		api_router.Any("/news/list", api.NewsList)
		//资讯详情，参数：参数：id
		api_router.Any("/news/detail", api.NewsDetail)



	}



}