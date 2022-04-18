package datasource

import (
	. "GoLangABP/demo.Web.Host/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)

type Db struct {
	Conn *gorm.DB
}

func (d *Db) Connect() error {
	//var (
	//	dbType, dbName, user, pwd, host string
	//)

	conf := Configs.Database
	//dbType = conf.Type
	//dbName = conf.Name
	//user = conf.User
	//pwd = conf.Password
	//host = conf.Host
	fmt.Println(112211111111111111)
	dsn := conf.User + ":" + conf.Password + "@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return err
	}
	d.Conn = db

	log.Println("Connect Mysql Success")

	return nil
}

func (d *Db) DB() *gorm.DB {
	return d.Conn
}
