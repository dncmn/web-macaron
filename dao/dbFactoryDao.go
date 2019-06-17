package dao

import (
	"code.dncmn.io/web-macaron/config"
	"code.dncmn.io/web-macaron/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var (
	engine *xorm.Engine
	dns    = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Dbname,
	)
)

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		fmt.Println(err)
		return
	}

	engine.SetMaxIdleConns(config.Config.Mysql.GOMaxIdleConns)
	engine.SetMaxOpenConns(config.Config.Mysql.MaxOpenConns)
	engine.SetConnMaxLifetime(config.Config.Mysql.ConnMaxLifetime)

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
