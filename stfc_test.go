package stfc

import (
	"io/ioutil"
	"testing"

	"github.com/go-yaml/yaml"
)

var secrets struct {
	AdhocUsername          string   `yaml:"adhoc_username"`
	AdhocPassword          string   `yaml:"adhoc_password"`
	ProfilesUserIDs        []string `yaml:"profiles_user_ids"`
	AlliancesPublicInfoIDs []uint64 `yaml:"alliances_public_info_ids"`
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
	t.Run("Login",          testLogin)
	t.Run("Sync2",          testSync2)
	if len(secrets.ProfilesUserIDs) > 0 {
		t.Run("Profiles",       testProfiles)
	}
	if len(secrets.AlliancesPublicInfoIDs) > 0 {
		//t.Run("AlliancesProto", testAlliancesProto)
		t.Run("AlliancesJson",  testAlliancesJson)
	}
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

//func testAlliancesProto(t *testing.T) {
//	alliances, err := s.AlliancesProto(secrets.AlliancesPublicInfoIDs)
//	if err != nil {
//		t.Fatalf("AlliancesProto failed: %s", err)
//	}
//	t.Logf("AlliancesProto succeeded: %v", alliances)
//}

func testAlliancesJson(t *testing.T) {
	alliances, err := s.AlliancesJson(secrets.AlliancesPublicInfoIDs)
	if err != nil {
		t.Fatalf("AlliancesJson failed: %s", err)
	}
	t.Logf("AlliancesJson succeeded: %+v", alliances)
}

