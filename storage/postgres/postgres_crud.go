package postgres

import (
	pb "auth/genproto/register"
	"auth/pkg/logger"
	"auth/storage"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewUserRepository(db *sql.DB) storage.IUserStorage {
	return &UserRepository{Db: db, Log: logger.NewLogger()}
}

func (u *UserRepository) CreateRegister(ctx context.Context, req *pb.CreateRegisterRequest) (*pb.CreateRegisterResponse, error) {
	query := `INSERT INTO register (
					id, email, first_name, last_name, phonenummer, created_at, updated_at
				) VALUES (
				 	$1, $2, $3, $4, $5, $6, $7)`

	newtime := time.Now()
	id := uuid.New().String()
	_, err := u.Db.Exec(query, id, req.Email, req.FirstName, req.LastName, req.Phonenummer, newtime, newtime)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while creating register: %v", err))
		return nil, err
	}

	query = `select role from register where email = $1 and deleted_at = 0`
	var role string
	err = u.Db.QueryRow(query, req.Email).Scan(&role)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while getting role: %v", err))
		return nil, err
	}

	register := &pb.Register{
		Id:          id,
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phonenummer: req.Phonenummer,
		Role:        role,
	}

	return &pb.CreateRegisterResponse{Register: register}, nil
}

func (u *UserRepository) Update(ctx context.Context,req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	query_update := `UPDATE register SET email = $1, first_name = $2, last_name = $3, phonenummer = $4, updated_at = $5 WHERE id = $6 and deleted_at = 0`

	query_old := `select id, email, first_name, last_name, phonenummer from register where id = $1`
	var old struct {
		Id string
		Email string
		FirstName string
		LastName string
		PhoneNummer string
	}

	err := u.Db.QueryRow(query_old,req.Id).Scan(&old.Id, &old.Email, &old.FirstName, &old.LastName, &old.PhoneNummer)
	if err != nil {
		u.Log.Error(fmt.Sprintf("error retrieving old data: %v", err))
		return nil, err
	}

	if req.Email == "" || old.Email == req.Email {
		req.Email = old.Email
	}

	if req.FirstName == "" || old.FirstName == req.FirstName {
		req.FirstName = old.FirstName
	}

	if req.LastName == "" || old.LastName == req.LastName {
		req.LastName = old.LastName
	}

	if req.Phonenummer == "" || old.PhoneNummer == req.Phonenummer {
		req.Phonenummer = old.PhoneNummer
	}

	newtime := time.Now()
	_, err = u.Db.Exec(query_update, req.Email, req.FirstName, req.LastName, req.Phonenummer, newtime, req.Id)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while updating register: %v", err))
		return nil, err
	}
 
	query_role := `select role from register where id = $1 and deleted_at = 0`
	var role string
	err = u.Db.QueryRow(query_role, req.Id).Scan(&role)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while getting role: %v", err))
		return nil, err
	}

	register := &pb.Register{
		Id:          req.Id,
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phonenummer: req.Phonenummer,
		Role:        role,
	}

	return &pb.UpdateResponse{Register: register}, nil
}

func (u *UserRepository) AddImage(ctx context.Context,req *pb.AddImageRequest) (*pb.AddImageResponse, error) {
	query := `UPDATE register SET image = $1 WHERE id = $2 and deleted_at = 0`

	_, err := u.Db.Exec(query, req.Image, req.Id)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while adding image: %v", err))
		return nil, err
	}

	return &pb.AddImageResponse{Id: req.Id, Image: req.Image}, nil
}

func (u *UserRepository) GetRegister(ctx context.Context,req *pb.GetRegisterRequest) (*pb.GetRegisterResponse, error) {
	query := `SELECT id, email, first_name, last_name, phonenummer, password, image, role FROM register WHERE id = $1 and deleted_at = 0`

	var id, email, first_name, last_name, phonenummer, password, image, role string
	err := u.Db.QueryRow(query, req.Id).Scan(&id, &email, &first_name, &last_name, &phonenummer, &password, &image, &role)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while getting register: %v", err))
		return nil, err
	}

	register := &pb.Getregister{
		Id:          id,
		Email:       email,
		FirstName:   first_name,
		LastName:    last_name,
		Phonenummer: phonenummer,
		Password:    password,
		Image:       image,
		Role:        role,
	}

	return &pb.GetRegisterResponse{Register: register}, nil
}

func (u *UserRepository) GetRegisters(ctx context.Context, req *pb.GetRegistersRequest) (*pb.GetRegistersResponse, error) {
	limit := req.Limit
	page := req.Page
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	query := `SELECT id, email, first_name, last_name, phonenummer, password, image, role FROM register 
			  WHERE deleted_at = 0 
			  ORDER BY created_at DESC 
			  LIMIT $1 OFFSET $2`

	rows, err := u.Db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		u.Log.ErrorContext(ctx, fmt.Sprintf("Error reading register: %v", err))
		return nil, err
	}
	defer rows.Close()

	var registers []*pb.Getregister
	for rows.Next() {
		var reg pb.Getregister
		err := rows.Scan(
			&reg.Id,
			&reg.Email,
			&reg.FirstName,
			&reg.LastName,
			&reg.Phonenummer,
			&reg.Password,
			&reg.Image,
			&reg.Role,
		)
		if err != nil {
			u.Log.ErrorContext(ctx, fmt.Sprintf("Error scanning row: %v", err))
			return nil, err
		}
		registers = append(registers, &reg)
	}

	if err = rows.Err(); err != nil {
		u.Log.ErrorContext(ctx, fmt.Sprintf("Error iterating rows: %v", err))
		return nil, err
	}

	totalQuery := `SELECT COUNT(*) FROM register WHERE deleted_at = 0`
	var total int32
	err = u.Db.QueryRowContext(ctx, totalQuery).Scan(&total)
	if err != nil {
		u.Log.ErrorContext(ctx, fmt.Sprintf("Error getting total count: %v", err))
		return nil, err
	}

	return &pb.GetRegistersResponse{
		Registers: registers,
		Total:     total,
	}, nil
}

func (u *UserRepository) DeleteRegister(ctx context.Context,req *pb.DeleteRegisterRequest) (*pb.DeleteRegisterResponse, error) {
	query := `UPDATE register SET deleted_at = $1 WHERE id = $2 and deleted_at = 0`

	newtime := time.Now()
	_, err := u.Db.Exec(query, newtime, req.Id)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while adding image: %v", err))
		return nil, err
	}

	return &pb.DeleteRegisterResponse{Message: "User deleted"},nil
}

func (u *UserRepository) GetByEmail(ctx context.Context,req *pb.GetByEmailRequest) (*pb.GetByEmailResponse, error) {
	query := `SELECT id, email, first_name, last_name, phonenummer, password, image, role FROM register WHERE email = $1 and deleted_at = 0`

	res := &pb.GetByEmailResponse{}
	err := u.Db.QueryRow(query, req.Email).Scan(&res.Id, &res.Email, &res.FirstName, &res.LastName, &res.Phonenummer, &res.Password, &res.Image, &res.Role)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while getting register: %v", err))
		return nil, err
	}

	return res, nil
}

func (u *UserRepository) UpdatePassword(ctx context.Context,req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	query := `UPDATE register SET password = $1 WHERE email = $2 and deleted_at = 0`
	_,err := u.Db.Exec(query, req.NewPassword, req.Email)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while updating password: %v", err))
		return nil, err
	}	

	return &pb.UpdatePasswordResponse{Message: "Password updated"}, nil
}

func (u *UserRepository) UpdateRole (ctx context.Context,req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	query := `UPDATE register SET role = $1 WHERE id = $2 and deleted_at = 0`
	_,err := u.Db.Exec(query, req.Role, req.Id)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while updating role: %v", err))
		return nil, err
	}	

	newtime := time.Now()
	query_adminn := `update adminn set deleted_at = $1 where user_id = $2 and deleted_at = 0`
	_,err = u.Db.Exec(query_adminn, newtime, req.Id)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while updating adminn: %v", err))
		return nil, err
	}

	return &pb.UpdateRoleResponse{Message: "Role updated"}, nil
}

func (u *UserRepository) Tobeanadmin (ctx context.Context,req *pb.TobeanadminRequest) (*pb.TobeanadminResponse, error) {
	query := `insert into adminn (id, user_id, role, created_at, updated_at) values ($1, $2, $3, $4, $5)`
	newtime := time.Now()
	id := uuid.New().String()

	query_role := `select role from register where id = $1 and deleted_at = 0`
	var role string
	err := u.Db.QueryRow(query_role, req.Id).Scan(&role)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while getting role: %v", err))
		return nil, err
	}

	_,err = u.Db.Exec(query, id, req.Id, role, newtime, newtime)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error while inserting into admin: %v", err))
		return nil, err
	}

	return &pb.TobeanadminResponse{Message: "You have been sent to review your material."}, nil
}

func (u *UserRepository) CheckUserId (ctx context.Context,req *pb.CheckUserIdRequest) (*pb.CheckUserIdResponse, error) {
	query := `select id from register where deleted_at = 0`

	rows, err := u.Db.Query(query)
	if err != nil {
		return &pb.CheckUserIdResponse{IsExist: false}, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return &pb.CheckUserIdResponse{IsExist: false}, err
		}
		if id == req.Id {
			return &pb.CheckUserIdResponse{IsExist: true}, nil
		}
	}

	if err = rows.Err(); err != nil {
		return &pb.CheckUserIdResponse{IsExist: false}, err
	}

	return &pb.CheckUserIdResponse{IsExist: false}, nil
}