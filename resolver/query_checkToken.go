package resolver

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/api"
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

	tokenNode, err := r.EntityDB.GetEntityByID(ctx, claim.Id)
	if err != nil {
		return nil, fmt.Errorf("CheckToken: failed geting ownerID by token, %s", err)
	}

	apiToken, ok := tokenNode.(*api.Token)
	if !ok {
		return nil, fmt.Errorf("CheckToken: failed to convert entity to api.token, %s", err)
	}

	userNode, err := r.EntityDB.GetEntityByID(ctx, apiToken.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("CheckToken: failed geting user by token, %s", err)
	}

	user, ok := userNode.(*api.User)
	if !ok {
		return nil, fmt.Errorf("CheckToken: failed to convert entity to api.user, %s", err)
	}

	return node.NewUserResolver(node.WithNode(user), node.WithDatastore(r.EntityDB)), nil
}
