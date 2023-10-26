package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type User struct {
	Id       int64    `pk:"auto" json:"id,omitempty" orm:""`
	Username string `json:"userName"`
	Password string `json:"password"`
	Phone    string `json:"phone,omitempty" orm:"null"`
	Mail     string `json:"mail,omitempty" orm:"null"`
	Reg_time     time.Time `json:"Reg_time,omitempty"`
	Real_name    string    `json:"Real_name,omitempty" orm:"null"`
	Permission   int    `json:"permission,omitempty" orm:"null"`
	Projectid  int64        `json:"projectid";orm:"rel(one)"`
}

func GetUserInfobyId(userid int64) (u *User, ret int) {
	dbObj := db.New()
	u = &User{}
	err := dbObj.Where("Id=?", userid).First(u).Error
	logs.Info("myerr", err)
	if err == nil {
		return u, RESULT_OK
	} else if err == gorm.ErrRecordNotFound {
		return nil, RESULT_OK
	} else {
		return nil, RESULT_ERROR
	}
}

func Login(username string, pwd string) (*User, int) {
	var user = User{}
	o := db.New()
	err := o.Where("username=?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, RESULT_USER_NOT_EXIST
	}

	if strings.Compare(user.Password, pwd) == 0 {
		return &user, RESULT_OK
	} else {
		return nil, RESULT_PASSWORD_ERROR
	}
}