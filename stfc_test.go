package stfc

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"testing"

	"github.com/go-yaml/yaml"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/path"
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
	g *GalaxyOptimised
	graph *simple.WeightedDirectedGraph
)

func TestAll(t *testing.T) {
	t.Run("Nodes",        testNodes)
	t.Run("BuildGraph",   testBuildGraph)
	t.Run("ShortestPath", testShortestPath)

	b, err := ioutil.ReadFile(".secrets.yaml")
	if err != nil {
		t.Fatalf("Couldn't read secrets file: %s", err)
	}
	err = yaml.Unmarshal(b, &secrets)
	if err != nil {
		t.Fatalf("Couldn't unmarshal secrets: %s", err)
	}
	//log.Printf("Using secrets: %+v", secrets)
	t.Run("ScopelyID",      testScopelyID)
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

func testNodes(t *testing.T) {
	b, err := ioutil.ReadFile("static/nodes.json")
	if err != nil {
		t.Fatalf("Nodes failed: %s", err)
	}
	err = json.Unmarshal(b, &g)
	if err != nil {
		t.Fatalf("Nodes failed: %s", err)
	}
	t.Logf("Nodes succeeded")
	t.Logf("Nodes: %d", len(g.Nodes))
	t.Logf("Paths: %d", len(g.Paths))
}

func testBuildGraph(t *testing.T) {
	graph = simple.NewWeightedDirectedGraph(0, math.Inf(1))
	for _, path := range g.Paths {
		graph.SetWeightedEdge(simple.WeightedEdge{
			F: simple.Node(path.SourceId),
			T: simple.Node(path.DestId),
			W: float64(path.Distance),
		})
	}
}

func testShortestPath(t *testing.T) {
	start  := 55         // Yridia System
	finish := 119043632  // Teranth System
	shortest := path.DijkstraFrom(simple.Node(start), graph)
	path, weight := shortest.To(int64(finish))
	t.Logf("ShortestPath from %d to %d succeeded (%d hops, %.0f distance)", start, finish, len(path), weight)
	t.Logf("Path: %v", path)
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

