package datasource

func OpenSql() {
	// c := MasterDbConf
	// driverSource := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%d)/%s?charset=%s",
	// 	c.User, c.Pwd, c.Host, c.Port, c.DbName, c.DbCharset,
	// )

	// fmt.Println(driverSource)

	// dbEngine,error:=xorm.NewEngine(c.DbName,driverSource)
	// if error!=nil {
	// 	fmt.Println(error)
	// 	return
	// }

}

const DriverName = "mysql"

var MasterDbConf DbConf = DbConf{

	Host: "127.0.0.1",

	Port: 3306,

	User: "root",

	Pwd: "123456",

	DbName: "app_market",

	DbCharset: "utf8",
}

type DbConf struct {
	Host string

	Port int

	User string

	Pwd string

	DbName string

	DbCharset string
}
