package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Create_time int64 `json:"create_time"`
	Role string `json:"role"`
	Parent string `json:"parent"`
	Money float64 `json:"money"`

}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Users))
}

func (u *Users) TableName() string {
	return "app_users"
}

