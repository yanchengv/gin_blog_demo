package main

import (
	"go_mars/app/models"
	"go_mars/routes"
)

func main() {
	models.InitDB()
	r := myroutes.SetupRoute()
	r.Run(":9000")

}
