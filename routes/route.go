package myroutes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	controllerscrms "go_mars/app/controllers/crms"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	//配置置cookie和session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mars_session", store))
	//自定义模板继承方法
	r.HTMLRender = controllerscrms.LoadTemplates("app/views")
	//设置静态文件的路由
	r.Static("/assets", "assets")

	//路由拆分
	//后台路由
	AdminRoutes(r)
	//api的路由
	ApiRoutes(r)

	return r

}
