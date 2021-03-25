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
	pagesize := 30
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
	c.Redirect(http.StatusSeeOther, "/crms/articles/new?isAlert=true&alertMsg=创建成功")
	//c.HTML(http.StatusOK, "articles/new.html", gin.H{
	//	"isAlert": true,
	//	"alertMsg": "创建成功",
	//})
}

func ArticleEdit(c *gin.Context) {
	var article models.Article
	id, _ := c.GetQuery("id")
	id1, _ := strconv.Atoi(id)
	models.DB.Where(&models.Article{ID: uint(id1)}).First(&article)
	c.HTML(http.StatusOK, "articles/edit.html", gin.H{
		"article": article,
	})
}

func ArticleUpdate(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	subtitle := c.PostForm("subtitle")
	content := c.PostForm("content")
	models.DB.Debug().Where("id = ?", id).Updates(models.Article{Title: title, Subtitle: subtitle, Content: content})
	url := "/crms/articles/edit?id=" + id + "isAlert=true&alertMsg=修改成功"
	c.Redirect(http.StatusSeeOther, url)

}

func ArticleDestroy(c *gin.Context) {
	id := c.PostForm("id")
	models.DB.Debug().Where("id = ? ", id).Delete(models.Article{})
	url := "/crms/articles?isAlert=true&alertMsg=删除成功"
	c.Redirect(http.StatusSeeOther, url)
}
