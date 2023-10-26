package conf
import (
	"github.com/astaxie/beego"

)

type AppConfig struct {
	MysqlConfig
}

type MysqlConfig struct {
	DBname     string
	Username   string
	MSPassword string
	Endpoint   string
}

func NewAppConfig() AppConfig {
	appconfig := AppConfig{}
	appconfig.MysqlConfig.DBname = beego.AppConfig.String("dbname")
	appconfig.MysqlConfig.Username = beego.AppConfig.String("dbuser")
	appconfig.MysqlConfig.MSPassword = beego.AppConfig.String("dbpwd")
	appconfig.MysqlConfig.Endpoint = beego.AppConfig.String("dbendpoint")
	return appconfig
}

func init() {
}
