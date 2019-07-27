package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"myadmin/models"
	"strconv"
	"time"
)

type AdminController struct {
	beego.Controller
}



func (a *AdminController) Prepare()  {
	admin :=a.GetSession("adminId")

	if  admin==nil {
		a.Redirect("/login", 302)
	}

}

func (a *AdminController) Index()  {
	a.Data["Website"] = "beego.me"
	a.Data["Email"] = "astaxie@gmail.com"
	a.TplName = "index.tpl"
}

func (a *AdminController) Lst()  {
	username :=a.GetString("username","")
	page,_ :=a.GetInt("page",1)
	pageSize,_ :=a.GetInt("pageSize",20)
	where :=""
	url :="lst?pageSize="+strconv.Itoa(pageSize)
	if username != ""{
		url +="&username="+username
		where +="where username='"+username+"'"
	}

	if Admins.Role != "超级管理员"{
		where += " and parent = '"+Admins.Username+"'"
	}

	info:= models.AdminPage(page,pageSize,where)

	a.Data["admins"] = info.Admin
	a.Data["pageSize"] = info.PageSize

	if info.NextPage ==0 {
		a.Data["nextPage"] = ""
	}else{
		a.Data["nextPage"] = url+"&page="+strconv.Itoa(info.NextPage)
	}

	a.Data["firstPage"] = url+"&page=1"
	a.Data["page"] = url+"&page="+strconv.Itoa(info.Page)
	a.Data["pageInfo"] = info.Page

	a.Data["lastPage"] =  url+"&page="+strconv.Itoa(info.LastPage)
	a.Data["total"] = info.Total
	a.Data["totalPage"] = info.TotalPage

	if info.PrePage ==0 {
		a.Data["prePage"] = ""
	}else{
		a.Data["prePage"] = url+"&page="+strconv.Itoa(info.PrePage)
	}
    menu :=Mes.Action
	a.Data["addName"] =menu[101][10101]["name"]
	a.Data["addAction"] =menu[101][10101]["action"]
	a.Data["editName"] =menu[101][10102]["name"]
	a.Data["editAction"] =menu[101][10102]["action"]
	a.Data["removeName"] =menu[101][10103]["name"]
	a.Data["removeAction"] =menu[101][10103]["action"]

	//jsonBytes, _ := json.Marshal(a.Data["admins"])
	//fmt.Println("adminList:" ,string(jsonBytes))
	//a.Data["json"] = info
	//a.ServeJSON()
	a.TplName = "admin/list.tpl"
}

func (a *AdminController) Add()  {
     var r []models.Role

     r = *models.RoleNameList()

     a.Data["role"] = r
	 a.Data["action"] ="initAdd"
	 a.TplName = "admin/add.tpl"
}

func (a *AdminController) InitAdd()  {
	username :=a.GetString("username","")
	role :=a.GetString("role","")
	pass :=a.GetString("pass","")
	repass :=a.GetString("repass","")

	if username == ""{
		a.Data["json"] =  map[string]interface{}{"code": 1, "msg": "请输入用户名"}
		a.ServeJSON()
		return
	}
	if role == ""{
		a.Data["json"] =  map[string]interface{}{"code": 2, "msg": "请选择管理员角色"}
		a.ServeJSON()
		return
	}

	if pass == ""{
		a.Data["json"] =  map[string]interface{}{"code": 3, "msg": "请输入密码"}
		a.ServeJSON()
		return
	}
	if pass != repass{
		a.Data["json"] =  map[string]interface{}{"code": 4, "msg": "两次输入的密码不一致"}
		a.ServeJSON()
		return
	}
	var ad *models.Admin
	ad= models.AdminByName(username)
	if ad != nil{
		a.Data["json"] =  map[string]interface{}{"code": 5, "msg": "该管理员已经存在"}
		a.ServeJSON()
		return
	}
	h := md5.New()
	h.Write([]byte(pass)) // 需要加密的字符串为 password
	password :=hex.EncodeToString(h.Sum(nil))
	newa := new(models.Admin)
	newa.Username = username
	newa.Password = password
	newa.Role = role
	newa.Parent = Admins.Username
	newa.Create_time = time.Now().Unix()

    id :=models.AddAdmin(newa)
    if id >0{
		a.Data["json"] =  map[string]interface{}{"code": 200, "msg": "ok","data":id}
		a.ServeJSON()
	}else{
		a.Data["json"] =  map[string]interface{}{"code": 6, "msg": "系统异常请稍候重试"}
		a.ServeJSON()
	}

}

func (a *AdminController) Remove()  {
	id,_ := a.GetInt64("id",0)
	if id == 0 {
		a.Data["json"] =  map[string]interface{}{"code": 1, "msg": "信息异常,请刷新页面后再重试"}
		a.ServeJSON()
		return
	}
	num := models.RemoveAdmin(id)
	if num==0{
		a.Data["json"] =  map[string]interface{}{"code": 1, "msg": "系统异常,请稍后重试"}
		a.ServeJSON()
		return
	}
	a.Data["json"] =  map[string]interface{}{"code": 200, "msg": "ok"}
	a.ServeJSON()

}

func (a *AdminController) Edit()  {
	id,_ := a.GetInt64("id",0)
	var initA  models.Admin
	initA = *models.AdminInfo(id)
    if &initA == nil {
		a.Data["json"] =  map[string]interface{}{"code": 1, "msg": "该管理员信息已不存在"}
		a.ServeJSON()
		return
	}
	var r []models.Role

	r = *models.RoleNameList()
	a.Data["username"] = initA.Username
	a.Data["rolename"] = initA.Role
	a.Data["id"] = initA.Id
	a.Data["role"] = r
	a.Data["action"] ="initEdit"
	a.TplName = "admin/edit.tpl"
}

func (a AdminController) InitEdit()  {
	id,_ := a.GetInt64("id",0)
	role := a.GetString("role","")
	username := a.GetString("username","")
	pass := a.GetString("pass","")
	if id == 0 {
		a.Data["json"] =  map[string]interface{}{"code": 1, "msg": "该管理员信息已不存在"}
		a.ServeJSON()
		return
	}
	if role == "" {
		a.Data["json"] =  map[string]interface{}{"code": 2, "msg": "请选择管理员角色"}
		a.ServeJSON()
		return
	}
	if username == "" {
		a.Data["json"] =  map[string]interface{}{"code": 3, "msg": "请输入管理员名称"}
		a.ServeJSON()
		return
	}
	var ad *models.Admin
	ad = models.AdminByName(username)
	if ad != nil {
		if ad.Id != id {
			a.Data["json"] =  map[string]interface{}{"code": 4, "msg": "该管理员已经存在"}
			a.ServeJSON()
			return
		}
	}
	h := md5.New()
	h.Write([]byte(pass)) // 需要加密的字符串为 password
	password :=hex.EncodeToString(h.Sum(nil))
	num := models.EditAdmin(&models.Admin{Id:id,Username:username,Role:role,Password:password})
	if num>0{
		a.Data["json"] =  map[string]interface{}{"code": 200, "msg": "ok"}
		a.ServeJSON()
	}else{
		a.Data["json"] =  map[string]interface{}{"code": 5, "msg": "系统异常,请稍后重试"}
		a.ServeJSON()
	}
}

func (a *AdminController) Role()  {
	//TO Do
	a.TplName = "role/role.tpl"
}
func (a *AdminController) Log()  {
	a.TplName = "admin/log.tpl"
}
func (a *AdminController) Info()  {
	username:=a.GetString("username","")
	fmt.Println("username:",username)
	a.SetSession("username",username)
	a.Data["json"] =  map[string]interface{}{"success": 1, "message": username}
	a.ServeJSON()
}