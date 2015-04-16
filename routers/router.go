// @APIVersion 1
// @Title 记录sms history API v1
package routers

import (
	"dream_api_sms_history/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/history",
			beego.NSInclude(
				&controllers.HistoryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
