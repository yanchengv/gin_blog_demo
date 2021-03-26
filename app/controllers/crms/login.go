package controllerscrms

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"net/http"
	"strconv"
)

func LoginIndex(c *gin.Context) {
	c.HTML(200, "logins/signin.html", gin.H{})
}

func Login(c *gin.Context) {
	var user models.User
	//用户身份校验
	email := c.PostForm("email")
	password := c.PostForm("password")
	models.DB.Where("email = ?", email).First(&user)
	if user.Email != email {
		fmt.Println("用户不存在！")
		c.Redirect(http.StatusSeeOther, "/crms/login")
		return
	}
	if user.Password != password {
		fmt.Println("密码不正确！")
		c.Redirect(http.StatusSeeOther, "/crms/login")
		return
	}
	userID := strconv.Itoa(int(user.ID))
	//session信息
	session := sessions.Default(c)
	session.Set("currentUID", userID)
	session.Set("currentUName", user.Name)
	session.Save()
	c.Redirect(http.StatusSeeOther, "/crms/homes")
}

//退出
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/crms/login")
}
