package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strconv"
	"strings"
	"sync"
)

// Registry 封装了 Consul 客户端，用于服务注册与发现
type Registry struct {
	client *api.Client // Consul 客户端
}

// 单例变量与初始化控制
var (
	consulClient *Registry // 仅初始化一次的 Registry 实例
	once         sync.Once // 控制只执行一次
	initErr      error     // 初始化过程中发生的错误
)

// New 初始化 Consul 客户端，只会执行一次，返回 Registry 单例
func New(consulAddr string) (*Registry, error) {
	once.Do(func() {
		// 配置 Consul 地址
		config := api.DefaultConfig()
		config.Address = consulAddr

		// 创建 Consul 客户端
		client, err := api.NewClient(config)
		if err != nil {
			initErr = err
			return
		}

		// 创建单例实例
		consulClient = &Registry{
			client: client,
		}
	})

	// 如果初始化过程中有错误，返回错误
	if initErr != nil {
		return nil, initErr
	}
	return consulClient, nil
}

// Register 向 Consul 注册服务
func (r *Registry) Register(_ context.Context, instanceID, serviceName, hostPort string) error {
	// 拆分 "host:port" 格式
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("invalid host:port format [Example: localhost:8080 ]")
	}
	host := parts[0]
	post, _ := strconv.Atoi(parts[1]) // 将端口转换为 int

	// 注册服务
	return r.client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      instanceID,  // 实例唯一 ID
		Name:    serviceName, // 服务名称
		Address: host,        // 服务地址
		Port:    post,        // 服务端口
		Check: &api.AgentServiceCheck{
			CheckID:                        instanceID, // 健康检查 ID
			TTL:                            "5s",       // 每 5s 要上报一次状态
			Timeout:                        "5s",       // 检查超时时间
			DeregisterCriticalServiceAfter: "10s",      // 10s 内未上报将移除服务
		},
	})
}

// Deregister 注销服务（取消注册健康检查）
func (r *Registry) Deregister(ctx context.Context, instanceID, serviceName string) error {
	logrus.WithFields(logrus.Fields{
		"instanceID":  instanceID,
		"serviceName": serviceName,
	}).Info("deregister from consul")

	// 注销健康检查（间接注销服务）
	return r.client.Agent().CheckDeregister(instanceID)
}

// Discover 通过服务名获取所有可用实例（带健康检查通过）
func (r *Registry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, e := range entries {
		// 拼接 IP:PORT
		ips = append(ips, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port))
	}
	return ips, nil
}

// HealthCheck 上报 TTL 健康检查结果为通过（"online"）
func (r *Registry) HealthCheck(instanceID, serviceName string) error {
	// 将健康状态更新为通过
	return r.client.Agent().UpdateTTL(instanceID, "online", api.HealthPassing)
}
