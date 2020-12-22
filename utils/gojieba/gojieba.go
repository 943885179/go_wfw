package mzjgojieba

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
)
//全模式匹配
func CutAll(str string,newWords ...string)[]string{
	var seg = gojieba.NewJieba()
	defer seg.Free()
	for _, word := range newWords {
		seg.AddWord(word)
	}
	return seg.CutAll(str)
}
//Cut 精确匹配，
//str 字符串
//newWords 词库
func Cut(str string,newWords ...string)[]string{
	hmm:=true
	var seg = gojieba.NewJieba()
	defer seg.Free()
	for _, word := range newWords {
		seg.AddWord(word)
	}
	return seg.Cut(str,hmm)//精确模式
}
//搜索引擎模式
func CutForSearch(str string,newWords ...string)[]string{
	hmm:=true
	var seg = gojieba.NewJieba()
	defer seg.Free()
	for _, word := range newWords {
		seg.AddWord(word)
	}
	return seg.CutForSearch(str,hmm)//精确模式
}
//词性标注
func Tag(str string,newWords ...string)[]string{
	var seg = gojieba.NewJieba()
	defer seg.Free()
	for _, word := range newWords {
		seg.AddWord(word)
	}
	return seg.Tag(str)//精确模式
}
//TokenizeDefault  Tokenize Default Mode搜索引擎模式
func TokenizeDefault(str string, newWords ...string) []gojieba.Word {
	hmm:=false
	var seg = gojieba.NewJieba()
	defer seg.Free()
	for _, word := range newWords {
		seg.AddWord(word)
	}
	return seg.Tokenize(str, gojieba.DefaultMode, hmm)//精确模式
}
//TokenizeSearch Tokenize Search Mode 搜索引擎模式
func TokenizeSearch(str string, newWords ...string) []gojieba.Word {
	hmm:=false
	var seg = gojieba.NewJieba()
	defer seg.Free()
	for _, word := range newWords {
		seg.AddWord(word)
	}
	return seg.Tokenize(str, gojieba.SearchMode, hmm)//精确模式
}
func ExtractWithWeight(str string, topk int, newWords ...string) []gojieba.WordWeight {
	var seg = gojieba.NewJieba()
	defer seg.Free()
	for _, word := range newWords {
		seg.AddWord(word)
	}
	return seg.ExtractWithWeight(str,topk)//精确模式
}
func main()  {

	fmt.Println(strings.Join(CutForSearch("人生是一条射线，以我们的出生为起点，可以无限延伸。理想有多高远，学习有多勤奋，坚持有多长久，这条射线就有多长，我们的人生轨迹就有多深，价值就有多大，意义就有多远"),"|"))
}

/*
func main() {
	var seg = gojieba.NewJieba()
	defer seg.Free()
	var useHmm = false
	var separator = "|"
	var resWords []string
	var sentence = "万里长城万里长"
	resWords = seg.CutAll(sentence)
	fmt.Printf("%s\t全模式：%s \n", sentence, strings.Join(resWords, separator))
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t精确模式：%s \n", sentence, strings.Join(resWords, separator))
	var addWord = "万里长"
	seg.AddWord(addWord)
	fmt.Printf("添加新词：%s\n", addWord)
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t精确模式：%s \n", sentence, strings.Join(resWords, separator))
	sentence = "北京鲜花速递"
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t新词识别：%s \n", sentence, strings.Join(resWords, separator))
	sentence = "北京鲜花速递"
	resWords = seg.CutForSearch(sentence, useHmm)
	fmt.Println(sentence, "\t搜索引擎模式：", strings.Join(resWords, separator))
	sentence = "北京市朝阳公园"
	resWords = seg.Tag(sentence)
	fmt.Println(sentence, "\t词性标注：", strings.Join(resWords, separator))
	sentence = "鲁迅先生"
	resWords = seg.CutForSearch(sentence, !useHmm)
	fmt.Println(sentence, "\t搜索引擎模式：", strings.Join(resWords, separator))
	words := seg.Tokenize(sentence, gojieba.SearchMode, !useHmm)
	fmt.Println(sentence, "\tTokenize Search Mode 搜索引擎模式：", words)
	words = seg.Tokenize(sentence, gojieba.DefaultMode, !useHmm)
	fmt.Println(sentence, "\tTokenize Default Mode搜索引擎模式：", words)
	sentence = "万里长城万里长"
	word2 := seg.ExtractWithWeight(sentence, 5)
	fmt.Println(sentence, "\tExtract：", word2)
	return
}*/