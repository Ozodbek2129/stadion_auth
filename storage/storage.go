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
	UpdatePassword(context.Context, *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error)
	GetByEmail (context.Context, *pb.GetByEmailRequest) (*pb.GetByEmailResponse, error)
	UpdateRole (context.Context, *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error)
	Tobeanadmin (context.Context, *pb.TobeanadminRequest) (*pb.TobeanadminResponse, error)
	CheckUserId (context.Context, *pb.CheckUserIdRequest) (*pb.CheckUserIdResponse, error)
}
