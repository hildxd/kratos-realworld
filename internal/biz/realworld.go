package biz

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// RealWorldRepo is a Greater repo.
type RealWorldRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// RealWorldUsecase is a Greeter usecase.
type RealWorldUsecase struct {
	repo RealWorldRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo RealWorldRepo, logger log.Logger) *RealWorldUsecase {
	return &RealWorldUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRealworld creates a Greeter, and returns the new Greeter.
func (uc *RealWorldUsecase) CreateRealworld(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateRealworld: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
