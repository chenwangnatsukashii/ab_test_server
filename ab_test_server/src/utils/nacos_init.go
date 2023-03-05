package utils

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"strconv"
)

var IClient config_client.IConfigClient

// InitNacos 发布nacos配置
func InitNacos() {
	var err error
	log.Print("开始初始化nacos")
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

	IClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		log.Println(err)
		return
	}
	log.Print("成功初始化nacos")
}

// PublishConfig 发布配置到Nacos上
func PublishConfig(msg string) error {
	var success bool
	var err error
	success, err = IClient.PublishConfig(vo.ConfigParam{
		DataId:  Config.Nacos.A,
		Group:   Config.Nacos.Group,
		Content: msg})
	if err != nil {
		return err
	}

	if success {
		log.Println("发布成功")
	} else {
		log.Println("发布失败")
	}
	return nil
}
