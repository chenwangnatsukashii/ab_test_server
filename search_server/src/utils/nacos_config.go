package utils

import (
	"encoding/json"
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"line_china/common/model"
	"log"
	"strconv"
)

var ConfigData model.Publish
var IClient config_client.IConfigClient

// ListenConfig 监听nacos配置
func ListenConfig() error {
	var thisErr error
	log.Print("开始初始化Nacos")
	clientConfig := constant.ClientConfig{
		Endpoint:    Config.Nacos.Ip + ":" + strconv.Itoa(Config.Nacos.Port) + Config.Nacos.Path,
		NamespaceId: Config.Nacos.NameSpaceId,
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		TimeoutMs:           Config.Nacos.TimeoutMs, //http请求超时时间，单位毫秒
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}

	// Nacos配置,至少一个
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      Config.Nacos.Ip,
			ContextPath: Config.Nacos.Path,
			Port:        uint64(Config.Nacos.Port),
		},
	}

	IClient, thisErr = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if thisErr != nil {
		return thisErr
	}

	// 首次拉去实验配置
	var content string
	content, thisErr = IClient.GetConfig(vo.ConfigParam{
		DataId: Config.Nacos.A,
		Group:  Config.Nacos.Group})
	if thisErr != nil {
		return thisErr
	}

	thisErr = json.Unmarshal([]byte(content), &ConfigData)
	if thisErr != nil {
		return thisErr
	}

	// 开始监听Nacos更新
	log.Print("开始监听Nacos")
	thisErr = IClient.ListenConfig(vo.ConfigParam{
		DataId: Config.Nacos.A,
		Group:  Config.Nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			log.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)

			// 对写操作加锁
			RwLock.Lock()
			defer RwLock.Unlock()
			thisErr = json.Unmarshal([]byte(data), &ConfigData)

			// nacos推送来的数据为空
			if len(data) == 0 {
				thisErr = errors.New("nacos推送数据为空")
			}
		},
	})

	if thisErr != nil {
		log.Println(thisErr)
		return thisErr
	}

	return nil
}
