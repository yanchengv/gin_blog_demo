package controllerscrms

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"go_mars/lib/pagination"
	"html/template"
	"net/http"
	"strconv"
)

func ArticleIndex(c *gin.Context) {

	pageindex, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageindex = 1
	}
	pagesize := 1
	var articles []models.Article
	var totalCount int64
	models.DB.Model(&models.Article{}).Count(&totalCount)
	models.DB.Debug().Offset((pageindex - 1) * pagesize).Limit(pagesize).Order("created_at asc").Find(&articles)
	//创建一个分页器，每页30条
	pagination := pagination.Initialize(c.Request, totalCount, pagesize)
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
