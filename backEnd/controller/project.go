package controller

import (
	"data/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ProjectController struct{
	beego.Controller
}

func (c ProjectController) ProjectDetail() {
	if(!c.IsUserLoggedIn()) {
		return 
	}

	// GetInt64 returns input value as int64 or the default value while it's present and input is blank.
	id,err:=c.GetInt64("id")
	if err!=nil{
		c.Data["json"] = models.GetAnswerByErrorCode(models.RESULT_PARAM_ERROR)
		c.ServeJSON()
		return
	}

	projInfo,ret:=models.GetProjectInfo(id)
	if ret!= models.RESULT_OK {
		c.Data["json"]= models.GetAnswerByErrorCode(ret)
		c.ServeJSON()
		return
	}

	c.Data["json"]= models.GetAnswer(models.RESULT_OK,projInfo)
	c.ServeJSON()
}

func (c ProjectController) ProjectList(){
	if(!c.IsUserLoggedIn()) {
		return 
	}

	projList,ret:=models.GetProjectList()
	if ret!= models.RESULT_OK {
		c.Data["json"]= models.GetAnswerByErrorCode(ret)
		c.ServeJSON()
		return
	}

	c.Data["json"]= models.GetAnswer(models.RESULT_OK,projList)
	c.ServeJSON()
}

func (c ProjectController) AddProject(){
	if(!c.IsUserLoggedIn()) {
		return 
	}

	var prj models.Project
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by &prj. If &prj is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError. 
	err:=json.Unmarshal(c.Ctx.Input.RequestBody, &prj)
	if err!=nil{
		c.Data["json"] = models.GetAnswerByErrorCode(models.RESULT_PARAM_ERROR)
		c.ServeJSON()
		return
	}

	ret:=models.AddProject(prj)
	if ret!= models.RESULT_OK {
		c.Data["json"]= models.GetAnswerByErrorCode(ret)
		c.ServeJSON()
		return
	}

	c.Data["json"]= models.GetAnswer(models.RESULT_OK,nil)
	c.ServeJSON()
}

func (c ProjectController) DeleteProject(){
	if(!c.IsUserLoggedIn()) {
		return 
	}

	id,err:=c.GetInt64("id")
	if err!=nil{
		c.Data["json"] = models.GetAnswerByErrorCode(models.RESULT_PARAM_ERROR)
		c.ServeJSON()
		return
	}

	ret:=models.DeleteProject(id)
	if ret!= models.RESULT_OK {
		c.Data["json"]= models.GetAnswerByErrorCode(ret)
		c.ServeJSON()
		return
	}

	c.Data["json"]= models.GetAnswer(models.RESULT_OK,nil)
	c.ServeJSON()
}

func (c ProjectController) IsUserLoggedIn() bool {
	user:=c.Ctx.Input.GetData("userinfo").(*models.User)
	if user==nil{
		c.Data["json"] = models.GetAnswerByErrorCode(models.RESULT_NOT_LOGIN)
		c.ServeJSON()
		return false
	}

	return true
}
