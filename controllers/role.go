package controllers

import (
	"github.com/astaxie/beego"
	"myadmin/models"
	"strconv"
	"time"
)

type RoleController struct {
	beego.Controller
}


func (r *RoleController) Prepare()  {
	 adminId :=r.GetSession("adminId")
	if  adminId==nil {
		r.Redirect("/login", 302)
	}
}



func (r *RoleController) List()  {
    name := r.GetString("name","")
    page,_ := r.GetInt("page",1)
    pageSize,_ :=r.GetInt("pageSize",20)
    where :=""
    url :="role?pageSize="+strconv.Itoa(pageSize)
    if name !=""{
		url +="&name="+name
		where += " where name='"+name+"'"
	}
    if Admins.Role != "超级管理员"{
		where += " and create_admin ='"+Admins.Username+"'"
	}

	info:= models.RolePage(page,pageSize,where)
	r.Data["firstPage"] = url+"&page=1"
	r.Data["page"] = url+"&page="+strconv.Itoa(info.Page)
	r.Data["pageInfo"] = info.Page
	r.Data["list"] = info.Role
	r.Data["pageSize"] = info.PageSize

	r.Data["lastPage"] =  url+"&page="+strconv.Itoa(info.LastPage)
	r.Data["total"] = info.Total
	r.Data["totalPage"] = info.TotalPage

	if info.PrePage ==0 {
		r.Data["prePage"] = ""
	}else{
		r.Data["prePage"] = url+"&page="+strconv.Itoa(info.PrePage)
	}
	if info.NextPage ==0 {
		r.Data["nextPage"] = ""
	}else{
		r.Data["nextPage"] = url+"&page="+strconv.Itoa(info.NextPage)
	}

	menu :=Mes.Action
	r.Data["addName"] =menu[102][10201]["name"]
	r.Data["addAction"] =menu[102][10201]["action"]
	r.Data["editName"] =menu[102][10202]["name"]
	r.Data["editAction"] =menu[102][10202]["action"]
	r.Data["removeName"] =menu[102][10203]["name"]
	r.Data["removeAction"] =menu[102][10203]["action"]
	r.TplName = "role/role.tpl"
}

func (r *RoleController) Add(){
	r.Data["action"] = "initAdd"
	r.TplName = "role/add.tpl"
}

func (r *RoleController) InitAdd()  {
	name :=r.GetString("name","")
    role :=models.RoleByName(name)
    if role != nil{
		r.Data["json"]=map[string]interface{}{"code":1,"msg":"该角色已经存在"}
		r.ServeJSON()
		return
	}
	rn := new(models.Role)
	rn.Name=name
	rn.Minu_id=""
	rn.Create_admin=Admins.Username
	rn.Create_time = time.Now().Unix()
	id :=models.AddRole(rn)
	if id >0 {
		r.Data["json"]=map[string]interface{}{"code":200,"msg":"ok"}
		r.ServeJSON()
	}else{
		r.Data["json"]=map[string]interface{}{"code":2,"msg":"网络异常,请稍候重试"}
		r.ServeJSON()
	}

}

func (r *RoleController) Edit()  {
	id,_ :=r.GetInt64("id",0)
	role:=models.RoleById(id)
	menus :=models.Menus( "")
	r.Data["role"]=role
	r.Data["nav"]=menus.Nav
	r.Data["menu"]=menus.NavMune
	r.Data["action"]=menus.NavAction
	r.Data["act"] = "initEdit"
	//r.Data["json"]=map[string]interface{}{"menu":menus.NavMune,"Nav":menus.Nav,"action":menus.NavAction}
	//r.ServeJSON()
	r.TplName="role/edit.tpl"
}

func (r *RoleController) InitEdit()  {

	 ids :=r.GetString("ids","")
	 name :=r.GetString("name","")
	 id,_ := r.GetInt64("role_id",0)

	 if id == 0 && name == ""{
		 r.Data["json"]=map[string]interface{}{"code":1,"msg":"参数错误"}
		 r.ServeJSON()
		 return
	 }
	 role :=new(models.Role)
	 role.Id = id
	 role.Minu_id = ids
	 num := models.EditRole(role)
	 if num>0{
		 r.Data["json"]=map[string]interface{}{"code":200,"msg":"ok"}
		 r.ServeJSON()
	 }else{
		 r.Data["json"]=map[string]interface{}{"code":2,"msg":"操作失败,请稍候再重试"}
		 r.ServeJSON()
	 }

}

func (r *RoleController) Remove()  {
    id,_ :=r.GetInt64("id",0)
    if id == 0{
		r.Data["json"]=map[string]interface{}{"code":1,"msg":"参数错误"}
		r.ServeJSON()
		return
	}
    num :=models.RemoveRole(id)
    if num >0 {
		r.Data["json"]=map[string]interface{}{"code":200,"msg":"ok"}
		r.ServeJSON()
	}else{
		r.Data["json"]=map[string]interface{}{"code":2,"msg":"操作失败,请稍候再重试"}
		r.ServeJSON()
	}

}



