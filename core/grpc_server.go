package core

import (
	"context"
	"errors"
	"eva/biz/handler/private"
	"eva/global"
	proto "eva/idl/proto/private"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunGrpcServer() {
	// 监听端口
	s := grpc.NewServer()
	proto.RegisterAutoLabelServer(s, &private.PrivateRpc{})
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	// 注册健康检察服务， 支持consul来进行健康检查
	//grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	// 启动gRPC服务
	errCh := make(chan error)
	go func() {
		errCh <- s.Serve(lis)
	}()
	signalWaiter := waitSignal

	if err = signalWaiter(errCh); err != nil {
		global.EVA_LOG.Error("GRPC: Receive close signal: error=%v", zap.Error(err))
		if err = lis.Close(); err != nil {
			global.EVA_LOG.Error("GRPC: Close error=%v", zap.Error(err))
		}
		return
	}
	global.EVA_LOG.Error("微服务正在优雅退出,  wait at most num=5 seconds...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go ShutDone(ctx)
	s.GracefulStop()
}

func ShutDone(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			global.EVA_LOG.Error("优雅退出", zap.Error(ctx.Err()))
			return
		default:
			{
			}
		}
	}
}

func waitSignal(errCh chan error) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	select {
	case sig := <-signals:
		switch sig {
		case syscall.SIGTERM:
			// force exit
			return errors.New(sig.String()) // nolint
		case syscall.SIGHUP, syscall.SIGINT:
			// graceful shutdown
			return nil
		}
	case err := <-errCh:
		// error occurs, exit immediately
		return err
	}

	return nil
}

func GrpcToHttp() {
	conn, err := grpc.DialContext(
		context.Background(),
		":9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		global.EVA_LOG.Error("Failed to dial server:", zap.Error(err))
	}

	gwmux := runtime.NewServeMux()
	err = proto.RegisterAutoLabelHandler(context.Background(), gwmux, conn)
	if err != nil {
		global.EVA_LOG.Error("Failed to register gateway:", zap.Error(err))
	}
	gwServer := &http.Server{
		Addr:    ":8889",
		Handler: gwmux,
	}

	global.EVA_LOG.Info("Serving gRPC-Gateway on http://0.0.0.0:8889")
	err = gwServer.ListenAndServe()
	if err != nil {
		fmt.Printf("gwServer.ListenAndServe failed, err:%v", err)
		return
	}
}
