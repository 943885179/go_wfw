package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/proto/file"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjimg"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjuuid"
	"strconv"
	"strings"
)

var (
	cliName = "fileCli"
	svName  = "fileSrv"
	conf    models.APIConfig
	client  file.FileSrvService
	resp    = mzjgin.Resp{}
)

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}
func main() {
	//http://localhost:8705/static/upload/321988436372230144.png 访问图片
	service := conf.Services[cliName]
	cliName = service.Name
	svName = conf.Services[svName].Name
	client = file.NewFileSrvService(conf.Services[svName].Name, service.NewRoundSrv().Options().Client)
	s := service.NewGinWeb(SrvGin())
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

//图像显示：https://github.com/webp-sh/webp_server_go/releases，直接使用改程序然后做部署先、后续在此基础上做修改

/**
 * @Author mzj
 * @Description SrvGin 初始化file
 * @Date 上午 8:48 2020/10/13 0013
 * @Param
 * @return
 **/
func SrvGin() *gin.Engine {
	g := mzjgin.NewGin().Default(cliName)
	g.MaxMultipartMemory = 100
	r := g.Group("/")
	{
		r.GET("/", func(c *gin.Context) {
			resp.APIOK(c, "文件webapi")
		})
		r.POST("/", func(c *gin.Context) {
			resp.APIOK(c, gin.H{
				"webconfig": conf.Services[cliName],
				"service":   conf.Services[svName].Name,
			})
		})
		r.POST("upload", upload)
		r.POST("uploadMutiple", uploadMultiple)
		r.GET("showFile", showFile)
		r.GET("/fileById/:id", fileById)
		//r.GET("/fileDownload",fileDownload)
		r.GET("/ImgWH/:img", ImgWH)
		//file.GET("getCaptcha", getCaptcha)
		//file.GET("verifyCaptcha", verifyCaptcha)
	}
	return g
}

func fileById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	req := &dbmodel.Id{
		Id: int64(id),
	} //c.Bind(req)
	//c.BindQuery(req)
	result, err := client.GetFile(context.Background(), req)
	if err != nil {
		resp.APIError(c, err.Error())
		return
	}
	c.Request.RequestURI = c.Request.RequestURI + result.FileSuffix
	c.File(result.Path + "/" + result.Name) //因为id和name一致所以随便用一个就可以，如果是有后缀文件只能用name
	/*copyFile(result.Path+"/"+result.Name,result.FileSuffix)
	f, _ := ioutil.TempFile(result.Path, result.Name+result.FileSuffix)
	defer f.Close()
	bt, _ := mzjfile.File2Bytes(f)
	c.Writer.WriteString(string(bt))
	os.Remove(result.Path+"/"+result.Name+result.FileSuffix)*/
	//c.Redirect(http.StatusMovedPermanently,"../"+result.Path+"/"+result.Name+result.FileSuffix)

}

/**
 * @Author mzj
 * @Description 加载资源文件方法一
 * @Date 上午 8:39 2020/10/13 0013
 * @Param
 * @return
 **/
func showFile(c *gin.Context) {
	//http://localhost:8705/img?url=321988436372230144.png
	//["jpg","png","jpeg","bmp"]
	/*for _, ext := range AllowedTypes {
	haystack := strings.ToLower(ImgFilename)
	needle := strings.ToLower("." + ext)
	if strings.HasSuffix(haystack, needle) {
	allowed = true
	break
	} else {
	allowed = false
	}
	}
	if !allowed {
		c.Redirect()
		return
	}*/
	url := c.Query("url")
	c.File(path.Join(conf.FilePath, url))
}

/**
 * @Author mzj
 * @Description  加载资源文件方法二
 * @Date 上午 8:38 2020/10/13 0013
 * @Param
 * @return
 **/
func getImage(c *gin.Context) {
	url := c.Query("url")
	file, _ := ioutil.ReadFile(path.Join(conf.FilePath, url))
	c.Writer.WriteString(string(file))
}

/**
 * @Author mzj
 * @Description 资源（图片）压缩显示
 * @Date 上午 8:43 2020/10/13 0013
 * @Param
 * @return
 **/
func ImgWH(c *gin.Context) {
	//http://localhost:8705/ImgWH/321988436372230144_100*100
	imgstr := c.Param("img")
	s := strings.Split(imgstr, "_")
	id, _ := strconv.Atoi(s[0])
	req := &dbmodel.Id{
		Id: int64(id),
	}
	result, err := client.GetFile(context.Background(), req)
	if err != nil {
		resp.APIError(c, err.Error())
		return
	}
	wh := strings.Split(s[1], "*")
	w, _ := strconv.Atoi(wh[0])
	h, _ := strconv.Atoi(wh[1])
	img, _ := mzjimg.ImageResizeImg(result.Name, w, h, 0)
	png.Encode(c.Writer, img)
	//c.Writer.WriteString(string(img))
}

/**
 * @Author mzj
 * @Description 上传
 * @Date 上午 8:43 2020/10/13 0013
 * @Param
 * @return
 **/
func upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		resp.APIError(c, "请选择上传文件")
		return
	}
	req := fileAttribute(c, f, 0)
	if err := uploadFile(c, f, req); err != nil {
		resp.APIError(c, fmt.Sprintf("上传失败!%s", err.Error()))
		return
	}
	result, err := client.UploadFile(context.Background(), req)
	resp.MicroResp(c, result, err)
}

/**
 * @Author mzj
 * @Description  资源基本信息配置
 * @Date 上午 8:41 2020/10/13 0013
 * @Param
 * @return
 **/
func fileAttribute(c *gin.Context, f *multipart.FileHeader, sort int) *dbmodel.SysFile {
	req := &dbmodel.SysFile{}
	c.Bind(req) //这里可以得到path(文件存放路径),file_type(文件业务类型)，file_explain（文件描述）
	req.FileSuffix = path.Ext(f.Filename)
	req.Id = mzjuuid.WorkerDefault()
	req.Name = strconv.Itoa(int(req.Id)) + path.Ext(f.Filename) //这个后缀看情况吧，需要先解决下载doc，excel等问题
	req.Size = f.Size
	req.Path = path.Join(conf.FilePath, req.Path) //文件夹前面加上文件系统路径
	req.Sort = int32(sort + 1)
	return req
}

func webPImg() {
	webpPath := "webp" //压缩图存放路径
	if _, err := os.Stat(webpPath); os.IsNotExist(err) {
		os.MkdirAll(webpPath, os.ModePerm) // 先创建文件夹
		os.Chmod(webpPath, 0777)           // 再修改权限
	}

}

/**
 * @Author mzj
 * @Description 保存资源
 * @Date 上午 8:41 2020/10/13 0013
 * @Param
 * @return
 **/
func uploadFile(c *gin.Context, file *multipart.FileHeader, req *dbmodel.SysFile) error {
	if _, err := os.Stat(req.Path); os.IsNotExist(err) { // 必须分成两步创建文件夹
		//os.Mkdir(Config.FilePath, 0777)//创建单级目录
		os.MkdirAll(req.Path, os.ModePerm) // 先创建文件夹
		os.Chmod(req.Path, 0777)           // 再修改权限
	}
	return c.SaveUploadedFile(file, path.Join(req.Path, req.Name))
}

/**
 * @Author mzj
 * @Description 批量上传
 * @Date 上午 8:42 2020/10/13 0013
 * @Param
 * @return
 **/
func uploadMultiple(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		resp.APIError(c, fmt.Sprintf("读取多个文件错误!%s", err))
		return
	}
	files := form.File["file"]
	if len(files) == 0 {
		resp.APIError(c, "请选择上传文件")
		return
	}
	var result []interface{}
	for i, f := range files {
		req := fileAttribute(c, f, i)
		if err := uploadFile(c, f, req); err != nil {
			resp.APIError(c, fmt.Sprintf("上传失败!%s", err.Error()))
			return
		}
		r, err := client.UploadFile(context.Background(), req)
		if err != nil {
			resp.APIError(c, err.Error())
			return
		}
		result = append(result, r)
	}
	resp.APIOK(c, result)
}

/**
 * @Author mzj
 * @Description fileDownload下载资源
 * @Date 上午 9:04 2020/10/13 0013
 * @Param
 * @return
 **/
func fileDownload(c *gin.Context) {
	url := c.Query("url")
	//适用了gzip压缩，下面的设置不生效
	//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	//c.Writer.Header().Add("Content-Type", "application/octet-stream")
	file, _ := ioutil.ReadFile(url)
	c.Writer.WriteString(string(file))
	//c.File(url+".ini")
	//os.Remove(url+".ini")
}
func copyFile(url, suffix string) {
	srcFile, _ := os.Open(url)
	defer srcFile.Close()
	dstFile, _ := os.Create(url + suffix)
	defer dstFile.Close()
	io.Copy(dstFile, srcFile)
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
