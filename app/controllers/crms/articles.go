package controllerscrms

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"go_mars/lib/pagination"
	"html/template"
	"net/http"
)

func ArticleIndex(c *gin.Context) {
	pageindex := 1
	pagesize := 30
	var articles []models.Article
	models.DB.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&articles)
	//创建一个分页器，一万条数据，每页30条
	pagination := pagination.Initialize(c.Request, 100, 10)
	//传到模板中需要转换成template.HTML类型，否则html代码会被转义
	c.HTML(http.StatusOK, "articles/index.html", gin.H{
		"articles": articles,
		"paginate": template.HTML(pagination.Pages()),
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
