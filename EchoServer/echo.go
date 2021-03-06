package EchoServer

import (
	context "context"
	"fmt"
	pb "golang_grpc_gin_jaeger_A/hello"
	"io"

	// "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
)

type EchoServer struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

func (e *EchoServer) SayHello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloReply, err error) {

	fmt.Println("[Server receive client request]" + req.GetMessage())

	span := e.Tracer.StartSpan("SayHello")
	span.SetTag("SayHello", "pong")
	defer span.Finish()

	{
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		example1(ctx)
	}

	return &pb.HelloReply{
		Message: "[Echo From Server] " + req.GetMessage(),
	}, nil

}
