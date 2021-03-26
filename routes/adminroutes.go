//后台路由
package myroutes

import (
	"github.com/gin-gonic/gin"
	controllerscrms "go_mars/app/controllers/crms"
)

func AdminRoutes(r *gin.Engine) {
	crms := r.Group("/crms")
	crms.GET("/login", controllerscrms.LoginIndex)
	crms.GET("/logout", controllerscrms.Logout)
	crms.Use(controllerscrms.AuthUser())
	{
		crms.GET("/homes", controllerscrms.HomeIndex)

		crms.POST("/login", controllerscrms.Login)
		users := crms.Group("/users")
		{
			users.GET("", controllerscrms.UserIndex)
			users.GET("/new", controllerscrms.UserNew)
			users.GET("/edit", controllerscrms.UserEdit)
			users.POST("/create", controllerscrms.UserCreate)
			users.POST("/update", controllerscrms.UserUpdate)
			users.POST("/destroy", controllerscrms.UserDestroy)
			users.GET("/getUserInfo", controllerscrms.GetUserInfo)
		}
		articles := crms.Group("/articles")
		{
			articles.GET("", controllerscrms.ArticleIndex)
			articles.GET("/new", controllerscrms.ArticleNew)
			articles.GET("/edit", controllerscrms.ArticleEdit)
			articles.POST("/create", controllerscrms.ArticleCreate)
			articles.POST("/update", controllerscrms.ArticleUpdate)
			articles.POST("/destroy", controllerscrms.ArticleDestroy)
		}

	}
}
