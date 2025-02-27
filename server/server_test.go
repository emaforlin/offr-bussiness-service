package grpc_server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildGrpcServer(t *testing.T) {
	builder := &GrpcServerBuilder{}
	builder.DisableDefaultHealthCheck(true)
	builder.EnableReflection(true)
	builder.SetUnaryInterceptors(nil)

	server := builder.Build()
	assert.NotNil(t, server)
}
