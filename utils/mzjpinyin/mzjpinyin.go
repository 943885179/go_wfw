package mzjpinyin

//使用的是网上的代码。不能识别数字和英文字母，会自动过滤掉

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

//单韵母
type vowel int32

var (
	ys = []vowel{'ā', 'ē', 'ī', 'ō', 'ū', 'ǖ', 'Ā', 'Ē', 'Ī', 'Ō', 'Ū', 'Ǖ'} // 单韵母 一声
	es = []vowel{'á', 'é', 'í', 'ó', 'ú', 'ǘ', 'Á', 'É', 'Í', 'Ó', 'Ú', 'Ǘ'} // 单韵母 二声
	ss = []vowel{'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ǚ', 'Ǎ', 'Ě', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǚ'} // 单韵母 三声
	fs = []vowel{'à', 'è', 'ì', 'ò', 'ù', 'ǜ', 'À', 'È', 'Ì', 'Ò', 'Ù', 'Ǜ'} // 单韵母 四声
	ws = []vowel{'a', 'e', 'i', 'o', 'u', 'v', 'A', 'E', 'I', 'O', 'U', 'V'} // 单韵母 无声调
)

// 从汉字到拼音的映射（带声调）
var (
	pinyinTemp map[vowel]interface{}
	toneTemp   map[vowel]vowel
)

const (
	//WithoutTone        string = "默认模式"      // 默认模式 例如：cao
	Tone               string = "带声调的拼音"    // 带声调的拼音 例如：Cào
	InitialsInCapitals string = "首字母大写不带声调" // 首字母大写不带声调 例如：Cao
	None               string = ""          //其他数字和字母
)

//读取拼音内容加入缓存
func get_VowelContent(filename string) (io.ReadCloser, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	return file, err
}
func DefaultInit() {
	LoadingPYFileName("pinyin.txt")
}

/*func init() {
	DefaultInit()
}*/
func LoadingPYFileName(filename string) {
	f, err := get_VowelContent(filename)
	pinyinTemp = make(map[vowel]interface{}, 0)
	toneTemp = make(map[vowel]vowel)
	for i, toneys := range ys {
		toneTemp[toneys] = ws[i]
	}
	for i, tonees := range es {
		toneTemp[tonees] = ws[i]
	}
	for i, toness := range ss {
		toneTemp[toness] = ws[i]
	}
	for i, tonefs := range fs {
		toneTemp[tonefs] = ws[i]
	}

	defer f.Close()
	if err != nil {
		panic(err) //自己看看路径是否正确，文件是否OK
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//单行分割取出拼音
		strs := strings.Split(scanner.Text(), "=>")
		if len(strs) < 2 {
			continue
		}
		i, err := strconv.ParseInt(strs[0], 16, 32)
		if err != nil {
			continue
		}
		/*t := vowel(i)
		fmt.Println(t)*/
		pinyinTemp[vowel(i)] = strs[1]
		/*var x = pinyinTemp
		fmt.Println(x)*/
	}
}

//汉字转换拼音
func Str2Pinyin(str string, split string, types string) (string, error) {
	hz := []vowel(str)
	words := make([]string, 0)
	for _, s := range hz {
		word, err := get_Vowel(s, types)
		if err != nil {
			return None, err
		}
		if len(word) > 0 {
			words = append(words, word)
		}
	}
	return strings.Join(words, split), nil
}

func get_Vowel(str vowel, types string) (string, error) {
	switch types {
	case Tone:
		return getTone(str), nil
	case InitialsInCapitals:
		return getInitialsInCapitals(str), nil
	default:
		return getDefault(str), nil
	}
}

func getTone(str vowel) string {
	if pinyinTemp[str] != nil {
		return pinyinTemp[str].(string)
	} else {
		return None
	}
	return None
}

//首字母大写不带声调
func getInitialsInCapitals(str vowel) string {
	def := getDefault(str)
	var objstr string
	if def == "" {
		return def
	}
	strarr := []vowel(def)
	if strarr[0] > 32 {
		strarr[0] = strarr[0] - 32
	}
	for _, v := range strarr {
		objstr += string(v)
	}
	return objstr
}

//带拼音字母转换为不带拼音的字母并返回拼音字母
func getDefault(str vowel) string {
	tone := getTone(str)
	var objstr string
	if tone == "" {
		return None
	}
	resultlen := make([]vowel, utf8.RuneCountInString(tone))
	count := 0
	for _, t := range tone {
		//fmt.Println(toneTemp,t,toneTemp[vowel(t)])
		changes, ok := toneTemp[vowel(t)] //有声调和无声调替换
		if ok {
			resultlen[count] = changes
		} else {
			resultlen[count] = vowel(t)
		}
		count++
	}
	for _, v := range resultlen {
		objstr += string(v)
	}
	//fmt.Println(objstr)
	return objstr
}

/*func main() {
	//LoadingPYFileName("./pinyin.txt")
	//DefaultInit()
	//fmt.Println(Str2Pinyin("是的啊a123", "", "")) //
	fmt.Println(strconv.ParseInt("1", 16, 32))
}
*/
