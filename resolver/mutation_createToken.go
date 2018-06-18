package resolver

import (
	"context"
	"fmt"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/token"
)

type CreateTokenInput struct {
	ID   graphql.ID
	Pass *string
}

// GetID получение ID
func (i *CreateTokenInput) GetID() string {
	return string(i.ID)
}

// HasPass проверка, что значение пароль задано
func (i *CreateTokenInput) HasPass() bool {
	if i.Pass == nil {
		return false
	}
	return true
}

// GetPass получение пароля из переданных аргументов
func (i *CreateTokenInput) GetPass() *string {
	if i.HasPass() {
		return i.Pass
	}
	return nil
}

// CreateToken создание токена
func (r *Resolver) CreateToken(ctx context.Context, in CreateTokenInput) (string, error) {
	if !in.HasPass() {
		return "", fmt.Errorf("password is empty")
	}

	expireAt := time.Now().Add(time.Hour * 10).Unix()
	tokenStr, err := token.Create(in.GetID(), expireAt)
	if err != nil {
		return "", fmt.Errorf("CreateToken, %s", err)
	}

	_, err = r.EntityDB.CreateEntity(ctx, &api.Token{
		Id:    in.GetID(),
		Value: tokenStr,
	})
	if err != nil {
		return "", fmt.Errorf("CreateToken, %s", err)
	}

	return tokenStr, nil
}
