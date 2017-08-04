package database

import (
	"reflect"
	"github.com/go-xorm/xorm"
)

var database interface{}

/**
 * 实例化数据库引擎
 * param(数据库类型[, 数据库主机用户名, 数据库主机密码, 数据库主机URL, 数据库端口, 数据库])
 */
func New(dbType interface{}, args ...interface{}) {
	switch dbType {
	case "mysql":
		database = new(MySql)
		reflect.ValueOf(database).MethodByName("New").Call([]reflect.Value{reflect.ValueOf(args[0]), reflect.ValueOf(args[1]), reflect.ValueOf(args[2]), reflect.ValueOf(args[3]), reflect.ValueOf(args[4])})
	}
}

/**
 * 获取数据库引擎
 */
func Get() *xorm.Engine {
	engine := reflect.ValueOf(database).MethodByName("Get").Call([]reflect.Value{})[0]
	return engine.Interface().(*xorm.Engine)
}
