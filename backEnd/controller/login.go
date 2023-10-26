package controller

import (
	"data/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c LoginController)Login() {
	username := c.GetString("username")
	pwd := c.GetString("password")
	if pwd == "" || username == "" {
		c.Data["json"] = models.GetAnswerByErrorCode(models.RESULT_PARAM_ERROR)
		c.ServeJSON()
		return
	}

	user, ret := models.Login(username, pwd)
	if ret != models.RESULT_OK {
		c.Data["json"] = models.GetAnswerByErrorCode(ret)
		c.ServeJSON()
		return
	}

	jwt, ret := models.GetJWT(user)
	if ret != models.RESULT_OK {
		c.Data["json"] = models.GetAnswerByErrorCode(ret)
		c.ServeJSON()
		return
	}
	
	c.Data["json"] = models.GetAnswer(ret, jwt)
	c.ServeJSON()
	return
}
