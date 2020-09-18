package mzjjson
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
//JSONRead 读取json文件返回string
func JSONRead(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("读取文件错误:%s", err))
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(fmt.Sprintf("读取文件错误:%s", err))
	}
	content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))
	//fmt.Println(string(content))
	return string(content)
}
//JSONReadEntity 读取json文件返回实体
func JSONReadEntity(path string, resp interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("读取文件错误:%s", err))
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(fmt.Sprintf("读取文件错误:%s", err))
	}
	//fmt.Println(string(content))
	content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}
	return json.Unmarshal([]byte(string(content)), resp)
}
//EntityToJSONOS 将实体写入json
func EntityToJSONOS(path string, t interface{}) {
	// 最后面4个空格，让json格式更美观
	result, _ := json.MarshalIndent(t, "", "    ")
	fmt.Println(string(result))
	_ = ioutil.WriteFile(path, result, 0644)
}
//StringToJSONOS 将实体写入json
func StringToJSONOS(path string, str string) {
	_ = ioutil.WriteFile(path, []byte(str), 0644)
}

//ObjectToJson interface转json
func Marshal(data interface{})(string,error)  {
	bt,err:= json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bt),err
}
//UnMarshal 字符串转interface
func UnMarshal(str string,resp interface{})error  {
	return json.Unmarshal([]byte(str),resp)
}