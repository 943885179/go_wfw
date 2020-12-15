package mzjnacos

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

type NacOsServerConfig struct {
	Scheme      string `json:"scheme"`       //the nacos server scheme
	ContextPath string `json:"context_path"` //the nacos server contextpath
	IpAddr      string `json:"ip_addr"`      //the nacos server address
	Port        uint64 `json:"port"`         //the nacos server port
}

//GetConfigJsonToEntity 配置文件是json的话直接读取成实体
func (n *NacOsServerConfig) GetConfigJsonToEntity(dataId, group string, resp interface{}) error {
	c, err := n.GetConfig(dataId, group)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(c), resp)
}

// GetConfig 读取string
func (n *NacOsServerConfig) GetConfig(dataId, group string) (string, error) {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      n.IpAddr,
			ContextPath: n.ContextPath,
			Port:        n.Port,
			Scheme:      n.Scheme,
		},
	}

	c, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		log.Fatal(err)
	}
	content, err := c.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	return content, nil
}

//ConfigByNacOS 通过nacos做配置中心获取config
//github.com/nacos-group/nacos-sdk-go
func main() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}

	c, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		log.Fatal(err)
	}
	content, err := c.GetConfig(vo.ConfigParam{
		DataId: "test1",
		Group:  "DEFAULT_GROUP",
	})
	fmt.Println(content, err)
	// 封装后调用
	t := NacOsServerConfig{
		IpAddr:      "127.0.0.1",
		ContextPath: "/nacos",
		Port:        8848,
	}
	fmt.Println(t.GetConfig("test1", "DEFAULT_GROUP"))
}
