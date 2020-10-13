package mzjimg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/goki/freetype"
	"golang.org/x/image/math/fixed"
)

//AddWatermarkByImg 添加水印图片
func AddWatermarkByImg(fromPath, watPath, resultPath string, position, option int) error {
	//操作原始图片
	imgb, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	var img image.Image
	if strings.HasSuffix(imgb.Name(), "jpg") {
		img, err = jpeg.Decode(imgb)
		if err != nil {
			return err
		}
	}
	if strings.HasSuffix(imgb.Name(), "png") {
		img, err = png.Decode(imgb)
		if err != nil {
			return err
		}
	}
	defer imgb.Close()
	//水印图片
	wat, err := os.Open(watPath)
	if err != nil {
		return err
	}
	var watimg image.Image
	if strings.HasSuffix(wat.Name(), "jpg") {
		watimg, err = jpeg.Decode(wat)
		if err != nil {
			return err
		}
	}
	if strings.HasSuffix(wat.Name(), "png") {
		watimg, err = png.Decode(wat)
		if err != nil {
			return err
		}
	}
	//把水印写到右下角，并向0坐标各偏移10个像素
	var offset image.Point
	switch position {
	case 0:
		offset = image.Pt(img.Bounds().Dx()-watimg.Bounds().Dx(), img.Bounds().Dy()-watimg.Bounds().Dy()) //右下
	case 1:
		offset = image.Pt(0, 0) //左上
	case 2:
		offset = image.Pt(img.Bounds().Dx()-watimg.Bounds().Dx(), 0) //右上
	case 3:
		offset = image.Pt(0, img.Bounds().Dy()-watimg.Bounds().Dy()) //左下
	default:
		offset = image.Pt(img.Bounds().Dx()-watimg.Bounds().Dx(), img.Bounds().Dy()-watimg.Bounds().Dy())
	}
	//创建新图
	m := image.NewRGBA(img.Bounds())
	draw.Draw(m, img.Bounds(), img, image.ZP, draw.Src)

	draw.Draw(m, watimg.Bounds().Add(offset), watimg, image.ZP, draw.Over)
	//生成新的图片
	imgw, _ := os.Create(resultPath)
	//设置模糊程度1-100
	jpeg.Encode(imgw, m, &jpeg.Options{option})
	defer imgw.Close()
	return nil
}

//AddWatermarkByFont 添加水印文字
func AddWatermarkByFont(fromPath, str, resultPath string, position, option int) error {
	//操作原始图片
	imgb, err := os.Open(fromPath)
	defer imgb.Close()
	if err != nil {
		return err
	}
	var img image.Image
	if strings.HasSuffix(imgb.Name(), "jpg") {
		img, err = jpeg.Decode(imgb)
		if err != nil {
			return err
		}
	}
	if strings.HasSuffix(imgb.Name(), "png") {
		img, err = png.Decode(imgb)
		if err != nil {
			return err
		}
	}

	imgnew := image.NewRGBA(img.Bounds())
	for y := 0; y < imgnew.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			imgnew.Set(x, y, img.At(x, y))
		}
	}
	//拷贝一个字体文件到运行目录
	fontBytes, err := ioutil.ReadFile("lib/simsun.ttc")
	if err != nil {
		log.Println(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}

	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(font)
	f.SetFontSize(12)
	f.SetClip(img.Bounds())
	f.SetDst(imgnew)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
	var pt fixed.Point26_6
	pt = freetype.Pt(img.Bounds().Dx()-len(str)*5, img.Bounds().Dy()-2)
	switch position {
	case 0:
		pt = freetype.Pt(img.Bounds().Dx()-len(str)*5, img.Bounds().Dy()-2) //右下
	case 1:
		pt = freetype.Pt(2, 12) //左上
	case 2:
		pt = freetype.Pt(img.Bounds().Dx()-len(str)*5, 12) //右上
	case 3:
		pt = freetype.Pt(2, img.Bounds().Dy()-2) //左下
	default:
		pt = freetype.Pt(img.Bounds().Dx()-len(str)*5, img.Bounds().Dy()-2)
	}
	fmt.Println(len(str))
	_, err = f.DrawString(str, pt)
	//保存到新文件中
	newfile, _ := os.Create(resultPath)
	defer newfile.Close()
	err = jpeg.Encode(newfile, imgnew, &jpeg.Options{option})
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

//ImageResize 压缩图片
func ImageResize(imgPath, resultPath string, width, height, ts int) error {
	image, err := ImageResizeImg(imgPath, width, height, ts)
	err = imaging.Save(image, resultPath)
	return err
}

/**
 * @Author mzj
 * @Description 图片剪切
 * @Date 上午 10:47 2020/10/13 0013
 * @Param imgPath string 图片路径, width 长度, height 高度, ts 压缩类型 int
 * @return
 **/
func ImageResizeImg(imgPath string, width, height, ts int) (image image.Image, err error) {
	imgData, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return image, err
	}
	buf := bytes.NewBuffer(imgData)
	image, err = imaging.Decode(buf)
	if err != nil {
		return image, err
	}
	//生成缩略图
	switch ts {
	case 0:
		image = imaging.Resize(image, width, height, imaging.Lanczos) //正常比例压缩,如果只输入x或y那么另外一边等比例压缩
	case 1:
		image = imaging.Fill(image, width, height, imaging.Center, imaging.Lanczos) //中间压缩,四周过滤掉
	case 2:
		image = imaging.Fill(image, width, height, imaging.Top, imaging.Lanczos) //以上边为中心,多余剪切
	case 3:
		image = imaging.Fill(image, width, height, imaging.TopLeft, imaging.Lanczos)
	case 4:
		image = imaging.Fill(image, width, height, imaging.TopRight, imaging.Lanczos)
	case 5:
		image = imaging.Fill(image, width, height, imaging.Left, imaging.Lanczos)
	case 6:
		image = imaging.Fill(image, width, height, imaging.Right, imaging.Lanczos)
	case 7:
		image = imaging.Fill(image, width, height, imaging.Bottom, imaging.Lanczos)
	case 8:
		image = imaging.Fill(image, width, height, imaging.BottomLeft, imaging.Lanczos)
	case 9:
		image = imaging.Fill(image, width, height, imaging.BottomRight, imaging.Lanczos)
	default:
		image = imaging.Resize(image, width, height, imaging.Lanczos)
	}
	return image, err
}

/**
 *图片转base64
 * @Author Administrator
 * @Description //TODO 
 * @Date  
 * @Param
 * @return 
 **/
func Img2Base64(m image.Image) string{
	buf:=bytes.NewBuffer(nil)
	jpeg.Encode(buf,m,nil)
	dist:=make([]byte,10)
	base64.StdEncoding.Encode(dist,buf.Bytes())
	return string(dist)
}

/**
 *base64转byte
 * @Author Administrator
 * @Description //TODO
 * @Date
 * @Param
 * @return
 **/
func Base642Byte(base64Str string) []byte {
	b,_:=base64.StdEncoding.DecodeString(base64Str)
	return b
}
func Base642Buff(base64Str string) *bytes.Buffer {
	b,_:=base64.StdEncoding.DecodeString(base64Str)
	return bytes.NewBuffer(b)
}
