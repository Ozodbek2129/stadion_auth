package service

import (
	pb "auth/genproto/register"
	"auth/storage"
	"auth/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type UserService struct {
	pb.UnimplementedRegisterServiceServer
	User storage.IUserStorage
	Log  *slog.Logger
}

func NewUserService(db *sql.DB, log *slog.Logger) *UserService {
	return &UserService{
		User: postgres.NewUserRepository(db),
		Log:  log,
	}
}

func (u *UserService) CreateRegister(ctx context.Context, req *pb.CreateRegisterRequest) (*pb.CreateRegisterResponse, error) {
	res, err := u.User.CreateRegister(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error creating register service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	res, err := u.User.Update(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error updating register service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) AddImage(ctx context.Context, req *pb.AddImageRequest) (*pb.AddImageResponse, error) {
	res, err := u.User.AddImage(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error adding image service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetRegister(ctx context.Context, req *pb.GetRegisterRequest) (*pb.GetRegisterResponse, error) {
	res, err := u.User.GetRegister(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error getting register service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetRegisters(ctx context.Context, req *pb.GetRegistersRequest) (*pb.GetRegistersResponse, error) {
	res, err := u.User.GetRegisters(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error getting registers service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) DeleteRegister(ctx context.Context, req *pb.DeleteRegisterRequest) (*pb.DeleteRegisterResponse, error) {
	res, err := u.User.DeleteRegister(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error deleting register service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	res, err := u.User.UpdatePassword(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error updating password service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetByEmail(ctx context.Context, req *pb.GetByEmailRequest) (*pb.GetByEmailResponse, error) {
	res, err := u.User.GetByEmail(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error getting by email service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	res, err := u.User.UpdateRole(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error updating role service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) Tobeanadmin(ctx context.Context, req *pb.TobeanadminRequest) (*pb.TobeanadminResponse, error) {
	res, err := u.User.Tobeanadmin(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error to be an admin service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (u *UserService) CheckUserId(ctx context.Context, req *pb.CheckUserIdRequest) (*pb.CheckUserIdResponse, error) {
	res, err := u.User.CheckUserId(ctx, req)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error checking user id service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}