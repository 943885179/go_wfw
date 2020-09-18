package mzjjwt


import (
	"errors"
	"fmt"
	"qshapi/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/goinggo/mapstructure"
)

//Token token
type Token struct {
	//ID         int    `json:"id"`
	Secret string `json:"-"` //token加密字段不返回给前端
	Token      string        `json:"token"`      //token
	CreateTime int64         `json:"createTime"` //创建时间
	OutTime    int64         `json:"outTime"`    //过期时间
	TimeOut    time.Duration `json:"timeOut"`    //过期时长
	Iss string `json:"-"`//用户
	Data  interface{}     `json:"data"`//携带的其他信息
}
//var secret string = "weixiao_token_secret" //token加密字段

func NewToken(tk models.Jwt) *Token {
	c:=Token{
		Secret: tk.Secret,
		TimeOut:  tk.TimeOut,
		Token: tk.Token,
		CreateTime : time.Now().Unix(),
		OutTime: time.Now().Add(tk.TimeOut).Unix(),
	}
	return &c
}

//CreateToken 创建Token
func (t *Token)CreateToken() (tokenValue string,err error) {
	//自定义claim iss: 签发者 sub: 面向的用户 aud: 接收方 exp: 过期时间 nbf: 生效时间 iat: 签发时间 jti: 唯一身份标识
	//t.TimeOut = time.Minute * time.Duration(30)
	//t.CreateTime = time.Now().Unix()
	//t.OutTime = time.Now().Add(t.TimeOut).Unix()
	claim := make(jwt.MapClaims)
	claim["exe"] = t.OutTime    //过期时间设定
	claim["iat"] = t.CreateTime //创建时间
	claim["nbf"] = t.CreateTime //生效时间
	claim["iss"] = t.Iss
	claim["data"] = t.Data
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claim
	tokenValue, err = token.SignedString([]byte(t.Secret))
	return  tokenValue, err
}

//ParseToken 解析token
func (t *Token)ParseToken(resp interface{}) (err error) {
	tk := func() jwt.Keyfunc {
		return func(token *jwt.Token) (interface{}, error) {
			return []byte(t.Secret), nil
		}
	}
	token, err := jwt.Parse(t.Token, tk())
	if err != nil {
		return err
	}
	//验证token，检查token是否被修改
	if !token.Valid {
		err = errors.New("token is invalid")
		return err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := errors.New("cannot convert claim to mapclaim")
		return err
	}
	return mapstructure.Decode(claim["data"],resp)
}
func main(){
	//TokenUser 用户基本信息（展示用）
	type TokenUser struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
		Email  string `json:"email"`
	}
	var secret = "weixiao_token_secret" //token加密字段
	ts:=time.Minute * time.Duration(30)
	tu:= TokenUser{
		Name: "Test",
		Avatar: "asd",
		Email: "943885179@qq.com",
	}
	tk:=Token{
		Secret: secret,
		TimeOut:  ts,
		CreateTime : time.Now().Unix(),
		OutTime: time.Now().Add(ts).Unix(),
		Data: tu,
	}
	tk.Token,_=tk.CreateToken()
	fmt.Println(tk.Token)
	tus:=TokenUser{}
	tk.ParseToken(&tus)
	fmt.Println(tus)
}