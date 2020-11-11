package tencentcloud

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"qshapi/utils/mzjstruct"
)

var (
	client *ocr.Client
)

//TxOcrAPI 腾讯文字OrC
type TxOcrAPI struct {
	Region    string `json:"region"`
	SecretID  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Endpoint  string `json:"endpoint"`
	IsDebug   bool   `json:"isDebug"` //是否为调试模式
}

//OCRConfig 图像基础设置
type OCRConfig struct {
	ImageBase64 string `json:"ImageBase64,omitempty"` //图片base64
	ImageURL    string `json:"ImageUrl,omitempty"`    //图片地址

	Scene        string `json:"Scene,omitempty"`
	LanguageType string `json:"LanguageType,omitempty"`

	CardSide string `json:"CardSide,omitempty"`
	Config   string `json:"Config,omitempty"`
}

//TxOcrInit 初始化
func (c TxOcrAPI) TxOcrInit() {
	credential := common.NewCredential(c.SecretID, c.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = c.Endpoint
	clients, err := ocr.NewClient(credential, c.Region, cpf)
	if err != nil {
		log.Info(fmt.Sprintf("腾讯ocrsdk初始化失败,%s", err.Error()))
	}
	client = clients
	fmt.Println("初始化成功", client)
}

//GeneralOCR 文字识别
func (c TxOcrAPI) GeneralOCR(config OCRConfig, tp int) (resp *ocr.GeneralBasicOCRResponse, err errors.TencentCloudSDKError) {
	params, _ := json.Marshal(config)
	request := ocr.NewGeneralBasicOCRRequest()
	request.FromJsonString(string(params))
	resp, x := client.GeneralBasicOCR(request)
	mzjstruct.CopyStruct(&x, &err)
	return resp, err
	/*switch tp {
	case 1:
		request = ocr.NewGeneralFastOCRRequest() //高速版
	case 2:
		request = ocr.NewGeneralEfficientOCRRequest() //精简版
	case 3:
		request = ocr.NewGeneralAccurateOCRRequest() //高精度
	default:
		request = ocr.NewGeneralBasicOCRRequest() //基础版本
	}*/
	/*response, err := client.GeneralBasicOCR(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return resp, err
	}
	if err != nil {
		return resp, err
	}
	response.Response.TextDetections
	//fmt.Printf("%s", response.ToJsonString())*/
}

//IDCardOCR 身份证验证
func (c TxOcrAPI) IDCardOCR(config OCRConfig) (resp *ocr.IDCardOCRResponse, err errors.TencentCloudSDKError) {
	request := ocr.NewIDCardOCRRequest()
	params, _ := json.Marshal(config)
	request.FromJsonString(string(params))
	resp, x := client.IDCardOCR(request)
	mzjstruct.CopyStruct(&x, &err)
	return resp, err
}

//Bizlicense 营业执照
func (c TxOcrAPI) Bizlicense(config OCRConfig) (resp *ocr.BizLicenseOCRResponse, err errors.TencentCloudSDKError) {
	request := ocr.NewBizLicenseOCRRequest()
	params, _ := json.Marshal(config)
	request.FromJsonString(string(params))
	resp, x := client.BizLicenseOCR(request)
	mzjstruct.CopyStruct(&x, &err)
	return resp, err
}

//EnterpriseLicense 企业证照识别 支持智能化识别各类企业登记证书、许可证书、企业执照、三证合一类证书，结构化输出统一社会信用代码、公司名称、法定代表人、公司地址、注册资金、企业类型、经营范围等关键字段
func (c TxOcrAPI) EnterpriseLicense(config OCRConfig) (resp *ocr.EnterpriseLicenseOCRResponse, err errors.TencentCloudSDKError) {
	request := ocr.NewEnterpriseLicenseOCRRequest()
	params, _ := json.Marshal(config)
	request.FromJsonString(string(params))
	resp, x := client.EnterpriseLicenseOCR(request)
	mzjstruct.CopyStruct(&x, &err)
	return resp, err
}

//BanckCard 银行卡识别
func (c TxOcrAPI) BanckCard(config OCRConfig) (resp *ocr.BankCardOCRResponse, err errors.TencentCloudSDKError) {
	request := ocr.NewBankCardOCRRequest()
	params, _ := json.Marshal(config)
	request.FromJsonString(string(params))
	resp, x := client.BankCardOCR(request)
	mzjstruct.CopyStruct(&x, &err)
	return resp, err
}
func main() {

	credential := common.NewCredential(
		"AKIDnQsERTVtSuTtAenY2LTbZN4aeS5YuMYX ",
		"0d7NuAGF4k4veDmCSXSfnlYQA91Tr6BN",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-beijing", cpf)

	request := ocr.NewGeneralBasicOCRRequest()

	params := "{\"ImageUrl\":\"https://ocr-demo-1254418846.cos.ap-guangzhou.myqcloud.com/card/IDCardOCR/IDCardOCR1.jpg\"}"
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.GeneralBasicOCR(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
