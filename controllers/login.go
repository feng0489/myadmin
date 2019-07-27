package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"myadmin/models"
	"myadmin/mycashe"
	"strconv"
	"strings"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Login()  {
	l.TplName = "index/login.tpl"
}


func (l *LoginController) InLogin()  {
	username:=l.GetString("username","")
	password:=l.GetString("password","")
	fmt.Println("username:",username,"\npassword:",password)
	admin:=models.AdminByName(username)

	if admin == nil {
		l.Data["json"] =  map[string]interface{}{"code": 40102, "msg": "登录失败,用户名或密码错误"}
		l.ServeJSON()
		return
	}
	h := md5.New()
	h.Write([]byte(password)) // 需要加密的字符串为 password
	submit :=hex.EncodeToString(h.Sum(nil))
	if submit != admin.Password {
		l.Data["json"] =  map[string]interface{}{"code": 40103, "msg": "登录失败,用户名或密码错误"}
		l.ServeJSON()
		return
	}

	l.SetSession("adminId",admin.Id)
	mycashe.SetCashe("admin_"+strconv.FormatInt(admin.Id,10),admin,10000*time.Second)
	l.Data["json"] =  map[string]interface{}{"code": 200, "msg": "登录成功"}
	l.ServeJSON()

}

type Student struct {
	Name string `json:"name"`          //对应的json的可以也是大写，可以指定json的二次编码 ,-不会输出到json
	Age  int    `json:"age,string"` //输出为字符串
	Sex  byte   `json:"sex"`         //会转出数字！怎么转出字符？
	Is   bool
	Like []string
	Addr string
}

func (l *LoginController) Test1()  {
	stu := Student{"Laymond", 18, 'm', true, []string{"跑步", "爬山", "学习"}, "中国.重庆"}
	json_byte, err := json.MarshalIndent(stu, "", " ") //看上去更加格式化
	if err != nil {
		fmt.Println("\n struct 转json 失败:",err)
	}
	fmt.Println("success:",string(json_byte))
	l.Data["Website"]= string(json_byte)
    str := string(json_byte)
	var stud Student
	json.Unmarshal([]byte(str), &stud)

    fmt.Println("name:",stud.Name,"\n age:",stud.Age,"\n sex:",stud.Sex)
	l.TplName = "index.tpl"


}
func (l *LoginController) Test2()  {
      menu := models.Menus("超级管理员")
      if menu ==nil {
      	fmt.Println("getMenu nil:")
      }

	  l.Data["json"] =  map[string]interface{}{"data": menu}

		s := strings.Split("1,101", ",")
		fmt.Println("inarr",models.InArr("40101",s))
	  l.ServeJSON()
}