package controllerscrms

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"go_mars/lib/pagination"
	"html/template"
	"strconv"
)

func UserIndex(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageindex = 1
	}
	pagesize := 30
	var users []models.User
	var totalCount int64
	models.DB.Model(&models.User{}).Count(&totalCount)
	//创建一个分页器
	pagination := pagination.Initialize(c.Request, totalCount, pagesize)
	models.DB.Debug().Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&users)
	c.HTML(200, "users/index.html", gin.H{
		"users":    users,
		"paginate": template.HTML(pagination.Pages()),
	})
}
