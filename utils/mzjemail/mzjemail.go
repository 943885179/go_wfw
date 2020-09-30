package mzjemail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"reflect"
)


type EmailType int
const (
	//QQ邮箱
	QQ EmailType =iota
	//QQ企业邮箱
	QQC
	//163.com:
	E163
	//谷歌邮箱
	GMAIL

)

type EmailServer struct {
	Host string `json:"host"`//服务器地址
	Port int `json:"port"`// 端口
}
func (t EmailType) Server() *EmailServer{
	var result EmailServer
	switch t {
	case QQ:
		result= EmailServer{Host: "smtp.qq.com", Port: 25}
	case QQC:
		result= EmailServer{Host: "smtp.exmail.qq.com", Port: 587} //SMTP服务器地址：smtp.exmail.qq.com（SSL启用 端口：587/465）
	case E163:
		result= EmailServer{Host: "smtp.163.com", Port: 25}
	case GMAIL:
		result= EmailServer{Host: "smtp.gmail.com", Port: 587}
	}
	return &result
}

type EmailConfig struct {
	EmailType EmailType `json:"emailType"`
	UserName string `json:"userName"`
	Password string	`json:"password"`
	Bcc []string `json:"bcc"`
	Cc []string `json:"cc"`
	To          []string `json:"to"`
	Subject     string `json:"subject"`
	Text       string `json:"text"`// Plaintext message (optional)
	HTML        string `json:"html"`// Html message (optional)
	AttachFiles []string `json:"attachFiles"`//附件
}
func (conf *EmailConfig)Copy()*EmailConfig{
	newConf:=EmailConfig{}
	sval := reflect.ValueOf(conf).Elem()
	dval := reflect.ValueOf(&newConf).Elem()
	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
	}
	return &newConf
}
func (conf *EmailConfig) Send() error {
	fmt.Println("发送有限的配置",conf)
	e:=email.NewEmail()
	e.From = conf.UserName
	e.To =conf.To
	e.Bcc =conf.Bcc
	e.Cc = conf.Cc
	e.Subject = conf.Subject
	e.Text = []byte(conf.Text)
	e.HTML = []byte(conf.HTML)
	for _, file := range conf.AttachFiles {
		e.AttachFile(file)
	}
	return e.Send(fmt.Sprintf("%s:%d",conf.EmailType.Server().Host,conf.EmailType.Server().Port), smtp.PlainAuth("", conf.UserName, conf.Password, conf.EmailType.Server().Host))
}


func main(){
	c:=EmailConfig{
		EmailType: 0,
		UserName: "943885179@qq.com",
		Password: "ppiqtqrtrzdpbcjf",
		To: []string{"1603044069@qq.com"},
		Subject: "这个是主题",
		Text: "你好",
	}
	err:=c.Send()
	fmt.Println(err)
}