package consul

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/register"
)

//Config 启动的配置项
type Config struct {
	Address string `toml:"address"`
	Check   Check  `toml:"check"`
}

type Check struct {
	Interval string `toml:"interval"`
	Timeout  string `toml:"timeout"`
	Notes    string `toml:"notes"`
}

type consulRegistrar struct {
	api.Config
	api.AgentServiceCheck
}

func NewConsulRegister(cfg api.Config, check api.AgentServiceCheck) register.Registrar {
	return &consulRegistrar{Config: cfg, AgentServiceCheck: check}
}
func (r *consulRegistrar) Register(serviceIP string, servicePort int, serviceName string, logger log.Logger) (sd.Registrar, error) {
	//1. 创建Consul客户端连接Consul service
	var client consul.Client
	{
		// consulCfg := api.DefaultConfig()
		// consulCfg.Address = serviceIP + ":" + servicePort
		consulClient, err := api.NewClient(&r.Config)
		if err != nil {
			logger.Log("create consul client err:", err)
			return nil, err
		}
		client = consul.NewClient(consulClient)
	}
	//2. 设置Consul对服务健康检查参数
	// check := api.AgentServiceCheck{
	// 	GRPC:     fmt.Sprintf("%v:%v/%v", serviceIP, servicePort, serviceName),
	// 	Interval: "10s",
	// 	Timeout:  "1s",
	// 	Notes:    "Consul check service health status",
	// }
	//3. 设置要注册的微服务信息
	registeration := api.AgentServiceRegistration{
		ID:      serviceName + uuid.New().String(),
		Name:    serviceName,
		Address: serviceIP,
		Port:    servicePort,
		Check:   &r.AgentServiceCheck,
	}
	//4. 注册服务
	return consul.NewRegistrar(client, &registeration, logger), nil
}
