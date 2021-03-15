package  adminroutes

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/controllers/crms"
)

func SetupRoute() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("app/views/**/**/*")
	r.Static("/assets","assets")
	//后台路由
	crms := r.Group("/crms")
	crms.GET("/login",controllerscrms.LoginIndex)
	crms.Use(controllerscrms.AuthUser())
	{
		crms.GET("/homes", controllerscrms.HomeIndex)
		crms.GET("/users", controllerscrms.UserIndex)
		crms.POST("/login",controllerscrms.Login)
	}

	return r

}

