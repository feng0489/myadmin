package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"myadmin/models"
	"myadmin/mycashe"
	"strconv"
	"time"
)

type IndexController struct {
	beego.Controller
}

var Admins *models.Admin
var Mes models.MenuList
func (i *IndexController) Prepare()  {
	 adminId :=i.GetSession("adminId")
	if  adminId==nil {
		i.Redirect("login", 302)
	}else{
        adminStr := mycashe.GetCashe("admin_"+strconv.FormatInt(adminId.(int64),10))
		json.Unmarshal(adminStr, &Admins)
		if Admins == nil {
			i.Redirect("/login", 302)
		}

	}

}



func (i *IndexController) Index()  {
	key :=mycashe.StrToBase64(Admins.Role)

	meCashe := mycashe.GetCashe("menus_"+key)
	json.Unmarshal(meCashe, &Mes)
	if Mes.Nav == nil{
		Mes =*models.MenuByRoleName(Admins.Role)
		if &Mes != nil{
			mycashe.SetCashe("menus_"+key,Mes,  10000000 * time.Second)
		}
	}
	i.Data["nav"] = Mes.Nav
	i.Data["menu"] = Mes.NavMune
	i.Data["action"] = Mes.Action
	i.TplName = "index/index.tpl"
}

func (i *IndexController) Welcome()  {
	i.TplName = "index/welcome.tpl"
}


func (i *IndexController) Info()  {
	username:=i.GetString("username","")
	fmt.Println("username:",username)
	i.SetSession("username",username)
	i.Data["json"] =  map[string]interface{}{"success": 1, "message": username}
	i.ServeJSON()
}

