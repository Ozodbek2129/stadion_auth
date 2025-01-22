package main

import (
	"auth/api"
	"auth/api/handler"
	"auth/config"
	pb "auth/genproto/register"
	"auth/pkg/logger"
	"auth/service"
	"auth/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().USER_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}

	logs := logger.NewLogger()

	service := service.NewUserService(db, logs)

	server := grpc.NewServer()
	pb.RegisterRegisterServiceServer(server, service)

	log.Printf("Server listening at %v", listener.Addr())
	go func() {
		err := server.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
	}()

	hand := handler.NewHandler()
	router := api.NewRouter(hand)
	err = router.Run(config.Load().AUTH_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
}
