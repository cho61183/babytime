package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type T_user struct {
	User_id   int64 `orm:"pk;auto"`
	User_name string
	Device_id string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@/yoyojie_user?charset=utf8", 30)
	orm.RegisterModel(new(T_user))
	//orm.RunSyncdb("default", false, true)
}

func Get_user_info_name(uid int64) string {
	fmt.Println(uid)
	o := orm.NewOrm()
	//var maps []orm.Params
	var t_user T_user
	// 获取 QuerySeter 对象，user 为表名
	err := o.QueryTable("t_user").Filter("user_id", 16).One(&t_user, "user_name", "device_id")
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
		return ""
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return ""
		fmt.Printf("Not row found")
	}
	return t_user.User_name

}
