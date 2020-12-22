package mzjpinyin

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"unicode/utf8"
)

//go get -u github.com/mozillazg/go-pinyin/cmd/pinyin 同样会自动过滤掉英文和数字
//默认 [[zhong] [guo] [ren]]
func Pinyin(str string) [][]string {
	a := pinyin.NewArgs()
	return pinyin.Pinyin(str, a)
}

// 包含声调 [[zhōng] [guó] [rén]]
func PinyinTone(str string) [][]string {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone
	return pinyin.Pinyin(str, a)
}

// 包含声调[[zho1ng] [guo2] [re2n]]
func PinyinTone2(str string) [][]string {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone2
	return pinyin.Pinyin(str, a)
}

// 开启多音字模式[[zhong zhong] [guo] [ren]]
func PinyinHeteronym(str string) [][]string {
	a := pinyin.NewArgs()
	a.Heteronym = true
	return pinyin.Pinyin(str, a)
}

// 开启多音字模式 [[zho1ng zho4ng] [guo2] [re2n]]
func PinyinHeteronymTone(str string) [][]string {
	a := pinyin.NewArgs()
	a.Heteronym = true
	a.Style = pinyin.Tone
	return pinyin.Pinyin(str, a)
}

// 开启多音字模式 [[zho1ng zho4ng] [guo2] [re2n]]
func PinyinHeteronymTone2(str string) [][]string {
	a := pinyin.NewArgs()
	a.Heteronym = true
	a.Style = pinyin.Tone2
	return pinyin.Pinyin(str, a)
}

//[zhong guo ren]
func LazyPinyin(str string) []string {
	return pinyin.LazyPinyin(str, pinyin.NewArgs())
}
func LazyConvert(str string) []string {
	return pinyin.LazyConvert(str, nil)
}

// [[zhong] [guo] [ren]]
func Convert(str string) [][]string {
	return pinyin.Convert(str, nil)
}

func DefalutStrToPinyin(str string) string {
	return strings.Join(LazyPinyin(str), "")
}

func DefalutStrToPinyinSZM(str string) string {
	arr := LazyPinyin(str)
	var result = ""
	for _, s := range arr {
		_, size := utf8.DecodeRuneInString(s)
		result = result + strings.ToUpper(s[:size])
	}
	return result
}

//返回用-份额的字符串
func Slug(str string) string {
	a := pinyin.NewArgs()
	return pinyin.Slug(str, a)
}
func main() {
	var str = "我是weixiao123m,啥"
	fmt.Println(DefalutStrToPinyinSZM(str)) //去掉了数字和英文
	fmt.Println(DefalutStrToPinyin(str))    //去掉了数字和英文
	// 段落转换，支持完整支持多音字，保留符号
	//fmt.Println(pinyin.Paragraph("交给团长，告诉他我们给予期望。前线的供给一定要能自给自足！"))
	fmt.Println(pinyin.Convert(str, nil))
	fmt.Println(PinyinTone(str))
	fmt.Println(PinyinTone2(str))
	fmt.Println(PinyinHeteronym(str))
	fmt.Println(PinyinHeteronymTone(str))
	fmt.Println(PinyinHeteronymTone2(str))
	fmt.Println(Convert(str))
	fmt.Println(LazyConvert(str))
	fmt.Println(Slug(str))
}
