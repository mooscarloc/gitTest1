package models
import (
	"data/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
func InitDB(){
	appconfig := conf.NewAppConfig()
	mysqluser := appconfig.MysqlConfig.Username
	mysqlpass := appconfig.MysqlConfig.MSPassword
	mysqlurl := appconfig.MysqlConfig.Endpoint
	mysqldb := appconfig.MysqlConfig.DBname
	var err error
	db,err=gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			mysqluser,
			mysqlpass,
			mysqlurl,
			mysqldb))

	if err!=nil{
		panic(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "placeholder_" + defaultTableName
	}

	// Disable table name's pluralization
	db.SingularTable(true)
	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed. It WON’T delete unused columns to protect your data.
	db.AutoMigrate(new(Project))
	db.AutoMigrate(new(Record))
	db.AutoMigrate(new(User))
}

