package mzjcaptcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type DriverType int

const (
	DIGIT   DriverType = iota //数字验证码
	MATH                      //公司验证码
	STRING                    //数字字母验证码
	CHINESE                   //中文验证码
	AUDIO                     //语言验证码
)

//ConfigJSONBody json request body.
type ConfigJSONBody struct {
	ID            string                       `json:"id" form:"id"`
	CaptchaType   DriverType                   `json:"captchaType" form:"captchaType"`
	VerifyValue   string                       `json:"verifyValue" form:"verifyValue"`
	DriverAudio   *base64Captcha.DriverAudio   `json:"driverAudio" form:"driverAudio"`
	DriverString  *base64Captcha.DriverString  `json:"driverString" form:"driverString"`
	DriverChinese *base64Captcha.DriverChinese `json:"driverChinese" form:"driverChinese"`
	DriverMath    *base64Captcha.DriverMath    `json:"driverMath" form:"driverMath"`
	DriverDigit   *base64Captcha.DriverDigit   `json:"driverDigit" form:"driverDigit"`
}


func Default() (string, string, error) {
	c := ConfigJSONBody{
		CaptchaType: DIGIT,
	}
	return NewCaptcha(&c)
}

//NewCaptcha 创建新的验证码 base64Captcha create http handler
func NewCaptcha(c *ConfigJSONBody) (id string,base64 string,err error) {
	//初始化赋值
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch c.CaptchaType {
	case AUDIO:
		if c.DriverAudio == nil {
			c.DriverAudio = base64Captcha.DefaultDriverAudio
		}
		driver = c.DriverAudio
		break
	case STRING:
		if c.DriverString == nil {
			c.DriverString = base64Captcha.NewDriverString(80, 240, 6, 2, 6, "abcdefghijklmnopqrstuvwxyz1234567890", &color.RGBA{R: 255, G: 255, A: 0, B: 255}, []string{})
		}
		driver = c.DriverString.ConvertFonts()
		break
	case MATH:
		if c.DriverMath == nil {
			c.DriverMath = base64Captcha.NewDriverMath(80, 240, 5, 6, &color.RGBA{R: 255, G: 255, A: 255, B: 0}, []string{})
		}
		driver = c.DriverMath.ConvertFonts()
		break
	case CHINESE:
		if c.DriverChinese == nil {
			c.DriverChinese = base64Captcha.NewDriverChinese(80, 240, 2, 1, 12, "你是什么东西啊我才想你打野在哪里呢我不知大你是射门东西啊测试百度一家人进入电脑完善许霆啥变化内衣媒体没有", &color.RGBA{R: 255, G: 255, A: 0, B: 255}, []string{})
		}
		driver = c.DriverChinese.ConvertFonts()
		break
	case DIGIT:
		if c.DriverDigit == nil {
			c.DriverDigit = base64Captcha.DefaultDriverDigit
		}
		driver = c.DriverDigit
		break
	default:
		if c.DriverDigit == nil {
			c.DriverDigit = base64Captcha.DefaultDriverDigit
		}
		driver = c.DriverDigit
		break
	}
	x := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return x.Generate()
}

//Verify 是否存在
func Verify(id, value string) bool {
	//verify the captcha
	//fmt.Println(id,store.Get(id,true))
	//return store.Get(id,true)==value
	return base64Captcha.DefaultMemStore.Verify(id, value, true)
}
func Set(id,value string){//赋值
	base64Captcha.DefaultMemStore.Set(id,value)
}
func Get(id string,clear bool)string  {//取值
	return base64Captcha.DefaultMemStore.Get(id,clear)
}
func main() {
	i, b, _ := Default()
	fmt.Println(i, b)
	var code string
	fmt.Println("请输入验证码:")
	// 当程序执行到 fmt.Scanl(&name), 程序会停止这里, 等待用户输入, 并回车
	fmt.Scanln(&code)
	//fmt.Scanf("验证码:%s",&code)
	fmt.Println(code)
	fmt.Println(Verify(i, code))
}
