package stfc

import (
	"encoding/json"
	"errors"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/path"
)

// SystemRadius is the radius of the game playfield. When setting system
// coordinates, the x, y values should fall within the circle described by
// this radius.
const SystemRadius = 1100.0

var (
	ErrLengthMismatch = errors.New("length mismatch")
)

type Galaxy struct {
	Nodes                []*GalaxyNode
	Paths                []*GalaxyPath
	UnlockData           []*GalaxyUnlockData
	SuperHighwaysIndices []uint32

	byId  map[uint64]*GalaxyNode
	graph *simple.WeightedDirectedGraph
}

type GalaxyNode struct {
	Id                   uint64
	XCoord               int
	YCoord               int
	ConnectionCount      uint32
	ConnectionOffset     uint32
	Level                uint32
	Priority             uint32
	Faction              int64
	FactionInfluence     uint32
	TransId              uint32
	AssetId              int32
	MiningSetups         []uint64  // TODO: Convert to enum
	MarauderSpawnRuleIds []uint64
	IsDeepSpace          bool

	// Not part of galaxy_optimised endpoint
	IsVisited            bool
}

type GalaxyPath struct {
	SourceId       uint64
	DestId         uint64
	Distance       uint32
	UnlockReqCount uint32
	UnlockOffset   uint32
}

type GalaxyUnlockData struct {
	Type     uint32
	Source   uint64
	Quantity uint32
}

func (g *Galaxy) UnmarshalJSON(b []byte) error {
	var r GalaxyOptimisedRaw
	err := json.Unmarshal(b, &r)
	d := &r.Data
	if err != nil {
		return err
	}
	n := len(d.NodeIds)
	if len(d.XCoords) != n || len(d.YCoords) != n || len(d.ConnectionsCount) != n || len(d.ConnectionsOffset) != n || len(d.Levels) != n || len(d.Priorities) != n || len(d.Factions) != n || len(d.FactionInfluences) != n || len(d.TransIds) != n || len(d.AssetIds) != n || len(d.MiningSetups) != n || len(d.MarauderSpawnRuleIds) != n || len(d.IsDeepSpace) != n {
		return ErrLengthMismatch
	}
	g.Nodes = make([]*GalaxyNode, n)
	for i := 0; i < n; i++ {
		g.Nodes[i] = &GalaxyNode{
			Id: d.NodeIds[i],
			XCoord: d.XCoords[i],
			YCoord: d.YCoords[i],
			ConnectionCount: d.ConnectionsCount[i],
			ConnectionOffset: d.ConnectionsOffset[i],
			Level: d.Levels[i],
			Priority: d.Priorities[i],
			Faction: d.Factions[i],
			FactionInfluence: d.FactionInfluences[i],
			TransId: d.TransIds[i],
			AssetId: d.AssetIds[i],
			MiningSetups: d.MiningSetups[i],
			MarauderSpawnRuleIds: d.MarauderSpawnRuleIds[i],
			IsDeepSpace: d.IsDeepSpace[i],
		}
	}
	n = len(d.SourceIds)
	if len(d.DestIds) != n || len(d.Distances) != n || len(d.UnlockReqCount) != n || len(d.UnlockOffset) != n {
		return ErrLengthMismatch
	}
	g.Paths = make([]*GalaxyPath, n)
	for i := 0; i < n; i++ {
		g.Paths[i] = &GalaxyPath{
			SourceId: d.SourceIds[i],
			DestId: d.DestIds[i],
			Distance: d.Distances[i],
			UnlockReqCount: d.UnlockReqCount[i],
			UnlockOffset: d.UnlockOffset[i],
		}
	}
	n = len(d.UnlockReqTypes)
	if len(d.UnlockReqSources) != n || len(d.UnlockReqQuantities) != n {
		return ErrLengthMismatch
	}
	g.UnlockData = make([]*GalaxyUnlockData, n)
	for i := 0; i < n; i++ {
		g.UnlockData[i] = &GalaxyUnlockData{
			Type: d.UnlockReqTypes[i],
			Source: d.UnlockReqSources[i],
			Quantity: d.UnlockReqQuantities[i],
		}
	}
	g.SuperHighwaysIndices = d.SuperHighwaysIndices
	return nil
}

func (g *Galaxy) Get(s uint64) (*GalaxyNode, bool) {
	// Lazy populate
	if g.byId == nil {
		for i, n := range g.Nodes {
			g.byId[n.Id] = g.Nodes[i]
		}
	}
	val, ok := g.byId[s]
	return val, ok
}

type ShortestOptions struct {
	MaxWarp *int
	MinWarp *int
}

func (g *Galaxy) Shortest(from, to uint64, options ...*ShortestOptions) ([]uint64, uint64) {
	//opts := ShortestOptionsArray(options).Gather()
	// Create a graph
	graph := simple.NewWeightedDirectedGraph(0, math.Inf(1))
	for _, path := range g.Paths {
		// TODO: Check warp range options
		graph.SetWeightedEdge(simple.WeightedEdge{
			F: simple.Node(path.SourceId),
			T: simple.Node(path.DestId),
			W: float64(path.Distance),
		})
	}
	// Calculate shortest path
	shortest := path.DijkstraFrom(simple.Node(from), graph)
	path, distance := shortest.To(int64(to))
	ret := make([]uint64, len(path))
	for i, _ := range path {
		ret[i] = uint64(path[i].ID())
	}
	return ret, uint64(distance)
}

type ShortestOptionsArray []*ShortestOptions

func (soa ShortestOptionsArray) Gather() *ShortestOptions {
	// Later options override earlier options
	if len(soa) < 1 {
		soa = append(soa, &ShortestOptions{})
	}
	for i := 1; i < len(soa); i++ {
		if soa[i].MaxWarp != nil {
			soa[0].MaxWarp = soa[i].MaxWarp
		}
		if soa[i].MinWarp != nil {
			soa[0].MinWarp = soa[i].MinWarp
		}
	}
	return soa[0]
}

/*
 * UTILITY FUNCTIONS
 */

// TranslateCoords maths polar coordinates (degrees of angle, and radius) into
// rectangular coordinates (x, y). This function does NOT guarantee that the
// values you pass are within the game's playfield. See [SystemRadius].
func TranslateCoords(degrees, radius float64) (float64, float64) {
	theta := math.Pi / 180.0
	return radius * math.Cos(theta), radius * math.Sin(theta)
}

// RandomCoords returns a set of x, y coordinates that is guaranteed to be
// within the game playfield area.
func RandomCoords() (x, y float64) {
	r := rand.Float64() * SystemRadius
	d := rand.Float64() * 360.0
	return TranslateCoords(d, r)
}

