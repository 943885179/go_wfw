package mzjgin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/micro/go-micro/v2/util/log"
	"io"
	"net/http"
	"os"
	"qshapi/models"
	"qshapi/utils/mzjinit"
	"strings"

	"github.com/gin-gonic/gin"
)
var (
	conf models.APIConfig
)
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
}
//RespCode 返回代码
type RespCode int
const (
	//APIOK 自定义返回状态（成功，登录成功等前端显示#67C23A）
	APIOK RespCode = 10000
	//APIError 自定义返回状态(错误，前端提示#F56C6C)
	APIError RespCode = 10005
	//APIInfo 自定义返回状态（提示,前端显示#909399）
	APIInfo RespCode = 10006
	//APIWary 自定义返回状态（警告，类似密码快过期等前端显示#E6A23C）
	APIWary RespCode = 1008
)

func (c RespCode) String() string {
	switch c {
	case APIOK:
		return "成功"
	case APIWary:
		return "警告"
	case APIInfo:
		return "提示"
	case APIError:
		return "错误"
	default:
		return http.StatusText(int(c))
	}
}

//Resp api接口返回实体
type Resp struct {
	Code   RespCode    `json:"code"`   //编码
	Msg    string      `json:"msg"`    //消息
	Method string      `json:"method"` //请求方式
	URL    string      `json:"url"`    //请求地址
	IP     string      `json:"ip"`     //请求ip
	Result interface{} `json:"result"` //返回值
}

type gomicroErrResp struct {
	ID string `json:"id"`
	Code int `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

//APIResult 返回
func (r Resp) APIResult(c *gin.Context, code int, result interface{}) {
	var respCode = RespCode(code)
	r.Code = respCode
	r.Msg = respCode.String()
	r.Result = result
	r.Method = c.Request.Method
	r.URL = c.Request.URL.String()
	r.IP = c.ClientIP()
	/*str, _ := json.MarshalIndent(resp, "", "   asd ")
	c.String(http.StatusOK, string(str))*/
	switch respCode {
	case APIOK:
		c.JSON(http.StatusOK, r)
		break
	case APIWary:
		c.JSON(http.StatusOK, r)
		break
	case APIInfo:
		c.JSON(http.StatusOK, r)
		break
	case APIError:
		c.JSON(http.StatusOK, r)
		break
	default:
		c.JSON(code, r)
		break
	}
	c.Abort() //退出
}

//APIOK 成功返回
func (r Resp) APIOK(c *gin.Context, result interface{}) {
	r.APIResult(c, int(APIOK), result)
}

//APIError 失败返回
func (r Resp) APIError(c *gin.Context, errMsg string) {
	r.APIResult(c, int(APIError), errMsg)
}

func (r Resp)MicroResp(c *gin.Context,result interface{},err error)  {
	if err != nil {
		er:=gomicroErrResp{}
		json.Unmarshal([]byte(err.Error()),&er)
		r.APIError(c,er.Detail)
	}else {
		r.APIOK(c,result)
	}
}

//APIInfo 提示返回
func (r Resp) APIInfo(c *gin.Context, errMsg string) {
	r.APIResult(c, int(APIInfo), errMsg)
}

//APIWary 警告返回
func (r Resp) APIWary(c *gin.Context, errMsg string) {
	r.APIResult(c, int(APIWary), errMsg)
}
var (
  notoken=[]string{"/user/login", "/user/addUser","/static","/swagger","/favicon.ico","/login","/registry","/codeVerify","/sendCode" }
  apiresp= Resp{}
)
//APIGin 自定义gin
type APIGin struct {
}

func NewGin() *APIGin {
	result:=APIGin{}
	return &result
}
func(api *APIGin) Default() *gin.Engine{
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	fe,_:=os.Create("gin_error.log")
	gin.DefaultErrorWriter=io.MultiWriter(fe)
	g:=gin.Default()
	g.Use(api.cors())//支持跨域
	g.NoMethod(handleNotFound)
	g.NoRoute(handleNotFound)
	g.Use(gzip.Gzip(gzip.DefaultCompression)) //使用gzip压缩
	//添加Token中间件
	//g.Use(APITokenMiddleware)
	//或者使用下面的方法
	//g.engine.Use(TokenAuthMiddleware())
	// 加载html文件，即template包下所有文件
	//g.engine.LoadHTMLGlob("wwwroot/*")
	//g.engine.LoadHTMLGlob("template/*")
	//g.Static("/assets", "/var/www/tizi365/assets")// /assets/images/1.jpg 这个url文件，存储在/var/www/tizi365/assets/images/1.jpg
	// 静态资源加载
	 g.StaticFS("static", http.Dir("./static"))
	//g.StaticFS("/", http.Dir("./static/upload"))
	//g.StaticFS("/public", http.Dir("D:/goproject/src/github.com/ffhelicopter/tmm/website/static"))
	//g.StaticFile("/favicon.ico", "./resources/favicon.ico")
	return g
}
func (api *APIGin)Run(addr string) {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	fe,_:=os.Create("gin_error.log")
	gin.DefaultErrorWriter=io.MultiWriter(fe)
	g:=gin.Default()
	g.Use(api.cors())//支持跨域
	g.NoMethod(handleNotFound)
	g.NoRoute(handleNotFound)
	//添加Token中间件
	g.Use(APITokenMiddleware)
	//或者使用下面的方法
	//g.engine.Use(TokenAuthMiddleware())
	// 加载html文件，即template包下所有文件
	//g.engine.LoadHTMLGlob("wwwroot/*")
	//g.engine.LoadHTMLGlob("template/*")
	g.StaticFS("/static", http.Dir("./static"))
	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Run(addr)
}
//跨域
func (g *APIGin)cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func handleNotFound(c *gin.Context) {
	apiresp.APIResult(c, http.StatusNotFound,  "Not Found")
}
//TokenAuthMiddleware token验证中间件(方法一)
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, r := range notoken {
			if strings.Contains(strings.ToLower(c.Request.URL.String()), strings.ToLower(r))  {
				c.Next()
				return
			}
		}
		conf.Jwt.Token = c.Request.Header.Get("api_token")
		if conf.Jwt.Token  == "" {
			apiresp.APIResult(c, http.StatusForbidden,  "Not token")
			return
		}
		user := &models.SysUser{}
		if err := conf.Jwt.ParseToken(user); err != nil {
			apiresp.APIResult(c, http.StatusBadRequest,  fmt.Sprintf("Token is Bad:%s", err.Error()))
			return
		}
		c.Next()
	}
}

//APITokenMiddleware token验证中间件(方法二)
func APITokenMiddleware(c *gin.Context) {
	for _, r := range notoken {
		if strings.Contains(strings.ToLower(c.Request.URL.String()), strings.ToLower(r))  {
			c.Next()
			return
		}
	}
	conf.Jwt.Token = c.Request.Header.Get("api_token")
	if conf.Jwt.Token == "" {
		apiresp.APIResult(c, http.StatusForbidden,  "权限不足")
		return
	}
	user := &models.SysUser{}
	if err :=conf.Jwt.ParseToken(user); err != nil {
		apiresp.APIResult(c, http.StatusBadRequest, fmt.Sprintf("Token is Bad:%s", err.Error()))
		return
	}
	//Todo：根据返回的用户查询操作权限范围
	c.Next()
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		resp:=Resp{}
		resp.APIOK(c,nil)
	})
	r.Run(":8080")
}