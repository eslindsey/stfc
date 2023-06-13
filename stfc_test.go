package stfc

import (
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/go-yaml/yaml"
)

var secrets struct {
	ScopelyIDEmail         string   `yaml:"scopelyid_email"`
	ScopelyIDPassword      string   `yaml:"scopelyid_password"`
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
	t.Run("ScopelyID",    testScopelyID)
	t.Run("Login",        testLogin)
	t.Run("ShortestPath", testShortestPath)  // needed to ensure Galaxy is populated
	t.Run("Sync2",        testSync2)
	if len(secrets.ProfilesUserIDs) > 0 {
		t.Run("Profiles",       testProfiles)
	}
	if len(secrets.AlliancesPublicInfoIDs) > 0 {
		//t.Run("AlliancesProto", testAlliancesProto)
		t.Run("AlliancesJson",  testAlliancesJson)
	}
}

func testScopelyID(t *testing.T) {
	_, err := ScopelyID(secrets.ScopelyIDEmail, secrets.ScopelyIDPassword)
	if err != nil {
		t.Fatalf("Scopely ID failed: %s", err)
	}
	t.Logf("Scopely ID succeeded")
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

func testShortestPath(t *testing.T) {
	g, err := s.Galaxy()
	if err != nil {
		t.Fatalf("Couldn't get galaxy: %s", err)
	}
	n := len(g.Nodes)
	n1 := g.Nodes[rand.Intn(n)]
	n2 := g.Nodes[rand.Intn(n)]
	path, distance := g.Shortest(n1.Id, n2.Id)
	t.Logf("Shortest path from %d to %d (%d hops, %d distance):", n1.Id, n2.Id, len(path), distance)
	t.Logf("%v", path)
}

func testSync2(t *testing.T) {
	sync, err := s.Sync(2)
	if err != nil {
		t.Fatalf("Sync2 failed: %s", err)
	}
	g, err := s.Galaxy()
	if err != nil {
		t.Fatalf("Couldn't get galaxy: %s", err)
	}
	var visited float32
	for _, v := range sync.VisitedSystems {
		if v {
			visited++
		}
	}
	n := float32(len(g.Nodes))
	pct1 := float32(len(sync.VisitedSystems)) / n * 100.0
	pct2 := visited / n * 100.0
	t.Logf(`Sync2 succeeded:
User ID:    %v
Resources:  %d types
Visit Data: info for %d of %.0f total (%.1f%%), visited %.0f (%.1f%%)`, sync.Starbase.UserID, len(sync.Resources), len(sync.VisitedSystems), n, pct1, visited, pct2)
}

func testProfiles(t *testing.T) {
	profiles, err := s.UserProfileProfiles(secrets.ProfilesUserIDs)
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

