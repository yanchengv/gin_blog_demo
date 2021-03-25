package controllerscrms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mars/app/models"
	"net/http"
	"strconv"
	"time"
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
	//生成cookie
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 30) //设置cookie的过期时间
	//cookie信息
	cookie := http.Cookie{Name: "userID", Value: userID, Expires: expiration}
	http.SetCookie(c.Writer, &cookie)
	c.Redirect(http.StatusSeeOther, "/crms/homes")
}
