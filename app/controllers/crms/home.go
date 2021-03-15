package controllerscrms

import "github.com/gin-gonic/gin"

func HomeIndex(c *gin.Context){
	c.HTML(200,"homes/index.html",gin.H{
		"title": "主页",
	})

}
