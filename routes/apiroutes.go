//api接口路由
package myroutes

import (
	"github.com/gin-gonic/gin"
	controllerapi "go_mars/app/controllers/api"
)

func ApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	articles := api.Group("/articles")
	{
		articles.GET("", controllerapi.ArticleIndex)
	}
}
