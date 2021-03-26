package controllerapi

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"net/http"
	"strconv"
)

func ArticleIndex(c *gin.Context) {
	var totalCount int64
	pagesize := 30
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	var articles []models.Article
	arr := make([]map[string]interface{}, 0)
	models.DB.Model(&models.Article{}).Count(&totalCount)
	models.DB.Offset((page - 1) * pagesize).Limit(pagesize).Order("id desc").Find(&articles)
	for _, article := range articles {
		createdAt := article.CreatedAt.Format("2006-01-02 15:04:05")
		articleMap := map[string]interface{}{"id": article.ID, "title": article.Title, "subtitle": article.Subtitle, "createdAt": createdAt, "content": article.Content}
		arr = append(arr, articleMap)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "成功",
		"data":   arr,
	})
}
