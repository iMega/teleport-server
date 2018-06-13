package resolver

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/health"
	"github.com/imega/teleport-server/mysql"
	"github.com/imega/teleport-server/shutdown"
	"github.com/imega/teleport-server/token"
	"github.com/imega/teleport-server/uuid"
)

type mockDatastore struct{}

func (mockDatastore) CreateEntity(context.Context, mysql.Entity) (mysql.Entity, error) {
	return &api.User{}, nil
}
func (mockDatastore) GetEntityByID(context.Context, string) (mysql.Entity, error) {
	user := &api.User{
		Id: "112233",
	}
	return user, nil
}
func (mockDatastore) RemoveEntity(context.Context, uuid.UID) error {
	return nil
}
func (mockDatastore) CreateRelation(ctx context.Context, subject uuid.UID, predicate string, object uuid.UID, priority int) error {
	return nil
}
func (mockDatastore) DeleteRelation(ctx context.Context, subject, object uuid.UID) error {
	return nil
}
func (mockDatastore) HealthCheckFunc() health.HealthCheckFunc {
	return func() bool { return false }
}
func (mockDatastore) ShutdownFunc() shutdown.ShutdownFunc {
	return func() {}
}

func createToken() (string, error) {
	private, err := ioutil.ReadFile("../token/testkeys/private.pem")
	if err != nil {
		return "", err
	}
	token.RsaPrivateKey = private

	public, err := ioutil.ReadFile("../token/testkeys/public.pem")
	if err != nil {
		return "", err
	}
	token.RsaPublicKey = public

	token, err := token.Create("1", time.Now().Add(time.Hour*10).Unix())
	if err != nil {
		return "", err
	}
	return token, nil
}

func Test_checkToken(t *testing.T) {
	r := Resolver{
		EntityDB: mockDatastore{},
	}
	tokenStr, err := createToken()
	if err != nil {
		t.Errorf("failed to create token, %s", err)
	}

	userNode, err := r.CheckToken(context.Background(), QueryCheckTokenInput{Token: &tokenStr})
	if err != nil {
		t.Errorf("failed to check token, %s", err)
	}

	if userNode.ID() != "112233" {
		t.Error("user ID is not equal")
	}
}
