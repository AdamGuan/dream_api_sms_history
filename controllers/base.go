package controllers

import (
	"dream_api_sms_history/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms_history/helper"
	"github.com/astaxie/beego/config"
)

//公共controller
type BaseController struct {
	beego.Controller
}

//json echo
func (u *BaseController) jsonEcho(datas map[string]interface{}) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
//		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
//		u.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
		u.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	} 
	
	datas["responseMsg"] = ""
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}

	u.Data["json"] = datas
	//log
	u.logEcho(datas)

	u.ServeJson()
}

//记录请求
func (u *BaseController) logRequest() {
	var logObj *models.MLog
	logObj.LogRequest(u.Ctx)
}

//记录返回
func (u *BaseController) logEcho(datas map[string]interface{}) {
	var logObj *models.MLog
	logObj.LogEcho(datas)
}