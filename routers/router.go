package routers

import (
	"myadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//首页
    beego.Router("/", &controllers.IndexController{},"*:Index")
	beego.Router("/welcome", &controllers.IndexController{},"*:Welcome")

    //登录
    beego.Router("/login", &controllers.LoginController{},"*:Login")
	beego.Router("/InLogin", &controllers.LoginController{},"*:InLogin")
    beego.Router("/test1", &controllers.LoginController{},"*:Test1")
    beego.Router("/test2", &controllers.LoginController{},"*:Test2")


    //管理员
    beego.Router("/admin/lst", &controllers.AdminController{},"*:Lst")
    beego.Router("/admin/add", &controllers.AdminController{},"*:Add")
    beego.Router("/admin/initAdd", &controllers.AdminController{},"*:InitAdd")
    beego.Router("/admin/remove", &controllers.AdminController{},"*:Remove")
    beego.Router("/admin/edit", &controllers.AdminController{},"*:Edit")
    beego.Router("/admin/initEdit", &controllers.AdminController{},"*:InitEdit")
    beego.Router("/admin/log", &controllers.AdminController{},"*:Log")

    //角色
	beego.Router("/role/lst", &controllers.RoleController{},"*:List")
	beego.Router("/role/add", &controllers.RoleController{},"*:Add")
	beego.Router("/role/initAdd", &controllers.RoleController{},"*:InitAdd")
	beego.Router("/role/edit", &controllers.RoleController{},"*:Edit")
	beego.Router("/role/initEdit", &controllers.RoleController{},"*:InitEdit")
	beego.Router("/role/remove", &controllers.RoleController{},"*:Remove")


    //用户
	beego.Router("/users/lst", &controllers.AdminController{},"*:Lst")

}
