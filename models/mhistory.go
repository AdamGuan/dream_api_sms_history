package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"dream_api_sms_history/helper"
)

func init() {
}

type MHistory struct {
}

func (u *MHistory) AddASmsLog(pkg string,debug int) {
	o := orm.NewOrm()
	o.Raw("INSERT INTO t_sms_send_history SET F_pkg=?,F_debug=?,F_create_datetime=?",pkg,debug,helper.GetNowDateTime()).Exec()
}