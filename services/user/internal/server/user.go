package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
	"pad/services/common/hash"
	"pad/services/common/jwt"
	"pad/services/user/internal/store"
	"pad/services/user/services/user"
)

type User struct {
	datastore store.User

	user.UnimplementedUserServer
}

func NewUserServer(db store.User) user.UserServer {
	return &User{
		datastore: db,
	}
}

func (u *User) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	creds, err := u.datastore.GetCredentials(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	if !hash.ValidatePassword(req.Password, creds.Password) {
		return nil, fmt.Errorf("unauthorized")
	}

	userToken, err := jwt.CreateToken(creds.ID)
	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		Token: userToken,
	}, nil
}

func (u *User) Register(ctx context.Context, req *user.RegisterRequest) (*emptypb.Empty, error) {
	hashedPassword, err := hash.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if err := u.datastore.CreateCredentials(ctx, &store.Credentials{
		Name:     req.Name,
		Password: hashedPassword,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
