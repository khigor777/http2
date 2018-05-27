package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/khigor777/http2/proto"
	"golang.org/x/net/context"

	"github.com/khigor777/http2/users"
)
//protoc --go_out=plugins=grpc:. proto/users.proto
func main()  {
	l, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalf("can' listsen server %v", err)
	}
	srv := grpc.NewServer()
	proto.RegisterAccountServer(srv, new(Server))
	log.Fatal(srv.Serve(l))
}

type Server struct {
	users *users.Users
}
func (s *Server)AddUser(ctx context.Context, addUser *proto.AddUserRequest) (*proto.AddUserResponse, error){
	s.users.Add(&users.Account{
		Name:addUser.Name,
		Email:addUser.Email,
	})
	resp := new(proto.AddUserResponse)
	resp.Result = "ok"
	return resp, nil
}
func (s *Server)GetUsers(ctx context.Context, getUser *proto.GetUsersRequest) (*proto.GetUsersResponse, error){
	u := s.users.GetAll()
	var res *proto.GetUsersResponse
	for _, v := range u {
		res.Users = append(res.Users, &proto.User{
			Name:v.Name,
			Email:v.Email,
		})
	}
	return res, nil
}