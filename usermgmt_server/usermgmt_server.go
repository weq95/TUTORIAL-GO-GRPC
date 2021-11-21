package main

import (
	"context"
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
)

const port = ":50051"

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	//userList *pb.UserList
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		/*userList: &pb.UserList{
			Users: make([]*pb.User, 0),
		},*/
	}
}

func (s *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserManagementServer(server, s)
	log.Printf("server listening at :%v", lis.Addr())

	return server.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, user *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", user.GetName())

	var (
		readBytes, err = ioutil.ReadFile("user.json")
		userList       = &pb.UserList{}
		userId         = int32(rand.Intn(1000))
		createUser     = &pb.User{
			Name: user.GetName(),
			Age:  user.GetAge(),
			Id:   userId,
		}
	)

	//s.userList.Users = append(s.userList.Users, createUser)

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("File not found. Creating a new file")

			userList.Users = append(userList.Users, createUser)
			jsonBytes, err := protojson.Marshal(userList)
			if err != nil {
				log.Fatalf("json marshaling filed: %v", err)
			}

			if err := ioutil.WriteFile("user.json", jsonBytes, 0664); err != nil {
				log.Fatalf("failed write to file: %v", err)
			}

			return createUser, nil
		}

		log.Fatalln("error reading file: ", err)
	}

	if err := protojson.Unmarshal(readBytes, userList); err != nil {
		log.Fatalf("failed to parse user list: %v", err)

	}

	userList.Users = append(userList.Users, createUser)
	jsonBytes, err := protojson.Marshal(userList)
	if err != nil {
		log.Fatalf("json marshaling filed: %v", err)
	}

	if err := ioutil.WriteFile("user.json", jsonBytes, 0664); err != nil {
		log.Fatalf("failed write to file: %v", err)
	}

	return createUser, nil
}

/*func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUserParams) (*pb.UserList, error) {
	return s.userList, nil
}*/

func (s *UserManagementServer) GetUsers(ctx context.Context, params *pb.GetUserParams) (*pb.UserList, error) {
	jsonBytes, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Fatalf("failed read from file: %v", err)
	}

	var userList = &pb.UserList{}

	if err := protojson.Unmarshal(jsonBytes, userList); err != nil {
		log.Fatalf("unmarshaling failed: %v", err)
	}

	return userList, nil
}

func main() {
	var userMgmtServer = NewUserManagementServer()
	if err := userMgmtServer.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
