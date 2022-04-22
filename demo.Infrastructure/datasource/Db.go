package datasource

import (
	. "GoLangABP/demo.Web.Host/Conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Db struct {
	Conn *sqlx.DB
}

func (d *Db) Connect() error {
	conf := Configs.Database
	//dbType = Conf.Type
	//dbName = Conf.Name
	//User = Conf.User
	//pwd = Conf.Password
	//host = Conf.Host

	dsn := conf.User + ":" + conf.Password + "@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		return err
	}
	d.Conn = db
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	log.Println("Connect Mysql Success")

	return nil
}

func (d *Db) DB() *sqlx.DB {
	return d.Conn
}