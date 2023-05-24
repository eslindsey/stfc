package stfc

import (
	"io/ioutil"
	"testing"

	"github.com/go-yaml/yaml"
)

var secrets struct {
	AdhocUsername string `yaml:"adhoc_username"`
	AdhocPassword string `yaml:"adhoc_password"`
}

var (
	s *Session
)

func TestAll(t *testing.T) {
	b, err := ioutil.ReadFile(".secrets.yaml")
	if err != nil {
		t.Fatalf("Couldn't read secrets file: %s", err)
	}
	err = yaml.Unmarshal(b, &secrets)
	if err != nil {
		t.Fatalf("Couldn't unmarshal secrets: %s", err)
	}
	t.Run("Login", testLogin)
	t.Run("Sync2", testSync2)
}

func testLogin(t *testing.T) {
	var err error
	s, err = Login(secrets.AdhocUsername, secrets.AdhocPassword)
	if err != nil {
		t.Fatalf("Login failed: %s", err)
	}
	r := s.LoginResponse
	t.Logf("Login succeeded as %s with session %s to instance %03d", r.InstanceAccount.Name, r.InstanceSessionID, r.InstanceAccount.InstanceIDCurrent)
}

func testSync2(t *testing.T) {
	body, err := s.Sync(2)
	if err != nil {
		t.Fatalf("Sync2 failed: %s", err)
	}
	t.Logf("Sync2 succeeded: %v", string(body))
}

