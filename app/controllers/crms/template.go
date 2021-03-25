package controllerscrms

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"path/filepath"
	"strings"
)

//Gin框架默认都是使用单模板，如果需要使用block template功能（模板继承），可以通过"github.com/gin-contrib/multitemplate"库实现
func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	//加载layout模板
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}
	//继承layout子模板
	includes, err := filepath.Glob(templatesDir + "/crms/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	//生成以application.html为模板的页面
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		//为每个页面模板定义名字，默认为去掉[app/views/crms/]后剩余的目录加文件的名字
		//例如 homes/index.html，在route中直接使用即可
		include = strings.Replace(include, "app/views/crms/", "", 1)
		fmt.Println(include)
		r.AddFromFiles(include, files...)
	}

	//单独定义登陆页面的模板
	r.AddFromFiles("logins/signin.html", "app/views/logins/signin.html")
	return r
}
