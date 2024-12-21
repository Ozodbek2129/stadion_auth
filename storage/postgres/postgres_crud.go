package postgres

import (
	"auth/pkg/logger"
	"auth/storage"
	"context"
	"database/sql"
	"log/slog"
	pb "auth/genproto/register"
)

type UserRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewUserRepository(db *sql.DB) storage.IUserStorage {
	return &UserRepository{Db: db, Log: logger.NewLogger()}
}

func (u *UserRepository) CreateRegister(context.Context, *pb.CreateRegisterRequest) (*pb.CreateRegisterResponse, error) {
	return nil, nil
}

func (u *UserRepository) Update(context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return nil, nil
}

func (u *UserRepository) AddImage(context.Context, *pb.AddImageRequest) (*pb.AddImageResponse, error) {
	return nil, nil
}

func (u *UserRepository) GetRegister(context.Context, *pb.GetRegisterRequest) (*pb.GetRegisterResponse, error) {
	return nil, nil
}

func (u *UserRepository) GetRegisters(context.Context, *pb.GetRegistersRequest) (*pb.GetRegistersResponse, error) {
	return nil, nil
}

func (u *UserRepository) DeleteRegister(context.Context, *pb.DeleteRegisterRequest) (*pb.DeleteRegisterResponse, error) {
	return nil, nil
}

func (u *UserRepository) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (u *UserRepository) UpdatePassword(context.Context, *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	return nil, nil
}

func (u *UserRepository) CreatePassword(context.Context, *pb.CreatePasswordRequest) (*pb.CreatePasswordResponse, error) {
	return nil, nil
}

func (u *UserRepository) ConfirmationPassword(context.Context, *pb.ConfirmationPasswordRequest) (*pb.ConfirmationPasswordResponse, error) {
	return nil, nil
}
