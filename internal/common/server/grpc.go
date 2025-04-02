package server

import (
	"net"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// RunGRPCServer 根据服务名称启动 gRPC 服务器，并调用传入的注册函数来注册具体的服务
func RunGRPCServer(serviceName string, registerServer func(server *grpc.Server)) {
	// 从配置中获取指定服务的 gRPC 地址
	addr := viper.Sub(serviceName).GetString("grpc-addr")
	if addr == "" {
		// 如果配置中没有找到服务地址，则使用 fallback 地址
		// TODO: 添加警告日志，提示未配置服务地址
		addr = viper.GetString("fallback-grpc-server")
	}
	// 根据获取的地址启动 gRPC 服务器
	RunGRPCServerOnAddr(addr, registerServer)
}

// RunGRPCServerOnAddr 根据指定的地址启动 gRPC 服务器，并注册服务
func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	// 创建一个 logrus 日志条目，用于日志记录
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	// 创建 gRPC 服务器，并添加一系列拦截器
	grpcServer := grpc.NewServer(
		// 配置 Unary 拦截器，用于处理单次请求
		grpc.ChainUnaryInterceptor(
			// 使用 grpc_tags 拦截器为请求提取字段信息
			grpc_tags.UnaryServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			// 使用 grpc_logrus 拦截器，将请求日志写入 logrus
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			// 以下拦截器被注释掉，可根据需要启用：
			// otelgrpc.UnaryServerInterceptor(),
			// srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			// logging.UnaryServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			// selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			// recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		// 配置 Stream 拦截器，用于处理流式请求
		grpc.ChainStreamInterceptor(
			// 使用 grpc_tags 拦截器为流请求提取字段信息
			grpc_tags.StreamServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			// 使用 grpc_logrus 拦截器，将流请求日志写入 logrus
			grpc_logrus.StreamServerInterceptor(logrusEntry),
			// 以下拦截器被注释掉，可根据需要启用：
			// otelgrpc.StreamServerInterceptor(),
			// srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			// logging.StreamServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			// selector.StreamServerInterceptor(auth.StreamServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			// recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)
	// 调用传入的注册函数，将具体的服务注册到 gRPC 服务器中
	registerServer(grpcServer)

	// 在指定地址上监听 TCP 连接
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		// 如果监听过程中发生错误，则记录 Panic 日志并终止程序
		logrus.Panic(err)
	}
	// 记录服务器启动日志，显示监听的地址
	logrus.Infof("Starting gRPC server, Listening: %s", addr)
	// 启动 gRPC 服务器开始服务
	if err := grpcServer.Serve(listen); err != nil {
		// 如果服务过程中出现错误，则记录 Panic 日志并终止程序
		logrus.Panic(err)
	}
}
