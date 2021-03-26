package controllerscrms

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(200, "homes/index.html", gin.H{
		"currentUName": session.Get("currentUName"),
		"title":        "主页",
	})
}
