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
	"qshapi/utils/mzjimg"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjuuid"
	"strconv"
)

var (
	cliName="fileCli"
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
	//http://localhost:8705/static/upload/321988436372230144.png 访问图片
	service := conf.Services[cliName]
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
				"webconfig":conf.Services[cliName],
				"service":conf.Services[svName].Name,
			})
		})
		r.POST("upload", upload)
		r.POST("uploadMutiple", uploadMultiple)
		r.GET("showFile",showFile)
		r.GET("/fileById/:id",fileById)
		r.POST("ImgWH",ImgWH)
		//file.GET("getCaptcha", getCaptcha)
		//file.GET("verifyCaptcha", verifyCaptcha)
	}
	return g
}

func fileById(c *gin.Context){
	id,_:=  strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	 req := &file.FileId{
		Id: int64(id),
	} //c.Bind(req)
	//c.BindQuery(req)
	result, err := client.GetFile(context.Background(),req)
	if err!=nil{
		resp.APIError(c,err.Error())
		return
	}
	c.File(result.Path+"/"+result.Name)
}
func showFile(c *gin.Context){
	//http://localhost:8705/img?imageName=static/upload/321988436372230144.png
	url := c.Query("url")
	c.File(url)
}

type imgShow struct {
	Name string `json:"name"`
	W int `json:"w"`
	H int `json:"h"`
}

func ImgWH(c * gin.Context)  {
	i:=&imgShow{}
	c.Bind(i)
	fmt.Println(i)
	img, _ := mzjimg.ImageResizeImg(i.Name, i.W, i.H, 98)
	resp.APIOK(c,mzjimg.Img2Base64(img))
}
/*
func getImage(c *gin.Context){
	imageName := c.Query("imageName")
	file, _ := ioutil.ReadFile(imageName)
	c.Writer.WriteString(string(file))
}*/
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

func fileAttribute(c *gin.Context,f *multipart.FileHeader,sort int) *file.FileInfo {
	req:=&file.FileInfo{
	}
	c.Bind(req) //这里可以得到path(文件存放路径),file_type(文件业务类型)，file_explain（文件描述）
	req.FileSuffix=path.Ext(f.Filename)
	req.Id=mzjuuid.WorkerDefault()
	req.Name=strconv.Itoa(int(req.Id))+path.Ext(f.Filename)
	req.Size= f.Size
	req.Path=path.Join(conf.FilePath,req.Path)//文件夹前面加上文件系统路径
	req.Sort= int32(sort + 1)
	return req
}
//uploadFile 上传图片保存
func uploadFile(c *gin.Context, file *multipart.FileHeader,req *file.FileInfo)error  {
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