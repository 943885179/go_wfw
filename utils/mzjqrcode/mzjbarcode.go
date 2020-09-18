package mzjqrcode
import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"qshapi/utils/mzjimg"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
)

//NeWQrCode 创建新的二维码推荐使用
func NeWQrCode(context, qrPath string, size int) error {

	qrCode, err := qr.Encode(context, qr.M, qr.Auto)
	if err != nil {
		return err
	}
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return err
	}
	file, err := os.Create(qrPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, qrCode)
}

//NeWQrCodeImg 创建新的二维码推荐使用
func NeWQrCodeImg(context, qrPath, avatarPath string, size int) error {
	qrCode, err := qr.Encode(context, qr.M, qr.Auto)
	if err != nil {
		return err
	}
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return err
	}
	//读取中间图片
	avimg, err := mzjimg.ImageResizeImg(avatarPath, size/5, size/5, 1)
	//居中设置
	offset := image.Pt((qrCode.Bounds().Max.X-avimg.Bounds().Max.X)/2, (qrCode.Bounds().Max.Y-avimg.Bounds().Max.Y)/2)
	m := image.NewRGBA(qrCode.Bounds())
	draw.Draw(m, qrCode.Bounds(), qrCode, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, avimg.Bounds().Add(offset), avimg, image.Point{X: 0, Y: 0}, draw.Over)
	i, err := os.Create(qrPath)
	if err != nil {
		return err
	}
	if strings.HasSuffix(qrPath, ".png") {
		err = png.Encode(i, m)
		if err != nil {
			return err
		}
	}
	if strings.HasSuffix(qrPath, ".jpg") {
		err = jpeg.Encode(i, m, &jpeg.Options{100})
		if err != nil {
			return err
		}
	}
	return nil
	/*file, err := os.Create(qrPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, qrCode)*/
}

//ReadQrCode 读取二维码
func ReadQrCode(qrPath string) (result string, err error) {
	fi, err := os.Open(qrPath)
	if err != nil {
		return result, err
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		return result, err
	}
	result = qrmatrix.Content
	return result, err
}
