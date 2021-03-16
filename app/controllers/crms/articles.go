package controllerscrms

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"net/http"
)

func ArticleIndex(c *gin.Context) {
	pageindex := 1
	pagesize := 30
	var articles []models.Article
	models.DB.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&articles)
	c.HTML(http.StatusOK, "articles/index.html", gin.H{
		"articles": articles,
	})
}
