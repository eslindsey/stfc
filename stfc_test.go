package stfc

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/go-yaml/yaml"
)

var secrets struct {
	AdhocUsername   string   `yaml:"adhoc_username"`
	AdhocPassword   string   `yaml:"adhoc_password"`
	ProfilesUserIDs []string `yaml:"profiles_user_ids"`
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
	//log.Printf("Using secrets: %+v", secrets)
	t.Run("Login",    testLogin)
	t.Run("Sync2",    testSync2)
	t.Run("Profiles", testProfiles)
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
	sync, err := s.Sync(2)
	if err != nil {
		t.Fatalf("Sync2 failed: %s", err)
	}
	t.Logf("Sync2 succeeded: user ID %v has %d types of resources", sync.Starbase.UserID, len(sync.Resources))
}

func testProfiles(t *testing.T) {
	profiles, err := s.Profiles(secrets.ProfilesUserIDs)
	if err != nil {
		t.Fatalf("Profiles failed: %s", err)
	}
	t.Logf("Profiles succeeded: %v", profiles)
}

