package storage

import (
	"context"
	pb "auth/genproto/register"
)

type IStorage interface {
	User() IUserStorage
	Close()
}

type IUserStorage interface {
	CreateRegister(context.Context, *pb.CreateRegisterRequest) (*pb.CreateRegisterResponse, error)
	Update(context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error)
	AddImage(context.Context, *pb.AddImageRequest) (*pb.AddImageResponse, error)
	GetRegister(context.Context, *pb.GetRegisterRequest) (*pb.GetRegisterResponse, error)
	GetRegisters(context.Context, *pb.GetRegistersRequest) (*pb.GetRegistersResponse, error)
	DeleteRegister(context.Context, *pb.DeleteRegisterRequest) (*pb.DeleteRegisterResponse, error)
	Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)	
	UpdatePassword(context.Context, *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error)
	CreatePassword(context.Context, *pb.CreatePasswordRequest) (*pb.CreatePasswordResponse, error)
	ConfirmationPassword(context.Context, *pb.ConfirmationPasswordRequest) (*pb.ConfirmationPasswordResponse, error)
}
