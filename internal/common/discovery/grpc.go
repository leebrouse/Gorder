package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ServiceDiscoverer 持有一个 Registry
func RegisterToConsul(ctx context.Context, serviceName string) (func() error, error) {
	//registry, err := consul.New(viper.GetString("consul.addr"))
	//if err != nil {
	//	return func() error { return nil }, err
	//}

	// Depends on the Register interface not the instances
	discoverer := NewServiceDiscoverer()

	instanceID := GenerateInstanceID(serviceName)
	grpcAddr := viper.Sub(serviceName).GetString("grpc-addr")
	if err := discoverer.reg.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		return func() error { return nil }, err
	}
	go func() {
		for {
			if err := discoverer.reg.HealthCheck(instanceID, serviceName); err != nil {
				logrus.Panicf("no heartbeat from %s to registry, err=%v", serviceName, err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	logrus.WithFields(logrus.Fields{
		"serviceName": serviceName,
		"addr":        grpcAddr,
	}).Info("registered to consul")
	return func() error {
		return discoverer.reg.Deregister(ctx, instanceID, serviceName)
	}, nil
}

func GetServiceAddr(ctx context.Context, serviceName string) (string, error) {
	//registry, err := consul.New(viper.GetString("consul.addr"))
	//if err != nil {
	//	return "", err
	//}

	discoverer := NewServiceDiscoverer()

	addrs, err := discoverer.reg.Discover(ctx, serviceName)
	if err != nil {
		return "", err
	}
	if len(addrs) == 0 {
		return "", fmt.Errorf("got empty %s addrs from consul", serviceName)
	}
	i := rand.Intn(len(addrs))
	logrus.Infof("Discovered %d instance of %s, addrs=%v", len(addrs), serviceName, addrs)
	return addrs[i], nil
}
