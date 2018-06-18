package resolver

import (
	"context"
	"io/ioutil"
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/imega/teleport-server/token"
)

func initCertificates(t *testing.T) {
	private, err := ioutil.ReadFile("../token/testkeys/private.pem")
	if err != nil {
		t.Error("failed to open private key")
	}
	token.RsaPrivateKey = private

	public, err := ioutil.ReadFile("../token/testkeys/public.pem")
	if err != nil {
		t.Error("failed to open public key")
	}
	token.RsaPublicKey = public
}

func Test_CreateToken(t *testing.T) {
	initCertificates(t)

	pass := "pass"
	in := CreateTokenInput{
		ID:   graphql.ID("id"),
		Pass: &pass,
	}
	r := Resolver{
		EntityDB: mockDatastore{},
	}
	_, err := r.CreateToken(context.Background(), in)
	if err != nil {
		t.Errorf("failed to create token, %s", err)
	}
}
