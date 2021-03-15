package controllerscrms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//判断用户是否登录
func AuthUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId,err := c.Cookie("userID")
		if (err != nil || userId == "") {
			//获取cookie报错或者userId为空则跳转登录界面
			fmt.Println("请先登录")
			c.Redirect(http.StatusMovedPermanently,"/crms/login")
		}
		//c.Next()
	}

}