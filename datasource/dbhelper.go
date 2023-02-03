package datasource

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"sync"

	"xorm.io/xorm"
)

// 数据库单例 dbhelper.go

var (
	masterEngine *xorm.Engine

	// slaveEngine *xorm.Engine

	lock sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	///锁定
	lock.Lock()
	//解锁
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	///配置信息
	c := MasterDbConf
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",

		c.User, c.Pwd, c.Host, c.Port, c.DbName)

	masterEngine, err := xorm.NewEngine(DriverName, driverSource)

	if err != nil {

		log.Fatal("dbhelper.DbInstanceMaster err:", err)

		return nil

	}

	// DEBUG 模式，打印全部sql, 帮助对比 orm 与 sql 执行对照关系

	masterEngine.ShowSQL(false)

	return masterEngine
}
