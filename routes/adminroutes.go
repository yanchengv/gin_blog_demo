package adminroutes

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/controllers/crms"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.HTMLRender = controllerscrms.LoadTemplates("app/views")
	r.Static("/assets", "assets")
	//后台路由
	crms := r.Group("/crms")
	crms.GET("/login", controllerscrms.LoginIndex)
	crms.Use(controllerscrms.AuthUser())
	{
		crms.GET("/homes", controllerscrms.HomeIndex)
		crms.GET("/users", controllerscrms.UserIndex)
		crms.POST("/login", controllerscrms.Login)
		articles := crms.Group("/articles")
		{
			articles.GET("", controllerscrms.ArticleIndex)
			articles.GET("/new", controllerscrms.ArticleNew)
			articles.POST("/create", controllerscrms.ArticleCreate)
		}

	}

	return r

}
