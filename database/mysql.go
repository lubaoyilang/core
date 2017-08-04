package database

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySql struct {
	engine *xorm.Engine
}

func (this *MySql) New(username, password, host, port, database string) {
	var err error
	this.engine, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8");
	if err != nil {
		log.Fatal(err)
	}
}

func (this *MySql) Get() *xorm.Engine {
	return this.engine
}

func (this *MySql) Close() {
	err := this.engine.Close()
	if err != nil {
		log.Fatal(err)
	}
}
