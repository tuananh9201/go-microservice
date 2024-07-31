package main

import (
	"context"
	"log"
	"strconv"

	common "github.com/tuananh9201/commons"
	pb "github.com/tuananh9201/commons/api"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewGRPCHandler(grpcServer *grpc.Server, db *gorm.DB) {
	log.Println("users: NewGRPCHandler")
	handler := &grpcHandler{}
	pb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateNewUserhandler(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	log.Println("users: New User received")
	user := &User{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Password:   req.Password,
		Role:       "user",
		DeleteFlag: false,
	}
	store := NewSQLStore(h.db)
	uc := NewUserUseCase(store)
	if err := uc.CreateNewUser(ctx, user); err != nil {
		return nil, err
	}
	return &pb.User{
		ID:         strconv.FormatUint(uint64(user.ID), 10),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Role:       string(user.Role),
		DeleteFlag: user.DeleteFlag,
	}, nil
}

func (h *grpcHandler) GetUsersHandler(ctx context.Context, req *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	log.Println("users: Get Users received")
	paging := &common.Paging{
		Page:  int(req.Paging.Page),
		Limit: int(req.Paging.Limit),
		Total: 0,
	}
	filter := &UserFilter{
		FirstName:  req.Filter.FirstName,
		LastName:   req.Filter.LastName,
		Email:      req.Filter.Email,
		DeleteFlag: req.Filter.DeleteFlag,
	}
	store := NewSQLStore(h.db)
	uc := NewUserUseCase(store)
	users, p, err := uc.GetUsers(ctx, paging, filter)
	if err != nil {
		return nil, err
	}
	// Convert the list of users to the protobuf type
	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			ID:         strconv.FormatUint(uint64(u.ID), 10),
			FirstName:  u.FirstName,
			LastName:   u.LastName,
			Email:      u.Email,
			DeleteFlag: u.DeleteFlag,
		})
	}

	// Create the response with users, updated paging, and filter
	resp := &pb.GetListUserResponse{
		Data: pbUsers,
		Paging: &pb.Paging{
			Page:  int32(p.Page),
			Limit: int32(p.Limit),
			Total: int32(p.Total),
		},
		Filter: &pb.UserFilter{
			FirstName:  req.Filter.FirstName,
			LastName:   req.Filter.LastName,
			Email:      req.Filter.Email,
			DeleteFlag: req.Filter.DeleteFlag,
		},
	}

	return resp, nil
}
