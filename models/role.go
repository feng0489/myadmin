package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type Role struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Minu_id string `json:"minu_id"`
	Create_time int64 `json:"create_time"`
	Create_admin string `json:"create_admin"`
}


//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Role))
}


func (u *Role) TableName() string {
	return "app_role"
}

type RoleList struct {
	LastPage int `json:"lastPage"`//最后一页
	Page int `json:"page"`//当前页
	PrePage int   `json:"prePage"` //上一页
	NextPage int   `json:"nextPage"` //下一页
	PageSize int `json:"pageSize"`//每页显示行数
	TotalPage int `json:"totalPage"` //总页数
	Total int `json:"total"` //总条数
	Role []Role `json:"role"`//管理员信息
}

func MenuByRoleName(name string) (m *MenuList) {
	var role Role
	var sql string
	if name == "超级管理员" {
		return Menus("")
	}else{

		sql="select id,minu_id from app_role where name='"+name+","
		o := orm.NewOrm()
		o.Raw(sql).QueryRow(&role)
		if role.Minu_id == ""{
			return nil
		}else{
			return	Menus(role.Minu_id)
		}

	}


}

func RoleNameList() (r *[]Role) {
	var role []Role
	sql :="select id,name from app_role"
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&role)
	return &role
}

func RolePage(page int,pageSize int,where string) (r *RoleList){
	var (
		role []Role
		info RoleList
		)
	o := orm.NewOrm()
	//计算总数
	var totalItem, totalpages int = 0, 0
	countSql := "select count(id) from app_role "+where
	o.Raw(countSql).QueryRow(&totalItem)

	if totalItem <=pageSize{
		totalpages = 1
	}else{
		temp := totalItem/pageSize
		if totalItem %pageSize !=0{
			temp +=1
		}
		totalpages = temp
	}

	info.TotalPage = totalpages
	info.Total = totalItem


	info.LastPage = totalpages
	info.Page = page
	info.PageSize = pageSize

	if page>1{
		info.PrePage = page-1
		if page < totalpages {
			info.NextPage = page+1
		}else{
			info.NextPage = 0
		}
	}else{
		info.PrePage = 0
		if totalpages>1 {
			info.NextPage = page+1
		}else{
			info.NextPage = 0
		}
	}
	pageInfo :=(page-1)*pageSize
    listSql :="select id,name,minu_id,create_time,create_admin from app_role "+where+" order by id desc limit "+strconv.Itoa(pageSize)+" offset "+strconv.Itoa(pageInfo)
    o.Raw(listSql).QueryRows(&role)
    //for k,v:=range role {
    //	if v.Minu_id != ""{
    //		var ne []Menu
    //		menuSql := "select id, name from app_menu where id in ("+v.Minu_id+")"
    //		o.Raw(menuSql).QueryRows(&ne)
    //        var name string=""
    //		for _,val := range ne{
    //			if name == ""{
    //				name += val.Name
	//			}else{
	//				name += ","+val.Name
	//			}
	//		}
	//		role[k].Minu_id = name
	//		fmt.Println("name",role[k].Minu_id)
	//	}
	//}
    //
	var ne []Menu
	menuSql := "select id, name from app_menu"
	o.Raw(menuSql).QueryRows(&ne)

	for k,v := range role {
		if v.Minu_id != ""{
			var name string=""
			s := strings.Split(v.Minu_id, ",")
			for _,val := range ne {
				if InArr(strconv.FormatInt(val.Id,10),s){
					if name == ""{
						name += val.Name
					}else{
						name += ","+val.Name
					}
				}
			}
			role[k].Minu_id = name
		}
	}

    info.Role = role

	return &info
}

func RoleByName(name string) (r *Role) {
	var role Role
	sql :="select id,name from app_role where name = '"+name+"'"

	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&role)
	if role.Id>0{
		return &role
	}else{
		return nil
	}

}

func AddRole(r *Role) int64  {
	o:=orm.NewOrm()
	id,err :=o.Insert(r)
	if err != nil {
		return 0
	}else {
		return id
	}
}

func RoleById(id int64) (r *Role) {
	o := orm.NewOrm()
	role := Role{Id: id}

	err := o.Read(&role)

	if err == orm.ErrNoRows {
		return nil
	} else if err == orm.ErrMissPK {
		return nil
	} else {
		return &role
	}
}

func EditRole(r *Role) int64 {
	o := orm.NewOrm()
	ad := Role{Id: r.Id}
	o.Read(&ad)
	if  &ad != nil {

		if num, err := o.Update(r,"minu_id"); err == nil {
			return num
		}else {
			return 0
		}
	}else {
		return 0
	}
}

func RemoveRole(id int64) int64 {
	o := orm.NewOrm()
	if num,err:=o.Delete(&Role{Id:id});err ==nil{
		return num
	}else{
		return 0
	}
}