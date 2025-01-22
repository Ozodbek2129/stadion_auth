package handler

import (
	"auth/config"
	pb "auth/genproto/register"
	"auth/pkg/logger"
	"log"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	AuthUser pb.RegisterServiceClient
	Log      *slog.Logger
}

func NewHandler() *Handler {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error while connecting to authentication service: %v", err)
	}

	return &Handler{
		AuthUser: pb.NewRegisterServiceClient(conn),
		Log:      logger.NewLogger(),
	}
}
