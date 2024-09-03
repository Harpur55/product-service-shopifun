package ports

import (
	oauthgoogleent "codebase-app/internal/integration/oauth2google/entity"
	"codebase-app/internal/module/user/entity"
	"context"
)

type UserRepository interface {
	Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error)
	FindByEmail(ctx context.Context, email string) (*entity.UserResult, error)
	FindById(ctx context.Context, id string) (*entity.ProfileResponse, error)
}

type UserService interface {
	Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error)
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	Profile(ctx context.Context, req *entity.ProfileRequest) (*entity.ProfileResponse, error)
	GetOauthGoogleUrl(ctx context.Context) (string, error)
	LoginGoogle(ctx context.Context, req *oauthgoogleent.UserInfoResponse) (*entity.LoginResponse, error)
}
