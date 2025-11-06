package session

import (
	"context"
	"errors"
	"broker/pkg/clients"
)

var (
	ErrNoAuth = errors.New("Unauthorized")
)

// сессии не храню, ибо они хранятся на клиентах
// (и js их отправляет в апи когда хочет)

type sessKey string

var SessionKey sessKey = "sessionKey"

// возращаем не сессию а пользователя, ибо это всё, что надо
func ClientFromContext(ctx context.Context) (*clients.Client, error) {
	u, ok := ctx.Value(SessionKey).(*clients.Client)
	if !ok || u == nil {
		return nil, ErrNoAuth
	}
	return u, nil
}
