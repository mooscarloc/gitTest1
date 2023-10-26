package router

import (
	"data/controller"
	"data/models"
	"github.com/astaxie/beego"
)

func init(){
	beego.Router("/v1/login",&controller.LoginController{},"get:Login")

	beego.Router("/v1/project/detail",&controller.ProjectController{},"get:ProjectDetail")
	beego.Router("/v1/config/project/list",&controller.ProjectController{},"get:ProjectList")
	beego.Router("/v1/config/project/add",&controller.ProjectController{},"post:AddProject")
	beego.Router("/v1/config/project/delete",&controller.ProjectController{},"get:DeleteProject")

	beego.InsertFilter("/v1/project/*",beego.BeforeRouter, models.Check)
	beego.InsertFilter("/v1/config/*",beego.BeforeRouter, models.Check)
	beego.InsertFilter("/v1/projects",beego.BeforeRouter, models.Check)
	beego.InsertFilter("/v1/data/*",beego.BeforeRouter, models.Check)
}
