package mzjjwt


import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/goinggo/mapstructure"
)

type Jwt struct {
	Secret string `json:"secret"`//jwt加密字段
	TimeOut    time.Duration `json:"timeOut"`    //过期时长
	Iss string `json:"-"`//用户
	Token   string `json:"token"`    //token
	Data  interface{}     `json:"data"`//携带的其他信息
}
//CreateToken 创建Token
func (t *Jwt)CreateToken() (tokenValue string,err error) {
	//自定义claim iss: 签发者 sub: 面向的用户 aud: 接收方 exp: 过期时间 nbf: 生效时间 iat: 签发时间 jti: 唯一身份标识
	//t.TimeOut = time.Minute * time.Duration(30)
	//t.CreateTime = time.Now().Unix()
	//t.OutTime = time.Now().Add(t.TimeOut).Unix()
	claim := make(jwt.MapClaims)
	claim["exe"] = time.Now().Add(t.TimeOut*time.Second).Unix()  //过期时间设定
	claim["iat"] =  time.Now().Unix() //创建时间
	claim["nbf"] =  time.Now().Unix() //生效时间
	claim["iss"] = t.Iss
	claim["data"] = t.Data
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claim
	tokenValue, err = token.SignedString([]byte(t.Secret))
	return  tokenValue, err
}

//ParseToken 解析token
func (t *Jwt)ParseToken(resp interface{}) (err error) {
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
	fmt.Sprintln("123")
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
	tk:=Jwt{
		Secret: secret,
		TimeOut:  ts,
		Data: tu,
	}
	tk.Token,_=tk.CreateToken()
	fmt.Println(tk.Token)
	tus:=TokenUser{}
	tk.ParseToken(&tus)
	fmt.Println(tus)
}