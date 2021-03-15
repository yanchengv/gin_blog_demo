package main

import (
	"go_mars/app/models"
	"go_mars/routes"

)



func main() {
	models.InitDB()
	r := adminroutes.SetupRoute()
	r.Run(":9000")

}
