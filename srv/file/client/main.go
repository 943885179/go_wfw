package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"mime/multipart"
	"os"
	"path"
	"qshapi/models"
	"qshapi/proto/file"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjuuid"
	"strconv"
)

var (
	webName="fileWeb"
	svName="fileSrv"
	conf models.APIConfig
	client file.FileSrvService
	resp =mzjgin.Resp{}
)
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
}
func main() {
	service := conf.Services[webName]
	client=file.NewFileSrvService(conf.Services[svName].Name,service.NewRoundSrv().Options().Client)
	s:= service.NewGinWeb(SrvGin())
	if err:=s.Run();err!= nil {
		log.Fatal(err)
	}
}
//SrvGin 初始化file
func SrvGin() *gin.Engine {
	g:=mzjgin.NewGin().Default()
	g.MaxMultipartMemory = 100
	r := g.Group("/")
	{
		r.GET("/", func(c *gin.Context) {
			resp.APIOK(c,"文件webapi")
		})
		r.POST("/", func(c *gin.Context) {
			resp.APIOK(c,gin.H{
				"webconfig":conf.Services[webName],
				"service":conf.Services[svName].Name,
			})
		})
		r.POST("upload", upload)
		r.POST("uploadMutiple", uploadMultiple)
		//file.GET("getCaptcha", getCaptcha)
		//file.GET("verifyCaptcha", verifyCaptcha)
	}

	return g
}

//上传
func upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		resp.APIError(c,"请选择上传文件")
		return
	}
	req:= fileAttribute(c,f,0)
	if err:=uploadFile(c,f,req);err != nil {
		resp.APIError(c, fmt.Sprintf("上传失败!%s", err.Error()))
		return
	}
	result, err := client.UploadFile(context.Background(),req)
	resp.MicroResp(c,result,err)
}

func fileAttribute(c *gin.Context,f *multipart.FileHeader,sort int) *file.FileReq {
	req:=&file.FileReq{
	}
	c.Bind(req) //这里可以得到path(文件存放路径),file_type(文件业务类型)，file_explain（文件描述）
	req.Name=strconv.Itoa(int(mzjuuid.WorkerDefault()))
	req.Size= f.Size
	req.FileSuffix=path.Ext(f.Filename)
	req.Path=path.Join(conf.FilePath,req.Path)//文件夹前面加上文件系统路径
	req.Sort= int32(sort + 1)
	return req
}
//uploadFile 上传图片保存
func uploadFile(c *gin.Context, file *multipart.FileHeader,req *file.FileReq)error  {
	if _, err := os.Stat(req.Path); os.IsNotExist(err) { // 必须分成两步创建文件夹
		//os.Mkdir(Config.FilePath, 0777)//创建单级目录
		os.MkdirAll(req.Path, os.ModePerm)// 先创建文件夹
		os.Chmod(req.Path, 0777) // 再修改权限
	}
	return c.SaveUploadedFile(file, path.Join(req.Path,req.Name))
}

//批量上传
func uploadMultiple(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		resp.APIError(c, fmt.Sprintf("读取多个文件错误!%s", err))
		return
	}
	files := form.File["file"]
	if len(files) ==0 {
		resp.APIError(c, "请选择上传文件")
		return
	}
	var result []interface{}
	for i, f := range files {
		req:= fileAttribute(c,f,i)
		if err:=uploadFile(c,f,req);err != nil {
			resp.APIError(c, fmt.Sprintf("上传失败!%s", err.Error()))
			return
		}
		r, err := client.UploadFile(context.Background(), req)
		if err != nil {
			resp.APIError(c,err.Error())
			return
		}
		result = append(result, r)
	}
	resp.APIOK(c,result)
}
/*
//生成验证码
func getCaptcha(c *gin.Context) {
	var u captcha.ConfigJSONBody
	//c.ShouldBindJSON(u)
	c.ShouldBind(&u)
	fmt.Println("读取到的数据", u)
	id, b64s, err := captcha.NewCaptcha(u)
	if err != nil {
		apiresp.APIError(c, err.Error())
		return
	}
	result := map[string]interface{}{"b64s": b64s, "captchaId": id}
	apiresp.APIOK(c, result)
}

//验证验证码是否正确
func verifyCaptcha(c *gin.Context) {
	var u captcha.ConfigJSONBody
	//c.ShouldBindJSON(u)
	c.ShouldBind(&u)
	if u.ID == "" || u.VerifyValue == "" {
		apiresp.APIError(c, "请输入验证码")
	}
	b := captcha.QshCaptchaVerifyHandle(u.ID, u.VerifyValue)
	apiresp.APIOK(c, b)
}
*/
//Base64Upload 稍等
func Base64Upload(context *gin.Context) {
}