package mzjcaptcha


import (
	"fmt"
	"image/color"
	"github.com/mojocn/base64Captcha"
)

//ConfigJSONBody json request body.
type ConfigJSONBody struct {
	ID            string                       `json:"id" form:"id"`
	CaptchaType   string                       `json:"captchaType" form:"captchaType"`
	VerifyValue   string                       `json:"verifyValue" form:"verifyValue"`
	DriverAudio   *base64Captcha.DriverAudio   `json:"driverAudio" form:"driverAudio"`
	DriverString  *base64Captcha.DriverString  `json:"driverString" form:"driverString"`
	DriverChinese *base64Captcha.DriverChinese `json:"driverChinese" form:"driverChinese"`
	DriverMath    *base64Captcha.DriverMath    `json:"driverMath" form:"driverMath"`
	DriverDigit   *base64Captcha.DriverDigit   `json:"driverDigit" form:"driverDigit"`
}

var store = base64Captcha.DefaultMemStore

//NewCaptcha 初始化
func(c *ConfigJSONBody) defaultDriver() {
	//初始化赋值
	if c.ID == "" {
		c.ID = "weixiao"
	}
	if c.VerifyValue == "" {
		c.VerifyValue = "weixiao121"
	}
	if c.DriverAudio == nil {
		c.DriverAudio = base64Captcha.DefaultDriverAudio
	}
	if c.DriverDigit == nil {
		c.DriverDigit = base64Captcha.DefaultDriverDigit
	}
	if c.DriverChinese == nil {
		c.DriverChinese = base64Captcha.NewDriverChinese(80, 240, 2, 1, 12, "你是什么东西啊我才想你打野在哪里呢我不知大你是射门东西啊测试百度一家人进入电脑完善许霆啥变化内衣媒体没有", &color.RGBA{R: 255, G: 255, A: 0, B: 255}, []string{})
	}
	if c.DriverMath == nil {
		c.DriverMath = base64Captcha.NewDriverMath(80, 240, 5, 6, &color.RGBA{R: 255, G: 255, A: 255, B: 0}, []string{})
	}
	if c.DriverString == nil {
		c.DriverString = base64Captcha.NewDriverString(80, 240, 6, 2, 6, "abcdefghijklmnopqrstuvwxyz1234567890", &color.RGBA{R: 255, G: 255, A: 0, B: 255}, []string{})
	}
}

//NewCaptcha 创建新的验证码 base64Captcha create http handler
func (c *ConfigJSONBody) NewCaptcha()(string, string, error) {
	c.defaultDriver()
	fmt.Println("开始生成新的验证码")
	fmt.Println(c)
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch c.CaptchaType {
	case "audio":
		driver = c.DriverAudio
		break
	case "string":
		driver = c.DriverString.ConvertFonts()
		break
	case "math":
		driver = c.DriverMath.ConvertFonts()
		break
	case "chinese":
		driver = c.DriverChinese.ConvertFonts()
		break
	case "digit":
		driver = c.DriverDigit
		break
	default:
		driver = c.DriverDigit
		break
	}
	x := base64Captcha.NewCaptcha(driver, store)
	return x.Generate()
}

//QshCaptchaVerifyHandle 是否存在
func QshCaptchaVerifyHandle(id, value string) bool {
	//verify the captcha
	return store.Verify(id, value, true)
}