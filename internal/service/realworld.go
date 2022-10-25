package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.RealWorldUsecase
}

func NewRealWorldService(uc *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
