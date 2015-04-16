package main

import (
	_ "dream_api_sms_history/docs"
	_ "dream_api_sms_history/routers"
	"dream_api_sms_history/controllers"

	"github.com/astaxie/beego"
	"runtime"
	"github.com/astaxie/beego/config" 
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
