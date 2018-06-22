package resolver

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/resolver/node"
	"github.com/imega/teleport-server/uuid"
)

// CreateUserInput переданные аргументы для создания пользователя
type CreateUserInput struct {
	Pass string
}

// HasPass проверка, что значение пароль задано
func (i *CreateUserInput) HasPass() bool {
	if len(i.Pass) < 1 {
		return false
	}
	return true
}

// GetPass получение пароля из переданных аргументов
func (i *CreateUserInput) GetPass() *string {
	if i.HasPass() {
		return &i.Pass
	}
	return nil
}

// CreateUser создание пользователя
func (r *Resolver) CreateUser(ctx context.Context, args CreateUserInput) (*node.UserResolver, error) {
	if !args.HasPass() {
		return nil, fmt.Errorf("password is empty")
	}

	ownerID := string(uuid.NewUUID())
	ctx = context.WithValue(ctx, "owner_id", ownerID)

	entity, err := r.EntityDB.CreateEntity(ctx, &api.User{
		Id:     ownerID,
		Pass:   *args.GetPass(),
		Active: false,
	})
	if err != nil {
		return nil, err
	}

	return node.NewUserResolver(node.WithNode(entity), node.WithDatastore(r.EntityDB)), nil
}
