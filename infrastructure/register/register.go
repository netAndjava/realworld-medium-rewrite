// Package register provides ...
package register

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

func Register(consulHost, consulPort, serviceIP, servicePort, serviceName string, logger log.Logger) (registrar sd.Registrar) {
	//1. 创建Consul客户端连接Consul service
	var client consul.Client
	{
		consulCfg := api.DefaultConfig()
		consulCfg.Address = serviceIP + ":" + servicePort
		consulClient, err := api.NewClient(consulCfg)
		if err != nil {
			logger.Log("create consul client err:", err)
			return
		}
		client = consul.NewClient(consulClient)
	}
	//2. 设置Consul对服务健康检查参数
	check := api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%v:%v/%v", serviceIP, servicePort, serviceName),
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status",
	}
	//3. 设置要注册的微服务信息
	registeration := api.AgentServiceRegistration{
		ID:      serviceName + uuid.New(),
		Name:    serviceName,
		Address: serviceIP,
		Port:    servicePort,
		Check:   &check,
	}
	//4. 注册服务
	return consul.NewRegistrar(client, registeration, logger)
}
