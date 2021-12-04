package gateway

import "demo-go/pkg/web/domain/model"

type UserGateway interface {
	Get(id int64) *model.User

	Del(id int64) bool
}
