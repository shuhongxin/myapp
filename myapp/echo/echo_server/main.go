package echo_server

import (
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	pb "myapp/echo/echo"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement EchoServer
type server struct{}

//Send implements Echo.EchoServer
func (s *server) Send(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReplay, error) {
	glog.Info(in.Message)
	return &pb.EchoReplay{Message: in.Message}, nil
}

func (s *server) Push(in *pb.Autoreply ,stream pb.Echo_PushServer) error {
		ticker := time.NewTicker(time.Minute * 1)
		for _ = range ticker.C {
			tm := time.Now()
            stream.Send(&pb.Autoreply{Message: tm.Format("2006-01-02 03:04:05 PM")})
		}
		return nil
}

func Start() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s,&server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
