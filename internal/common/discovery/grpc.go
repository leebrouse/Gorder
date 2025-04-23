package discovery

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/leebrouse/Gorder/common/discovery/consul"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

// RegisterToConsul 将服务注册到 Consul，并返回注销函数
// 参数：ctx - 上下文控制、serviceName - 当前服务的名称（用于注册标识）
// 返回：一个注销函数（用于在服务退出时调用），以及注册过程中可能出现的错误
func RegisterToConsul(ctx context.Context, serviceName string) (func() error, error) {
	// 创建 Consul 注册中心实例，读取 consul 地址
	registry, err := consul.New(viper.GetString("consul.addr"))
	if err != nil {
		// 返回空注销函数和错误
		return func() error {
			return nil
		}, err
	}

	// 生成当前实例的唯一 ID（常用于健康检查标识）
	instanceID := GenerateInstanceID(serviceName)

	// 从配置文件中获取当前服务监听的 gRPC 地址（如 "localhost:8080"）
	grpcAddr := viper.Sub(serviceName).GetString("grpc-addr")

	// 向 Consul 注册服务
	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		// 注册失败，返回空注销函数和错误
		return func() error {
			return nil
		}, err
	}

	// 启动 Goroutine 进行 TTL 健康检查的周期性上报
	go func() {
		for {
			// 每秒上报一次状态为“通过”，如果失败就 panic（强制退出）
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Panicf("no heartbeat from %s to registry, err=%v", serviceName, err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// 打印成功注册日志
	logrus.WithFields(logrus.Fields{
		"serviceName": serviceName,
		"addr":        grpcAddr,
	}).Info("registered to consul")

	// 返回注销函数（可以在 main() 中 defer 调用），以及 nil 错误
	return func() error {
		return registry.Deregister(ctx, instanceID, serviceName)
	}, nil
}

func GetServiceAddr(ctx context.Context, serviceName string) (string, error) {
	registry, err := consul.New(viper.GetString("consul.addr"))
	if err != nil {
		return "", nil
	}
	addrs, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return "", err
	}
	if len(addrs) == 0 {
		return "", fmt.Errorf("got empty addrs from consul")
	}
	i := rand.Intn(len(addrs))
	logrus.Infof("Discovered %d instance of %s,addrs=%v", len(addrs), serviceName, addrs)
	return addrs[i], nil
}
