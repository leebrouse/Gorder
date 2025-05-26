package discovery

import (
	"context"
	"fmt"
	"github.com/leebrouse/Gorder/common/discovery/consul"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

type Registry interface {
	Register(ctx context.Context, instanceID, serviceName, hostPort string) error
	Deregister(ctx context.Context, instanceID, serviceName string) error
	Discover(ctx context.Context, serviceName string) ([]string, error)
	HealthCheck(instanceID, serviceName string) error
}

type ServiceDiscoverer struct {
	reg Registry
}

// impelement part inject
func NewServiceDiscoverer() *ServiceDiscoverer {
	registry, err := consul.New(viper.GetString("consul.addr"))
	if err != nil {
		return nil
	}
	return &ServiceDiscoverer{reg: registry}
}

func GenerateInstanceID(serviceName string) string {
	x := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	return fmt.Sprintf("%s-%d", serviceName, x)
}
