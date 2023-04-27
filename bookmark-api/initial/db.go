package initial

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog"
	"os"
	"time"
)

const DbDriverName = "mysql"

func InitDb() {

	if err := orm.RegisterDriver(DbDriverName, orm.DRMySQL); err != nil {
		klog.Error(err.Error())
		os.Exit(1)
	}

	klog.Infof("[ orm ] orm driver is %s", DbDriverName)

	user, _ := config.String("mysql_user")
	password, _ := config.String("mysql_password")
	host, _ := config.String("mysql_host")
	port, _ := config.Int("mysql_port")
	db, _ := config.String("mysql_database")

	timeZone := config.DefaultString("mysql_time_zone", "Asia%2FShanghai")

	engineUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&&loc=%s",
		user,
		password,
		host,
		port,
		db,
		timeZone,
	)

	if err := orm.RegisterDataBase("default", DbDriverName, engineUrl); err != nil {
		klog.Error(err.Error())
		os.Exit(1)
	}

	ttl := config.DefaultInt("mysql_conn_ttl", 30)

	orm.ConnMaxLifetime(time.Duration(ttl) * time.Second)
	orm.Debug = config.DefaultBool("mysql_sql_debug", false)

	// table sync
	// create table
	if err := orm.RunSyncdb("default", false, true); err != nil {
		klog.Error(err.Error())
		os.Exit(1)
	}
}
