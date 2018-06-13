package token

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestToken_isValid_ReturnsNoError(t *testing.T) {
	private, err := ioutil.ReadFile("testkeys/private.pem")
	if err != nil {
		t.Fatalf("failed to read rsa key %s", err)
	}
	RsaPrivateKey = private

	public, err := ioutil.ReadFile("testkeys/public.pem")
	if err != nil {
		t.Fatalf("failed to read rsa key %s", err)
	}
	RsaPublicKey = public

	token, err := Create("1", time.Now().Add(time.Hour*10).Unix())
	if err != nil {
		t.Fatalf("failed to create rsa key %s", err)
	}

	if _, err := Valid(token); err != nil {
		t.Fatalf("token is invalid, %s", err)
	}
}

func TestToken_isExpired_ReturnsError(t *testing.T) {
	private, err := ioutil.ReadFile("testkeys/private.pem")
	if err != nil {
		t.Fatalf("failed to read rsa key %s", err)
	}
	RsaPrivateKey = private

	public, err := ioutil.ReadFile("testkeys/public.pem")
	if err != nil {
		t.Fatalf("failed to read rsa key %s", err)
	}
	RsaPublicKey = public

	token, err := Create("1", time.Now().Unix()-1)
	if err != nil {
		t.Fatalf("failed to create rsa key %s", err)
	}

	if _, err := Valid(token); err == nil {
		t.Fatalf("validate token, %s", err)
	}
}

func TestToken_WrongPubKey_ReturnsError(t *testing.T) {
	private, err := ioutil.ReadFile("testkeys/private.pem")
	if err != nil {
		t.Fatalf("failed to read rsa key %s", err)
	}
	RsaPrivateKey = private

	RsaPublicKey = []byte("wrong key")

	token, err := Create("1", time.Now().Add(time.Second).Unix())
	if err != nil {
		t.Fatalf("failed to create rsa key %s", err)
	}
	if _, err := Valid(token); err == nil {
		t.Fatalf("validate token, %s", err)
	}
}

func TestToken_check_ID_ReturnsNoError(t *testing.T) {
	private, err := ioutil.ReadFile("testkeys/private.pem")
	if err != nil {
		t.Fatalf("failed to read rsa key %s", err)
	}
	RsaPrivateKey = private

	public, err := ioutil.ReadFile("testkeys/public.pem")
	if err != nil {
		t.Fatalf("failed to read rsa key %s", err)
	}
	RsaPublicKey = public

	token, err := Create("1", time.Now().Add(time.Second).Unix())
	if err != nil {
		t.Fatalf("failed to create rsa key %s", err)
	}

	actual, err := Valid(token)
	if err != nil {
		t.Fatalf("validate token, %s", err)
	}

	if actual.Id != "1" {
		t.Fatalf("id is not equals")
	}
}
