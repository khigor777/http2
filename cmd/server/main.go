package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/khigor777/http2/proto"
	"golang.org/x/net/context"

)
//protoc --go_out=plugins=grpc:. proto/users.proto
func main()  {
	l, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalf("can' listsen server %v", err)
	}
	srv := grpc.NewServer()
	proto.RegisterAccountServer(srv, nil)
	log.Fatal(srv.Serve(l))
}

type Server struct {
}
func AddUser(ctx context.Context, addUser *proto.AddUserRequest) (*proto.AddUserResponse, error){
	return nil, nil
}
func GetUsers(ctx context.Context, getUser*proto.GetUsersRequest) (*proto.GetUsersResponse, error){
	return nil, nil
}