package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dream_api_sms_history/controllers:HistoryController"] = append(beego.GlobalControllerRouter["dream_api_sms_history/controllers:HistoryController"],
		beego.ControllerComments{
			"LogSmsSend",
			`/smssend`,
			[]string{"post"},
			nil})

}
