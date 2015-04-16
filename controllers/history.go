package controllers

import (
	"dream_api_sms_history/models"
	"dream_api_sms_history/helper"
)

//用户
type HistoryController struct {
	BaseController
}

// @Title 增加一条短信发送记录
// @Description 增加一条短信发送记录
// @Param	pkg					form	string	true	包名
// @Param	debug				form	int		true	debug
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /smssend [post]
func (u *HistoryController) LogSmsSend() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var historyObj *models.MHistory
	//parse request parames
	u.Ctx.Request.ParseForm()
	pkg := u.Ctx.Request.FormValue("pkg")
	debug := u.Ctx.Request.FormValue("debug")
	//
	if len(pkg) > 0 && len(debug) > 0{
		//记录一条sms发送的log
		debug2 := helper.StrToInt(debug)
		historyObj.AddASmsLog(pkg,debug2)
	}
	//return
	u.jsonEcho(datas)
}