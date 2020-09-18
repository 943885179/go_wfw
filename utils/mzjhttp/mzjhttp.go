package mzjhttp

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//Post post请求，返回string
func Post(url string, data interface{}) (string, int, error) {
	fmt.Printf("请求地址：%s\n", url) //读取请求
	contentType := "application/json;charset=utf-8"
	jsonStr, err := json.Marshal(data)
	fmt.Println(string(jsonStr))
	if err != nil {
		return "", 500, err
	}
	//fmt.Printf("请求地址：%s\n请求参数：%s\n", url, string(jsonStr)) //读取请求

	resp, err := http.Post(url, contentType, strings.NewReader(string(jsonStr)))
	if err != nil {
		//dbUtil.Exec(entity.EntityToAddSQL(HTTPLog, ""))
		return "", 500, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//dbUtil.Exec(entity.EntityToAddSQL(HTTPLog, ""))
			return "", 500, err
		}
		//fmt.Printf("返回内容：%s\n", string(content)) //返回内容
		//dbUtil.Exec(entity.EntityToAddSQL(HTTPLog, ""))
		log.Info(fmt.Sprintf("请求地址：%s\n请求参数：%s\n返回参数:%s\n", url, string(jsonStr), string(content)))
		return string(content), resp.StatusCode, nil
	} else if resp.StatusCode == 504 { //超时
		return "", resp.StatusCode, nil
	} else {
		return "", resp.StatusCode, fmt.Errorf("返回状态：%d，返回内容为:%v", resp.StatusCode, resp.Body)
	}
}

//Get get请求，返回string
func Get(url string) (string, int, error) {
	fmt.Printf("请求地址：%s\n", url) //读取请求
	resp, err := http.Get(url)
	if err != nil {
		//dbUtil.Exec(entity.EntityToAddSQL(HTTPLog, ""))
		return "", 500, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//dbUtil.Exec(entity.EntityToAddSQL(HTTPLog, ""))
		return "", 500, err
	}
	//fmt.Printf("返回内容：%s\n", string(content)) //返回内容
	//dbUtil.Exec(entity.EntityToAddSQL(HTTPLog, ""))
	log.Info(fmt.Sprintf("请求地址：%s\n返回参数:%s\n", url, string(content)))
	return string(content), resp.StatusCode, nil
}

//PostEntity post请求返回实体
func PostEntity(url string, reqData interface{}, resp interface{}) error {
	content, code, err := Post(url, reqData)
	if err != nil || code == 504 { //超时重试
		fmt.Printf("请求失败，开始重试")
		for i := 0; i < 4; i++ {
			fmt.Printf("正在重试%d...", i)
			time.Sleep(time.Second * 5)
			content, code, err = Post(url, reqData)
			if err == nil {
				break
			}
		}
	}
	if err != nil {
		fmt.Println("重试超时，请检查...")
		return err
	}
	return json.Unmarshal([]byte(string(content)), resp)
}

//GetEntity get请求返回实体
func GetEntity(url string, resp interface{}) error {
	content, code, err := Get(url)
	if err != nil || code == 504 {
		fmt.Printf("请求失败，开始重试")
		for i := 0; i < 4; i++ {
			fmt.Printf("正在重试%d...", i)
			time.Sleep(time.Second * 3)
			content, code, err = Get(url)
			if err == nil {
				break
			}
		}
	}
	if err != nil {
		fmt.Println("重试超时，请检查...")
		return err
	}
	return json.Unmarshal([]byte(string(content)), resp)
}

