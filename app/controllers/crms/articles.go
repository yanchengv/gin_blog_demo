package controllerscrms

import "C"
import (
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"net/http"
)

func ArticleIndex(c *gin.Context) {
	pageindex := 1
	pagesize := 30
	type pagination struct {
		CurrentPage int
		TotalPage   int
		PageSize    int
	}
	page := pagination{CurrentPage: pageindex, TotalPage: 10, PageSize: 30}
	var articles []models.Article
	models.DB.Offset((page.CurrentPage - 1) * pagesize).Limit(pagesize).Find(&articles)
	c.HTML(http.StatusOK, "articles/index.html", gin.H{
		"articles": articles,
		"page":     page,
	})
}

func ArticleNew(c *gin.Context) {
	isAlert := false
	alertMsg, ok := c.GetQuery("alertMsg")
	if ok {
		isAlert = true
	}
	c.HTML(http.StatusOK, "articles/new.html", gin.H{
		"isAlert":  isAlert,
		"alertMsg": alertMsg,
	})
}

func ArticleCreate(c *gin.Context) {
	content := c.PostForm("content")
	title := c.PostForm("title")
	subtitle := c.PostForm("subtitle")
	article := models.Article{Title: title, Subtitle: subtitle, Content: content}
	res := models.DB.Create(&article)
	if res.Error != nil {

	}
	c.Redirect(http.StatusMovedPermanently, "/crms/articles/new?isAlert=true&alertMsg=创建成功")
	//c.HTML(http.StatusOK, "articles/new.html", gin.H{
	//	"isAlert": true,
	//	"alertMsg": "创建成功",
	//})
}
