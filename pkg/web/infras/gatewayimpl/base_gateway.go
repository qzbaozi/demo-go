package gatewayimpl

import (
	"demo-go/pkg/web/domain/context"
	"demo-go/pkg/web/domain/gateway"
)

type GatewayFactoryInterface interface {
	UserGateWay(ctx context.DemoContext) gateway.UserGateway
}

var factory = new(gatewayFactory)

func UserGateWay(ctx context.DemoContext) gateway.UserGateway {
	return factory.UserGateWay(ctx)
}

type gatewayFactory struct {
}

func (g gatewayFactory) UserGateWay(ctx context.DemoContext) gateway.UserGateway {
	return &userGatewayImpl{baseGateWay{context.NewDemoContext(ctx)}}
}

type baseGateWay struct {
	context.DemoContext
}
