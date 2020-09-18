package tencentcloud


import (
"encoding/json"
"fmt"
	log "github.com/sirupsen/logrus"
	"qshapi/models"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

var (
	client *ocr.Client
)

//TxOcrInit 初始化
func TxOcrInit(config models.TxOcrAPI) {
	credential := common.NewCredential(config.SecretID, config.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.Endpoint
	clients, err := ocr.NewClient(credential, config.Region, cpf)
	if err != nil {
		log.Info(fmt.Sprintf("腾讯ocrsdk初始化失败,%s", err.Error()))
	}
	client = clients
	fmt.Println("初始化成功", client)
}

//GeneralOCR 文字识别
func GeneralOCR(config OCRConfig, tp int) (resp *ocr.GeneralBasicOCRResponse, err error) {
	params, err := json.Marshal(config)
	if err != nil {
		return resp, err
	}
	request := ocr.NewGeneralBasicOCRRequest()
	err = request.FromJsonString(string(params))
	if err != nil {
		return resp, err
	}
	return client.GeneralBasicOCR(request)
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
func IDCardOCR(config OCRConfig) (resp *ocr.IDCardOCRResponse, err error) {
	request := ocr.NewIDCardOCRRequest()
	params, err := json.Marshal(config)
	if err != nil {
		return resp, err
	}
	err = request.FromJsonString(string(params))
	if err != nil {
		return resp, err
	}
	fmt.Println("当前客户端检测", client)
	return client.IDCardOCR(request)
}

//Bizlicense 营业执照
func Bizlicense(config OCRConfig) (resp *ocr.BizLicenseOCRResponse, err error) {
	request := ocr.NewBizLicenseOCRRequest()
	params, err := json.Marshal(config)
	if err != nil {
		return resp, err
	}
	err = request.FromJsonString(string(params))
	if err != nil {
		return resp, err
	}
	return client.BizLicenseOCR(request)
}

//EnterpriseLicense 企业证照识别 支持智能化识别各类企业登记证书、许可证书、企业执照、三证合一类证书，结构化输出统一社会信用代码、公司名称、法定代表人、公司地址、注册资金、企业类型、经营范围等关键字段
func EnterpriseLicense(config OCRConfig) (resp *ocr.EnterpriseLicenseOCRResponse, err error) {
	request := ocr.NewEnterpriseLicenseOCRRequest()
	params, err := json.Marshal(config)
	if err != nil {
		return resp, err
	}
	err = request.FromJsonString(string(params))
	if err != nil {
		return resp, err
	}
	return client.EnterpriseLicenseOCR(request)
}

//BanckCard 银行卡识别
func BanckCard(config OCRConfig) (resp *ocr.BankCardOCRResponse, err error) {
	request := ocr.NewBankCardOCRRequest()
	params, err := json.Marshal(config)
	if err != nil {
		return resp, err
	}
	err = request.FromJsonString(string(params))
	return client.BankCardOCR(request)
}

//OCRConfig 图像基础设置
type OCRConfig struct {
	ImageBase64 string `json:"ImageBase64,omitempty"`
	ImageURL    string `json:"ImageUrl,omitempty"`

	Scene        string `json:"Scene,omitempty"`
	LanguageType string `json:"LanguageType,omitempty"`

	CardSide string `json:"CardSide,omitempty"`
	Config   string `json:"Config,omitempty"`
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

