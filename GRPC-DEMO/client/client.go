package main

import (
	"context"
	"example.com/go-usermgmt-grpc/GRPC-DEMO/helper"
	"example.com/go-usermgmt-grpc/GRPC-DEMO/services"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(helper.GetClientCreds()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userClient := services.NewUserServiceClient(conn)

	var uid int32
	stream, err := userClient.GetUserScoreByTWS(context.Background())

	for j := 0; j <= 3; j++ {
		req := services.UserScoreRequest{}
		req.Users = make([]*services.UserInfo, 0)

		for i := 1; i <= 5; i++ {
			req.Users = append(req.Users, &services.UserInfo{
				UserId: uid,
			})

			uid++

			err := stream.Send(&req)
			if err != nil {
				log.Println(err)
			}

			res, err := stream.Recv()
			if err == io.EOF {
				log.Println(err)
			}

			fmt.Println(res.Users)
		}
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("close succeed !")
}
