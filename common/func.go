package common

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"crypto/md5"
	"io"
)

//返回封装json数据
func JsonData(status bool, objects interface{}, msg string) (apijson *ApiJson) {
	apijson = &ApiJson{Status: status, Data: objects, Msg: msg}
	return apijson
}

//md5密码
func Md5(password string) string {
	w := md5.New()
	io.WriteString(w, password)   //将str写入到w中
	password_md5 := fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式
	return password_md5
}

//从字符串转换成int64方法
func Str2Int64(str string) int64 {
	str_int64, err := strconv.ParseInt(str, 10, 64)
	if err != nil{
		return 0
	}
	return str_int64
}

//从int64转换成字符串方法
func Int642Str(data int64) string {
	return strconv.FormatInt(data, 10)
}

//从float64转换成字符串方法
func Float642Str(f float64) string{
	return strconv.FormatFloat(f, 'f', 2, 64)
}

//从字符串转换成float64方法
func Str2Float64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

//获取为int64的参数
func ParamsInt64(data interface{}) (result int64) {
	result = 0
	switch data.(type) { //多选语句switch
	case string:
		result = Str2Int64(data.(string))
	case float64:
		result =  int64(data.(float64))
	}
	return result
}

//获取当前时间戳
func CurrentTime() int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location).UTC().Unix()
}

//截取字符串，含中文
func Substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}
		if sl + rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

//判断是否在数组中
func InArray(need string, needArr []string) bool {
	for _,v := range needArr{
		if need == v{
			return true
		}
	}
	return false
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}



//随机一段整数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min) + min
	return randNum
}

//字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}


//生成干扰字符串
func RandString(len int) string {
	str := "1123456789QWERTYUIOPASDFGHJKLZXVBNMqwertyuioplkjhgfdsamnbvcxzz"
	var return_str string = ""
	for i := 0;i < len; i++ {
		index := GenerateRangeNum(i+1, 57)
		return_str = return_str + str[index:index+1]
	}
	return return_str
}


//反转数组
func ReverseArray(arr []interface{}) []interface{} {
	len := len(arr)
	temp_arr := make([]interface{}, 0)
	for i:=len-1;i>=0;i--{
		temp_arr = append(temp_arr, arr[i])
	}
	return temp_arr
}



