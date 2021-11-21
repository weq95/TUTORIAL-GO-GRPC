package main

import (
	"example.com/go-usermgmt-grpc/GRPC-DEMO/helper"
	"example.com/go-usermgmt-grpc/GRPC-DEMO/services"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	rpcServer()

	httpServer()
}

func rpcServer() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))

	//商品服务
	services.RegisterOrderServiceServer(rpcServer, new(services.ProdService))

	//订单服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OrdersService))

	//用户服务
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))

	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	_ = rpcServer.Serve(listen)
}

func httpServer() {

	ctx, cancel := context.WithCancel(context.Background())

	gwmux := runtime.NewServeMux()
	defer cancel()

	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(helper.GetClientCreds()),
	}

	err := services.RegisterProdServiceHandlerFromEndpoint(
		ctx, gwmux, "localhost:8082", opt)
	if err != nil {
		log.Fatal(err)
	}

	err = services.RegisterOrderServiceHandlerFromEndpoint(
		ctx, gwmux, "localhost:8082", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	_ = httpServer.ListenAndServe()
}
