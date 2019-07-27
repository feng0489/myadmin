package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Admin struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Create_time int64 `json:"create_time"`
	Role string `json:"role"`
	Parent string `json:"parent"`
	Money float64 `json:"money"`

}

type AdminList struct {
	LastPage int `json:"lastPage"`//最后一页
	Page int `json:"page"`//当前页
	PrePage int   `json:"prePage"` //上一页
	NextPage int   `json:"nextPage"` //下一页
	PageSize int `json:"pageSize"`//每页显示行数
	TotalPage int `json:"totalPage"` //总页数
	Total int `json:"total"` //总条数
	Admin []Admin `json:"admin"`//管理员信息
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Admin))
}

func (u *Admin) TableName() string {
	return "app_admin"
}

func AdminByName( username string) (admin *Admin ) {
	o := orm.NewOrm()
	a := Admin{Username: username}
	err := o.Read(&a,"username")
	if err!=nil {
		return nil
	}else{
		return &a
	}
}

func AdminPage(page int,pageSize int,where string) (a *AdminList) {

	var (
		admin []Admin
		info AdminList
	)

	o := orm.NewOrm()
    //计算总数
	var totalItem, totalpages int = 0, 0
	o.Raw("SELECT count(id) FROM  app_admin " + where).QueryRow(&totalItem) //获取总条数

	if totalItem <= pageSize {
		totalpages = 1
	} else if totalItem > pageSize {
		temp := totalItem / pageSize
		if (totalItem % pageSize) != 0 {
			temp = temp + 1
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

	sql := "SELECT id, username, create_time, role, money FROM app_admin "+where+" ORDER BY id DESC LIMIT "+strconv.Itoa(pageSize) +" OFFSET "+strconv.Itoa(pageInfo)
	fmt.Println("sql:" ,sql)
	o.Raw(sql).QueryRows(&admin)
	info.Admin = admin

	return &info
}

func AdminInfo(id int64) (a *Admin) {
	var admin Admin
	o := orm.NewOrm()
	o.Raw("SELECT id,username,password,role,create_time,parent,money FROM  app_admin where id=" + strconv.FormatInt(id,10)).QueryRow(&admin) //获取总条数
	return &admin
}

func AddAdmin(a *Admin) int64 {
	o := orm.NewOrm()
	id,err :=o.Insert(a)
	if err != nil {
		return 0
	}else {
		return id
	}

}

func RemoveAdmin(id int64) int64 {
	o := orm.NewOrm()
	if num, err := o.Delete(&Admin{Id: id}); err == nil {
		return  num
	}else {
		return  0
	}
}

func EditAdmin(a *Admin) int64 {
	o := orm.NewOrm()
	ad := Admin{Id: a.Id}
	o.Read(&ad)
	if  &ad != nil {
		if a.Password == "" {
			a.Password = ad.Password
		}
		if a.Username == "" {
			a.Username = ad.Username
		}
		if num, err := o.Update(a,"username","role","password"); err == nil {
			return num
		}else {
			return 0
		}
	}else {
		return 0
	}
}

