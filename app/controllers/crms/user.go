package controllerscrms

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"go_mars/lib/pagination"
	"html/template"
	"net/http"
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

func UserNew(c *gin.Context) {
	c.HTML(http.StatusOK, "users/new.html", gin.H{})
}

func UserEdit(c *gin.Context) {
	var user models.User
	id, ok := c.GetQuery("id")
	if !ok {
		c.Redirect(http.StatusSeeOther, "/crms/users")
		return
	}
	models.DB.Where("id =?", id).First(&user)
	c.HTML(http.StatusOK, "users/edit.html", gin.H{
		"user": user,
	})
}

func UserCreate(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := models.User{Name: name, Email: email, Password: password}
	models.DB.Debug().Create(&user)
	c.Redirect(http.StatusSeeOther, "/crms/users")
}

func UserUpdate(c *gin.Context) {
	models.DB.Where("id = ?", c.PostForm("id")).Updates(models.User{Name: c.PostForm("name"), Email: c.PostForm("email"), Password: c.PostForm("password")})
	c.Redirect(http.StatusFound, "/crms/users")

}

func UserDestroy(c *gin.Context) {
	models.DB.Where("id =?", c.PostForm("id")).Delete(models.User{})
	c.Redirect(http.StatusFound, "/crms/users")
}

func GetUserInfo(c *gin.Context) {
	var user models.User
	session := sessions.Default(c)
	models.DB.Where("id = ?", session.Get("currentUID")).First(&user)
	c.JSON(http.StatusOK, gin.H{
		"userID":   user.ID,
		"userName": user.Name,
	})
}
