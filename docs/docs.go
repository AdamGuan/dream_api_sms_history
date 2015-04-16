package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1","swaggerVersion":"1.2","apis":[{"path":"/history","description":"用户\n"}],"info":{"title":"记录sms history API v1"}}`
    Subapi string = `{"/history":{"apiVersion":"1","swaggerVersion":"1.2","basePath":"","resourcePath":"/history","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/smssend","description":"","operations":[{"httpMethod":"POST","nickname":"增加一条短信发送记录","type":"","summary":"增加一条短信发送记录","parameters":[{"paramType":"form","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"debug","description":"debug","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]}],"models":{"MResp":{"id":"MResp","properties":{"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}}}}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	err := json.Unmarshal([]byte(Rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(Subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = BasePath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
