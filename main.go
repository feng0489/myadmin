package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	"myadmin/models"
	_ "myadmin/routers"
	"strconv"
	"strings"
	"time"
)

func init()  {

	InitMysql()
	InitSession()
}

func InitMysql()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	dbinfo,_ := fmt.Printf("%v:%v@/%v?charset=utf8",beego.AppConfig.String("mysqluser"), beego.AppConfig.String("mysqlpass"),beego.AppConfig.String("mysqldb"))
	fmt.Println(dbinfo ,"////")
	orm.RegisterDataBase("default","mysql","root:root@/myapp?charset=utf8")
}
var (
	globalSessions *session.Manager
)


func InitSession()  {
	sessionConfig := &session.ManagerConfig{
		CookieName:"myapp",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}

	globalSessions,_ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}
func main() {
    orm.Debug = true
	o := orm.NewOrm()
	beego.AddFuncMap("timeFormat",TimeFormat)
	beego.AddFuncMap("Checked",Checked)
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	beego.Run()
}


func  TimeFormat(int_time int64) string {
	return time.Unix(int_time, 0).Format("2006-01-02 15:04")
}

func Checked(menuId int64,roleId string) bool {
  var checked bool=false
  ok :=models.InArr(strconv.FormatInt(menuId,10),strings.Split(roleId, ","))
  if ok {
	  checked=true
  }
  return  checked
}