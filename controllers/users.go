package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UsersController struct {
	beego.Controller
}



func (u *UsersController) Prepare()  {
	admin :=u.GetSession("admin")
	if  admin==nil {
		u.Redirect("/login", 302)
	}

}

func (u *UsersController) Index()  {
	u.Data["Website"] = "beego.me"
	u.Data["Email"] = "astaxie@gmail.com"
	u.TplName = "index.tpl"
}

func (u *UsersController) Lst()  {
	u.TplName = "users/list.tpl"
}
func (u *UsersController) Role()  {
	u.TplName = "admin/role.tpl"
}
func (u *UsersController) Log()  {
	u.TplName = "admin/log.tpl"
}
func (u *UsersController) Info()  {
	username:=u.GetString("username","")
	fmt.Println("username:",username)
	u.SetSession("username",username)
	u.Data["json"] =  map[string]interface{}{"success": 1, "message": username }
	u.ServeJSON()
}