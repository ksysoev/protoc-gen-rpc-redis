package gen

import (
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
)

func TestGenerateMethod(t *testing.T) {
	plugin := &protogen.Plugin{}
	g := plugin.NewGeneratedFile("test.go", "test")
	method := &protogen.Method{
		Parent: &protogen.Service{
			GoName: "TestService",
		},
		GoName: "TestMethod",
		Input: &protogen.Message{
			GoIdent: protogen.GoIdent{
				GoName:       "TestInput",
				GoImportPath: "github.com/ksysoev/protoc-gen-rpc-redis/pkg/gen",
			},
		},
		Output: &protogen.Message{
			GoIdent: protogen.GoIdent{
				GoName:       "TestOutput",
				GoImportPath: "github.com/ksysoev/protoc-gen-rpc-redis/pkg/gen",
			},
		},
	}

	expectedMethod := `
func (x *RPCRedisTestService) handleTestMethod(req *rpc_redis.Request) (any, error) {
	var rpcReq gen.TestInput

	err := req.ParseParams(&rpcReq)
	if err != nil {
		return nil, fmt.Errorf("error parsing request: %v", err)
	}

	return x.service.TestMethod(req.Context(), &rpcReq)
}
`

	renderedMethod, err := generateMethod(g, method)
	if err != nil {
		t.Fatalf("failed to generate method: %v", err)
	}

	if renderedMethod != expectedMethod {
		t.Fatalf("invalid method code, expected:\n%s\ngot:\n%s", expectedMethod, renderedMethod)
	}
}

func TestGenerateService(t *testing.T) {
	plugin := &protogen.Plugin{}
	g := plugin.NewGeneratedFile("test.go", "test")

	service := &protogen.Service{
		GoName: "TestService",
		Methods: []*protogen.Method{
			{GoName: "TestMethod1"},
			{GoName: "TestMethod2"},
		},
	}

	expectedService := `
// TestService is the server API for example.TestService
type RPCRedisTestService struct {
	rpcSever *rpc_redis.Server
	service  *TestServiceService
}

func NewRedisTestService(redis *v9.Client, grpcService *TestServiceService, opts ...rpc_redis.ServerOption) *RPCRedisTestService {
	rpcServer := rpc_redis.NewServer(redis, "example.TestService", "TestServiceGroup", uuid.New().String(), opts...)
	service := &RPCRedisTestService{
		rpcSever: rpcServer,
		service:  grpcService,
	}

	// Register handlers
	rpcServer.AddHandler("TestMethod1", service.handleTestMethod1)
	rpcServer.AddHandler("TestMethod2", service.handleTestMethod2)

	return service
}

func (x *RPCRedisTestService) Serve() error {
	return x.rpcSever.Run()
}

func (x *RPCRedisTestService) Close() {
	x.rpcSever.Close()
}
`

	renderedService, err := generateService(g, service, "example.TestService")
	if err != nil {
		t.Fatalf("failed to generate method: %v", err)
	}

	if renderedService != expectedService {
		t.Fatalf("invalid method code, expected:\n%s\ngot:\n%s", expectedService, renderedService)
	}
}

func TestGenerateFileHeader(t *testing.T) {
	file := &protogen.File{
		GoPackageName: protogen.GoPackageName("example"),
	}

	expectedHeader := `
// Code generated by protoc-gen-rpc-redis. DO NOT EDIT.

package example
`

	renderedHeader, err := generateFileHeader(file)
	if err != nil {
		t.Fatalf("failed to generate file header: %v", err)
	}

	if renderedHeader != expectedHeader {
		t.Fatalf("invalid file header, expected:\n%s\ngot:\n%s", expectedHeader, renderedHeader)
	}
}
