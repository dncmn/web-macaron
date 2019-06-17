package dao

import (
	"code.dncmn.io/web-macaron/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var (
	engine *xorm.Engine
)

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/game?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}

	engine.SetMaxIdleConns(100)
	engine.SetMaxOpenConns(100)
	engine.SetConnMaxLifetime(100)

	engine.SetLogger(engine.Logger())
	engine.ShowSQL(true)

	if os.Getenv("MIGRATE_DB") == "true" {
		err = engine.Sync2(new(model.ConfigDept))
	}

	//engine.SetLogger()
}

func GetDB() *xorm.Engine {
	return engine
}
