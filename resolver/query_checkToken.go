package resolver

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/resolver/node"
)

// QueryCheckTokenInput переданные аргументы для проверки токена
type QueryCheckTokenInput struct {
	Token *string
}

// HasToken проверка, что значение токен задано
func (i *QueryCheckTokenInput) HasToken() bool {
	if i.Token == nil {
		return false
	}
	return true
}

// GetToken получение токена из переданных аргументов
func (i *QueryCheckTokenInput) GetToken() *string {
	if i.HasToken() {
		return i.Token
	}
	return nil
}

// CheckToken проверка токена
func (r *Resolver) CheckToken(ctx context.Context, in QueryCheckTokenInput) (*node.UserResolver, error) {
	if !in.HasToken() {
		return nil, fmt.Errorf("token is empty")
	}
	return nil, nil
}
