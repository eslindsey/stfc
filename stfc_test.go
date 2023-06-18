package stfc

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"testing"

	"github.com/go-yaml/yaml"
)

var secrets struct {
	ScopelyIdEmail         string   `yaml:"scopelyid_email"`
	ScopelyIdPassword      string   `yaml:"scopelyid_password"`
	AdhocUsername          string   `yaml:"adhoc_username"`
	AdhocPassword          string   `yaml:"adhoc_password"`
	ProfilesUserIds        []string `yaml:"profiles_user_ids"`
	AlliancesPublicInfoIds []uint64 `yaml:"alliances_public_info_ids"`
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
	t.Run("ScopelyId",     testScopelyId)
	t.Run("Login",         testLogin)
	t.Run("ShortestPath",  testShortestPath)  // needed to ensure Galaxy is populated
	t.Run("Sync2",         testSync2)
	t.Run("Fleet",         testFleet)
	t.Run("DeployedFleet", testDeployedFleet)
	//t.Run("WarpAtoSwarm",  testWarpAtoSwarm)
	if len(secrets.ProfilesUserIds) > 0 {
		t.Run("Profiles",       testProfiles)
	}
	if len(secrets.AlliancesPublicInfoIds) > 0 {
		//t.Run("AlliancesProto", testAlliancesProto)
		t.Run("AlliancesJson",  testAlliancesJson)
	}
}

func testScopelyId(t *testing.T) {
	_, err := ScopelyId(secrets.ScopelyIdEmail, secrets.ScopelyIdPassword)
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
	t.Logf("Login succeeded as %s with session %s to instance %03d", r.InstanceAccount.Name, r.InstanceSessionId, r.InstanceAccount.InstanceIdCurrent)
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
Visit Data: info for %d of %.0f total (%.1f%%), visited %.0f (%.1f%%)`, sync.Starbase.UserId, len(sync.Resources), len(sync.VisitedSystems), n, pct1, visited, pct2)
}

func testProfiles(t *testing.T) {
	profiles, err := s.UserProfileProfiles(secrets.ProfilesUserIds)
	if err != nil {
		t.Fatalf("Profiles failed: %s", err)
	}
	t.Logf("Profiles succeeded: %v", profiles)
}

//func testAlliancesProto(t *testing.T) {
//	alliances, err := s.AlliancesProto(secrets.AlliancesPublicInfoIds)
//	if err != nil {
//		t.Fatalf("AlliancesProto failed: %s", err)
//	}
//	t.Logf("AlliancesProto succeeded: %v", alliances)
//}

func testAlliancesJson(t *testing.T) {
	alliances, err := s.AlliancesJson(secrets.AlliancesPublicInfoIds)
	if err != nil {
		t.Fatalf("AlliancesJson failed: %s", err)
	}
	t.Logf("AlliancesJson succeeded: %+v", alliances)
}

func testFleet(t *testing.T) {
	for _, dock := range AllDrydocks {
		f, err := s.Fleet(dock)
		if err != nil {
			t.Fatalf("Fleet failed: %s", err)
		}
		status := "docked"
		if _, ok := s.MyDeployedFleets[f.StringId]; ok {
			status = "deployed"
		}
		t.Logf("ID for %s is %d (%s)", dock, f.Id, status)
	}
	t.Logf("Fleet succeeded")
}

func testDeployedFleet(t *testing.T) {
	t.Logf("%6s  %19s  %10s  %4s  %4s  %6s  %9s  %4s  %5s  %7s  %7s  %7s", "", "Ship ID", "Hull ID", "HHP", "SHP", "Mining", "Destroyed", "Warp", "Speed", "Impulse", "Cargo", "PC")
	for _, dock := range AllDrydocks {
		f, err := s.Fleet(dock)
		if err != nil {
			t.Fatalf("DeployedFleet failed: %s", err)
		}
		df, ok := s.MyDeployedFleets[f.StringId]
		if !ok {
			continue
		}
		id := strconv.FormatUint(df.ShipIds[0], 10)
		hhp := (df.ShipHps[id] - df.ShipDmg[id]) / df.ShipHps[id] * 100
		shp := (df.ShipShieldHps[id] - df.ShipShieldDmg[id]) / df.ShipShieldHps[id] * 100
		t.Logf("%-6s  %19d  %10d  %3.0f%%  %3.0f%%  %-6t  %-9t  %4.0f  %5.2f  %7.0f  %7d  %7d", dock, df.ShipIds[0], df.HullIds[0], hhp, shp, df.IsMining, df.IsDestroyed, df.WarpDistance, df.WarpSpeed, df.ImpulseSpeed, df.FleetData.CargoMax, df.FleetData.SafeCargo)
	}
	t.Logf("DeployedFleet succeeded")
}

func testWarpAtoSwarm(t *testing.T) {
	Fyrsta := uint64(1549883102)
	f, err := s.Fleet(ShipA)
	x, y := RandomCoords()
	err = f.WarpTo(Fyrsta, x, y, false)
	if err != nil {
		t.Fatalf("WarpAtoSwarm failed: %s", err)
	}
	t.Logf("WarpAtoSwarm succeeded")
}

