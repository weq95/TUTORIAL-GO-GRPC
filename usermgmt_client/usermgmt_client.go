package main

import (
	"context"
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var newUser = make(map[string]int32)

	newUser["Alice"] = 43
	newUser["Bob"] = 30

	for name, age := range newUser {
		r, err := c.CreateNewUser(ctx,
			&pb.NewUser{
				Name: name,
				Age:  age,
			})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`User Details: 
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
		fmt.Println("\r\n")
	}

	r, err := c.GetUsers(ctx, &pb.GetUserParams{})
	if err != nil {
		log.Fatalf("cloud not ==============> retrieve users: %v", err)
	}

	log.Printf("\rUser LIST \n")
	fmt.Printf("r.GetUsers() : %v\n", r.GetUsers())
}
