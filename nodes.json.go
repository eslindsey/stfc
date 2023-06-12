package stfc

import (
	"encoding/json"
	"errors"
)

var (
	ErrLengthMismatch = errors.New("length mismatch")
)

type GalaxyOptimisedMiningSetup uint64

type GalaxyOptimisedRaw struct {
	Data struct {
		// Count: 1367
		NodeIds              []uint64                       `json:"node_ids"`
		XCoords              []int                          `json:"x_coords"`
		YCoords              []int                          `json:"y_coords"`
		ConnectionsCount     []uint32                       `json:"connections_count"`
		ConnectionsOffset    []uint32                       `json:"connections_offset"`
		Levels               []uint32                       `json:"levels"`
		Priorities           []uint32                       `json:"priorities"`
		Factions             []int64                        `json:"factions"`
		FactionInfluences    []uint32                       `json:"faction_influences"`
		TransIds             []uint32                       `json:"trans_ids"`
		AssetIds             []int32                        `json:"asset_ids"`
		MiningSetups         [][]GalaxyOptimisedMiningSetup `json:"mining_setups"`
		MarauderSpawnRuleIds [][]uint64                     `json:"marauder_spawn_rule_ids"`
		IsDeepSpace          []bool                         `json:"is_deep_space"`

		// Count: 3292
		SourceIds      []uint64 `json:"source_ids"`
		DestIds        []uint64 `json:"dest_ids"`
		Distances      []uint32 `json:"distances"`
		UnlockReqCount []uint32 `json:"unlock_req_count"`
		UnlockOffset   []uint32 `json:"unlock_offset"`

		// Count: 231
		UnlockReqTypes      []uint32 `json:"unlock_req_types"`
		UnlockReqSources    []uint64 `json:"unlock_req_sources"`
		UnlockReqQuantities []uint32 `json:"unlock_req_quantities"`

		// Count: 34
		SuperHighwaysIndices []uint32 `json:"super_highways_indices"`
	} `json:"galaxy_optimised"`
}

type GalaxyOptimised struct {
	Nodes                []*GalaxyOptimisedNode
	Paths                []*GalaxyOptimisedPath
	UnlockData           []*GalaxyOptimisedUnlockData
	SuperHighwaysIndices []uint32

	NodesById map[uint64]*GalaxyOptimisedNode
}

type GalaxyOptimisedNode struct {
	Id                  uint64
	XCoord              int
	YCoord              int
	ConnectionCount     uint32
	ConnectionOffset    uint32
	Level               uint32
	Priority            uint32
	Faction             int64
	FactionInfluence    uint32
	TransId             uint32
	AssetId             int32
	MiningSetups        []GalaxyOptimisedMiningSetup
	MarauderSpawnRuleIds []uint64
	IsDeepSpace         bool
}

type GalaxyOptimisedPath struct {
	SourceId       uint64
	DestId         uint64
	Distance       uint32
	UnlockReqCount uint32
	UnlockOffset   uint32
}

type GalaxyOptimisedUnlockData struct {
	Type     uint32
	Source   uint64
	Quantity uint32
}

func (g *GalaxyOptimised) UnmarshalJSON(b []byte) error {
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
	g.Nodes = make([]*GalaxyOptimisedNode, n)
	for i := 0; i < n; i++ {
		g.Nodes[i] = &GalaxyOptimisedNode{
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
	g.Paths = make([]*GalaxyOptimisedPath, n)
	for i := 0; i < n; i++ {
		g.Paths[i] = &GalaxyOptimisedPath{
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
	g.UnlockData = make([]*GalaxyOptimisedUnlockData, n)
	for i := 0; i < n; i++ {
		g.UnlockData[i] = &GalaxyOptimisedUnlockData{
			Type: d.UnlockReqTypes[i],
			Source: d.UnlockReqSources[i],
			Quantity: d.UnlockReqQuantities[i],
		}
	}
	g.SuperHighwaysIndices = d.SuperHighwaysIndices
	return nil
}

