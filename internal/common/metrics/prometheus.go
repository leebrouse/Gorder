package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

// PrometheusMetricsClient 封装了 Prometheus 的 registry 注册器
type PrometheusMetricsClient struct {
	registry *prometheus.Registry
}

// dynamicCounter 是一个带标签的计数器，用于记录自定义 key 的次数
var dynamicCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "dynamic_counter",   // 指标名称，在 Prometheus 查询中使用
		Help: "Count custom keys", // 指标的帮助信息
	},
	[]string{"key"}, // 使用 "key" 作为标签，区分不同维度
)

// PrometheusMetricsClientConfig 配置结构体，包含服务监听地址和服务名称
type PrometheusMetricsClientConfig struct {
	Host        string // 监听地址，例如 ":8080"
	ServiceName string // 服务名称，用于指标标签
}

// NewPrometheusMetricsClient 创建一个新的 PrometheusMetricsClient 实例
func NewPrometheusMetricsClient(config *PrometheusMetricsClientConfig) *PrometheusMetricsClient {
	client := &PrometheusMetricsClient{}
	client.initPrometheus(config) // 初始化 Prometheus
	return client
}

// 初始化 Prometheus，包括注册默认采集器、自定义采集器，以及启动 HTTP 服务
func (p *PrometheusMetricsClient) initPrometheus(config *PrometheusMetricsClientConfig) {
	// 创建新的指标注册器
	p.registry = prometheus.NewRegistry()

	// 注册默认采集器：Go运行时指标、进程信息指标
	p.registry.MustRegister(
		collectors.NewGoCollector(),                                       // Go 运行时相关指标
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}), // 当前进程相关指标
	)

	// 注册自定义的 dynamicCounter
	p.registry.MustRegister(dynamicCounter)

	// 给注册器包装服务名称标签（注意：WrapRegistererWith 返回的是 Registerer，并不会影响已有 registry）
	prometheus.WrapRegistererWith(prometheus.Labels{"serviceName": config.ServiceName}, p.registry)

	// 启动 HTTP Server 暴露 /metrics 接口，供 Prometheus 拉取
	http.Handle("/metrics", promhttp.HandlerFor(p.registry, promhttp.HandlerOpts{}))

	// 启动 HTTP 服务（ListenAndServe 是阻塞的，因此放在 goroutine 中）
	go func() {
		logrus.Fatalf("failed to start prometheus metrics endpoint, err=%v", http.ListenAndServe(config.Host, nil))
	}()
}

// Inc 自增指标的值
func (p PrometheusMetricsClient) Inc(key string, value int) {
	dynamicCounter.WithLabelValues(key).Add(float64(value))
}
