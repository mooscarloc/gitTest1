package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Project struct {
	Id       int64   `pk:"auto" json:"id,omitempty" orm:"size(11)"`
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Industry int     `json:"industry" orm:"size(11)"`
	Mn       string  `json:"mn"`
	Type     string  `json:"-"`
}

var prjmap map[string]*Project
func GetProjectList() ([]*Project,int){
	o:=db.New()
	var projects []*Project
	err:=o.Find(&projects).Error
	if err!=nil{
		return nil, RESULT_ERROR
	}

	return projects, RESULT_OK
}

func GetProjectInfo(id int64) (*Project,int){
	o:=db.New()
	var projects Project
	err:=o.Where("id=?",id).First(&projects).Error
	if err!=nil{
		return nil, RESULT_ERROR
	}

	return &projects, RESULT_OK
}

func AddProject(prj Project) int{
	o:=db.New()
	prjmap[prj.Id]=&prj
	err=o.Create(&prj).Error
	if err!=nil{
		return RESULT_ERROR
	}

	return RESULT_OK
}

func DeleteProject(id int64) int{
	o:=db.New()
	err:=o.Delete(&Project{Id: id}).Error
	if err!=nil{
		return RESULT_ERROR
	}

	delete(prjmap, id)
	return RESULT_OK
}

func InitPrjMap(){
	prjmap = make(map[string]*Project)
	o:=db.New()
	var prjs []Project
	o.Find(&prjs)
	for _, prj := range prjs{
		prjmap[prj.Id]=&prj
	}
}
