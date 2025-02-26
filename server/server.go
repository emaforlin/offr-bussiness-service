package grpc_server

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type GrpcServer interface {
	Start(addr string) error
	RegisterService(reg func(*grpc.Server))
	AwaitTermination(shutdownHook func())
}

type GrpcServerBuilder struct {
	options                   []grpc.ServerOption
	enabledReflection         bool
	shutdownHook              func()
	enabledHealthCheck        bool
	disableDefaultHealthCheck bool
}

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

// RegisterService implements GrpcServer.
func (s *grpcServer) RegisterService(reg func(*grpc.Server)) {
	reg(s.server)
}

// Start the GrpcServer
func (s *grpcServer) Start(addr string) error {
	var err error

	s.listener, err = net.Listen("tcp", addr)

	if err != nil {
		msg := fmt.Sprintf("Failed to listen: %v", err)
		return errors.New(msg)
	}

	go s.serv()

	log.Infof("gRPC Server started on %s", addr)
	return nil
}

func (s *grpcServer) AwaitTermination(shutdownHook func()) {
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGTERM)
	<-interruptSignal

	s.cleanup()
	if shutdownHook != nil {
		shutdownHook()
	}
}

func (s *grpcServer) cleanup() {
	log.Info("Stopping the server...")
	s.server.GracefulStop()
	log.Info("Closing listner.")
	s.listener.Close()
	log.Info("Server stopped.")
}

func (s *grpcServer) serv() {
	if err := s.server.Serve(s.listener); err != nil {
		log.Errorf("failed to serve: %v", err)
	}
}

// Builds a new Server
func (sb *GrpcServerBuilder) Build() GrpcServer {
	srv := grpc.NewServer(sb.options...)

	if !sb.disableDefaultHealthCheck {
		grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
	}

	if sb.enabledReflection {
		reflection.Register(srv)
	}

	return &grpcServer{server: srv, listener: nil}
}

// SetUnaryInterceptors set a list of interceptors to the gRPC server for unary connections
func (sb *GrpcServerBuilder) SetUnaryInterceptors(interceptors []grpc.UnaryServerInterceptor) {
	chain := grpc.ChainUnaryInterceptor(interceptors...)
	sb.AddOption(chain)
}

// AddOption configures new options to set up the connection
func (sb *GrpcServerBuilder) AddOption(o grpc.ServerOption) {
	sb.options = append(sb.options, o)
}

// EnableReflection enables the reflection
// This provides information about publicly-accessible gRPC services on a server,
// and helps clients at runtime to build RPC requests and responses without precompiled service information.
// It is used by tools like grpc-curl and postman to instrospec server and help with debugging.
// WARNING! this shouldn't be enabled in production
func (sb *GrpcServerBuilder) EnableReflection(e bool) {
	sb.enabledReflection = e
}

// DisableDefaultHealthCheck disables the default health check service
// WARNING! if you disable the default health check service, you must provide a custom one
func (sb *GrpcServerBuilder) DisableDefaultHealthCheck(e bool) {
	sb.disableDefaultHealthCheck = e
}
