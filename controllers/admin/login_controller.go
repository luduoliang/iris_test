package admin

import (
	"github.com/kataras/iris"
	"github.com/mojocn/base64Captcha"
	"iris_test/models"
	"iris_test/common"
)

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

//登录
func Login(c iris.Context){
	if(c.Method() == "POST"){
		idKeyDTemp := common.GetSession(c,"captcha_code")
		if(idKeyDTemp == nil){
			c.JSON(common.JsonData(false, "", "验证码不正确"))
			return
		}

		idKeyD := idKeyDTemp.(string)
		verifyValue := c.FormValue("captcha")
		verifyResult := base64Captcha.VerifyCaptcha(idKeyD, verifyValue)
		if !verifyResult {
			c.JSON(common.JsonData(false, "", "验证码不正确"))
			return
		}

		username := c.FormValue("name")
		password := c.FormValue("pwd")
		if(username == ""){
			c.JSON(common.JsonData(false, "", "请输入用户名"))
			return
		}
		if(password == ""){
			c.JSON(common.JsonData(false, "", "请输入密码"))
			return
		}
		admin_info := models.GetAdminByUserName(username, password)
		if(admin_info.Id <= 0 || admin_info.Status != 1){
			c.JSON(common.JsonData(false, "", "用户名或密码不正确"))
			return
		}

		common.SetSession(c, "admin_info", admin_info)
		c.JSON(common.JsonData(true, admin_info, "登录成功"))
	}else{
		c.View("admin/login/login.html")
	}
}

//退出登录
func Logout(c iris.Context){
	common.DelSession(c, "admin_info")
	c.Redirect("/admin/login", iris.StatusFound)
	return
}

//验证码
func Captcha(c iris.Context){
	//config struct for digits
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     60,
		Width:      120,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 3,
	}
	/*//config struct for audio
	//声音验证码配置
	var configA = base64Captcha.ConfigAudio{
		CaptchaLen: 6,
		Language:   "zh",
	}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height:             60,
		Width:              240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}*/
	/*//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)*/
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	common.SetSession(c,"captcha_code", idKeyD)
	c.JSON(common.JsonData(true, base64stringD, "操作成功"))
	//fmt.Println(idKeyA, base64stringA, "\n")
	//fmt.Println(idKeyC, base64stringC, "\n")
	//fmt.Println(idKeyD, base64stringD, "\n")
}

