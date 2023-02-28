package user_config

import (
	"encoding/json"
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	_ "github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"py_gomall/v2/common/free_port"
	"py_gomall/v2/user_web/user_dial"
	"py_gomall/v2/user_web/user_run"
	"strconv"
)

func ParseNacosConf() {
	config := new(NacosConf)
	viper.SetConfigName("user_web_nacos")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./user_config")
	viper.AddConfigPath("./user_web/user_config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal(err)
	}
	if config.NacosAddr == "" {
		err := errors.New("read nacos config file error,please check the filepath")
		log.Fatal(err)
	}
	log.Printf("read nacos config file ok! viper serve and listening...")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			log.Printf("upgrade nacos config file:%v,error:%v\n", e.Name, err)
		} else {
			if err = viper.ReadInConfig(); err == nil {
				if err = viper.Unmarshal(config); err == nil {
					log.Printf("upgrage nacos config success , update project config now...")
					AppConf.NacosConf = config
					nacosAt(config)
				} else {
					log.Printf("upgrage nacos config unmarshal error:%v\n", err)
				}
			} else {
				log.Printf("upgrage nacos config readed error:%v\n", err)
			}
		}
	})
	nacosAt(config)
	return
}

func nacosAt(nc *NacosConf) {
	nacosServerConfig := []constant.ServerConfig{
		{IpAddr: nc.NacosAddr, Port: nc.NacosPort},
	}
	nacosClientConfig := constant.ClientConfig{
		NamespaceId:         nc.NamespaceId,
		TimeoutMs:           nc.TimeoutMs,
		NotLoadCacheAtStart: nc.LoadCache,
		LogDir:              nc.LogDir,
		CacheDir:            nc.CacheDir,
		LogLevel:            nc.LogLevel,
	}
	nacosClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": nacosServerConfig,
		"clientConfig":  nacosClientConfig,
	})
	if err != nil {
		log.Fatal(err)
	}
	content, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: nc.DataId,
		Group:  nc.Group,
	})
	conf := &Conf{}
	if err = json.Unmarshal([]byte(content), conf); err != nil {
		log.Fatal(err)
	}
	AppConf = conf
	addr()
	log.Printf("parse project config success! \n nacos serve on listening...")
	err = nacosClient.ListenConfig(vo.ConfigParam{
		DataId: nc.DataId,
		Group:  nc.Group,
		OnChange: func(namespace, group, dataId, data string) {
			if err = json.Unmarshal([]byte(data), conf); err != nil {
				log.Printf("ERROR:: Failed to update project config file, please restart the service\n error:%v\n", err)
			} else {
				AppConf = conf
				addr()
				log.Printf("update project config file ok:namespace:%s,group:%s,dataId:%s\ndata:%s\n", namespace, group, dataId, data)
				user_run.ReStart()
			}
		},
	})
	if err != nil {
		log.Printf("ERROR:: Failed to update project config file, please restart the service\n error:%v\n", err)
	}
}

func addr() {
	if AppConf.Listen == "" {
		zap.L().Info("No configuration or unable to read the listening address, start the default address: 0.0.0.0")
		AppConf.Listen = "0.0.0.0"
	}
	if AppConf.Port == 0 {
		if port, err := free_port.FreePort(); err != nil {
			zap.L().Info("No configuration or unable to read the listening port, start the default port: 8000")
			AppConf.Port = 8000
		} else {
			AppConf.Port = port
			zap.L().Info("No configuration or unable to read the listening port, start the free port:", zap.Int("port", port))
		}
	}
	user_run.AppHost = AppConf.Listen + ":" + strconv.Itoa(AppConf.Port)
	newDialConf()
}

func newDialConf() {
	user_dial.Services = AppConf.Services
	user_dial.Host = AppConf.Listen
	user_dial.Port = AppConf.Port
	user_dial.Name = AppConf.Name
	user_dial.ConsulHost = AppConf.ConsulHost
	user_dial.ConsulPort = AppConf.ConsulPort
	user_dial.TimeoutSec = AppConf.TimeoutSec
	user_dial.IntervalSec = AppConf.IntervalSec
	user_dial.RemoveAfterSec = AppConf.RemoveAfterSec
	user_dial.Tages = AppConf.Tage
}
