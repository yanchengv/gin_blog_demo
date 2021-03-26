package controllerscrms

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//判断用户是否登录
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		currentUserId := session.Get("currentUID")
		if currentUserId == nil {
			//获取cookie报错或者userId为空则跳转登录界面
			fmt.Println("请先登录")
			c.Redirect(http.StatusSeeOther, "/crms/login")
		}
		c.Next()
	}

}
