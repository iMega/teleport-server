package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/owner"
	"github.com/imega/teleport-server/token"
	"github.com/imega/teleport-server/uuid"
	"github.com/improbable-eng/go-httpwares/logging/logrus/ctxlogrus"
)

type CreateTokenInput struct {
	ID   graphql.ID
	Pass string
}

// GetID получение ID
func (i *CreateTokenInput) GetID() string {
	return string(i.ID)
}

// HasPass проверка, что значение пароль задано
func (i *CreateTokenInput) HasPass() bool {
	if len(i.Pass) < 1 {
		return false
	}
	return true
}

// GetPass получение пароля из переданных аргументов
func (i *CreateTokenInput) GetPass() *string {
	if i.HasPass() {
		return &i.Pass
	}
	return nil
}

// CreateToken создание токена
func (r *Resolver) CreateToken(ctx context.Context, in CreateTokenInput) (string, error) {
	var errorFailCreateToken = fmt.Errorf("failed to create token")

	logger := ctxlogrus.Extract(ctx)
	ownerID, err := owner.GetOwnerIDFromContext(ctx)
	if err != nil {
		logger.Errorf("CreateToken: failed getting ownerID, %s", err)
		return "", err
	}

	if !in.HasPass() {
		logger.Errorf("CreateToken: password is empty")
		return "", fmt.Errorf("password is empty")
	}

	userNode, err := r.EntityDB.GetEntityByID(ctx, string(ownerID))
	if err != nil {
		logger.Errorf("CreateToken: failed getting user record, %s", err)
		return "", fmt.Errorf("failed getting user")
	}

	user, ok := userNode.(*api.User)
	if !ok {
		logger.Errorf("CreateToken: failed to convert entity to api.user, %s", err)
		return "", errorFailCreateToken
	}

	if *in.GetPass() != user.GetPass() {
		logger.Errorf("CreateToken: password is invalid")
		return "", errorFailCreateToken
	}

	entityID := uuid.NewUUID()
	expireAt := time.Now().Add(time.Hour * 10).Unix()
	tokenStr, err := token.Create(string(entityID), expireAt)
	if err != nil {
		logger.Errorf("CreateToken: %s", err)
		return "", fmt.Errorf("CreateToken, %s", err)
	}

	_, err = r.EntityDB.CreateEntity(ctx, &api.Token{
		Id:      string(entityID),
		Value:   tokenStr,
		OwnerId: string(ownerID),
	})
	if err != nil {
		logger.Errorf("CreateToken: %s", err)
		return "", fmt.Errorf("CreateToken, %s", err)
	}

	return tokenStr, nil
}
