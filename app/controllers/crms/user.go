package controllerscrms

import (
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
)

func UserIndex(c *gin.Context){
	pageindex := 1
	pagesize  := 5
	var users []models.User
	models.DB.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&users)
	c.HTML(200,"users/index.html",gin.H{
		"users": users,
	})
}


