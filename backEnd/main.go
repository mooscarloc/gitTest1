package main

import (
	"data/models"
	_ "data/router"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.InitDB()
	models.InitPrjMap()
}

func main() {
	beego.Run()
}

