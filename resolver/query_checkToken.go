package resolver

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/resolver/node"
	"github.com/imega/teleport-server/token"
)

// QueryCheckTokenInput переданные аргументы для проверки токена
type QueryCheckTokenInput struct {
	Token string
}

// HasToken проверка, что значение токен задано
func (i *QueryCheckTokenInput) HasToken() bool {
	if len(i.Token) < 1 {
		return false
	}
	return true
}

// GetToken получение токена из переданных аргументов
func (i *QueryCheckTokenInput) GetToken() *string {
	if i.HasToken() {
		return &i.Token
	}
	return nil
}

// CheckToken проверка токена
func (r *Resolver) CheckToken(ctx context.Context, in QueryCheckTokenInput) (*node.UserResolver, error) {
	if !in.HasToken() {
		return nil, fmt.Errorf("token is empty")
	}
	claim, err := token.Valid(*in.GetToken())
	if err != nil {
		return nil, err
	}

	n, err := r.EntityDB.GetEntityByID(ctx, claim.Id)

	if err != nil {
		return nil, fmt.Errorf("CheckToken: failed geting login by token, %s", err)
	}
	return node.NewUserResolver(node.WithNode(n), node.WithDatastore(r.EntityDB)), nil
}
