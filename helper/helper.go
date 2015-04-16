package helper

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"crypto/sha1"
)

func init() {
}

//类型转化 string  to int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//类型转化 string  to float64
func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

//类型转化 int to string
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

//类型转化 int64 to string
func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

//获取16位Guid
func GetGuid() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x", b)
	return uuid[0:16]
}

//创建密码
func CreatePwd(num int) string {
	return GetGuid()[0:num]
}

//检查签名
func CheckSign(sign string, token string) bool {
	//sign = timestamp+md5(token+timestamp)
	if len(sign) == 46 && len(token) == 32 {
		timestamp := sign[0:14]
		//检测是否超时
//		appConf, _ := config.NewConfig("ini", "conf/app.conf")
//		debug,_ := appConf.Bool(beego.RunMode+"::debug")
//		if !debug{
		if 1 != 1{
			nowTime, _ := strconv.Atoi(time.Now().Format("20060102150405"))
			requestTime, _ := strconv.Atoi(timestamp)
			timedistince := nowTime - requestTime
			if timedistince > 60*5 {
				return false
			}
		}

		md5Str := sign[14:]
		str := token + timestamp
		md5Str2 := fmt.Sprintf("%x\r\n", md5.Sum([]byte(str)))
		md5Str2 = strings.TrimSpace(md5Str2)
		if string(md5Str) == string(md5Str2) {
			return true
		}
	}
	return false
}

//md5
func Md5(str string)string{
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}

//get now datatime
func GetNowDateTime()string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDateTimeBeforeMinute(num int)string{
	return time.Now().Add(-time.Minute * time.Duration(num)).Format("2006-01-02 15:04:05")
}

func GetDateTimeAfterMinute(num int)string{
	return time.Now().Add(time.Minute * time.Duration(num)).Format("2006-01-02 15:04:05")
}

func Split(str string,flag string)[]string{
	return strings.Split(str, ",")
}

func JoinString(list []string,flag string)string{
	result := ""
	if len(list) > 0{
		for _,v := range list{
			result += v+","
		}
		result = strings.Trim(result,",")
	}
	return result
}

func StringInArray(value string,list []string)bool{
	result := false
	for _,item := range list{
		if value == item{
			result = true
			break
		}
	}
	return result
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Sha1(str string)string{
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}