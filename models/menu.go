package models

import (
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id int64
	Menu_id int64
	Name string
	Action string
	Parent_id int64
}
type MenuList struct {
	Nav []map[string]interface{} `json:"nav"`
	NavMune []map[string]interface{} `json:"navmune"`
	NavAction []map[string]interface{} `json:"navaction"`
	Menus map[int64]map[int64]map[string]interface{} `json:"menus"`
	Action map[int64]map[int64]map[string]interface{} `json:"action"`
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Menu))
}

func (m *Menu) TableName() string {
	return "app_menu"
}

func AllMenu() (mene *[]*Menu,err error) {
	var menus []*Menu
	o := orm.NewOrm()
	total,errs := o.QueryTable("app_menu").All(&menus)
	if errs != nil{
		return nil,errs
	}
	if total ==0{
		return nil,nil
	}
	return &menus,nil
}

func Menus( ids string) (m *MenuList){
	var menus []Menu
	o := orm.NewOrm()
	var menu MenuList
	var sql string
	if ids == "" {
		sql = "select id,menu_id,name,action,parent_id from app_menu"
	}else{
		sql = "select id,menu_id,name,action,parent_id from app_menu where id in("+ids+")"
	}
	o.Raw(sql).QueryRows(&menus)
	if menus == nil{
		return nil
	}
	var nav []map[string]interface{}
    var navmune []map[string]interface{}
    var navaction []map[string]interface{}
	menuss :=make(map[int64]map[int64]map[string]interface{})
	action :=make(map[int64]map[int64]map[string]interface{})

	for _,v :=range menus {
		if v.Id<100 {
			nav = append(nav, map[string]interface{}{"id": v.Id,"name":v.Name,"action":v.Action})
		}
		if v.Id >100 && v.Id< 1000{
			if menuss[v.Menu_id] == nil{
				menuss[v.Menu_id] = make(map[int64]map[string]interface{})
			}
			if menuss[v.Menu_id][v.Id] == nil {
				menuss[v.Menu_id][v.Id] =  make(map[string]interface{})
			}
			navmune = append(navmune, map[string]interface{}{"id": v.Id,"name":v.Name,"action":v.Action,"mid":v.Menu_id})
			menuss[v.Menu_id][v.Id]["id"]=v.Id
			menuss[v.Menu_id][v.Id]["name"]=v.Name
			menuss[v.Menu_id][v.Id]["action"]=v.Action
			menuss[v.Menu_id][v.Id]["mid"]=v.Menu_id

		}
		if v.Id> 10000 {
			if action[v.Menu_id] == nil{
				action[v.Menu_id] = make(map[int64]map[string]interface{})
			}
			if action[v.Menu_id][v.Id] == nil {
				action[v.Menu_id][v.Id] =  make(map[string]interface{})
			}
			navaction = append(navaction,map[string]interface{}{"id": v.Id,"name":v.Name,"action":v.Action,"mid":v.Parent_id})
			action[v.Menu_id][v.Id]["id"]=v.Id
			action[v.Menu_id][v.Id]["name"]=v.Name
			action[v.Menu_id][v.Id]["action"]=v.Action
			action[v.Menu_id][v.Id]["mid"]=v.Menu_id

		}

	}

	menu.Nav = nav
	menu.Menus = menuss
	menu.Action = action
	menu.NavMune = navmune
	menu.NavAction = navaction
	return &menu
}