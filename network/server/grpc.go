package server

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zutim/ztm-demo/common/apis"
	"github.com/zutim/ztm-demo/common/pb/order"
	"google.golang.org/grpc"
	"net"
)

type grpcServer struct {
	server *grpc.Server
}

func NewGrpc() *grpcServer {
	srv := grpc.NewServer()

	return &grpcServer{
		server: srv,
	}
}

// register server to grpc
func (g *grpcServer) Register() *grpcServer {

	order.RegisterOrderServer(g.server, apis.Order())

	return g
}

func (g *grpcServer) Run() {

	addr := viper.GetString("rpc")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("err")
	}

	fmt.Println("Listening and serving Rpc on ", addr)

	if err := g.server.Serve(lis); err != nil {
		panic(err)
	}

}
