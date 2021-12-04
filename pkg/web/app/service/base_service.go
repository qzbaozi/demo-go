package service

import (
	service "demo-go/pkg/web/client/api"
	"demo-go/pkg/web/domain/context"
)

type ServiceFactoryInterface interface {
	UserService(ctx context.DemoContext) service.UserService
}

var factory = new(serviceFactory)

type serviceFactory struct {
}

func UserService(ctx context.DemoContext) service.UserService {
	return factory.UserService(ctx)
}

type baseService struct {
	context.DemoContext
}

func (g serviceFactory) UserService(ctx context.DemoContext) service.UserService {
	return &userServiceImpl{baseService{context.NewDemoContext(ctx)}}
}
