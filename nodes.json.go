package stfc

type NodeId uint64

type GalaxyOptimisedRaw struct {
	Data struct {
		// Count: 1367
		NodeIds              []NodeId   `json:"node_ids"`
		XCoords              []int      `json:"x_coords"`
		YCoords              []int      `json:"y_coords"`
		ConnectionsCount     []uint32   `json:"connections_count"`
		ConnectionsOffset    []uint32   `json:"connections_offset"`
		Levels               []uint32   `json:"levels"`
		Priorities           []uint32   `json:"priorities"`
		Factions             []int64    `json:"factions"`
		FactionInfluences    []uint32   `json:"faction_influences"`
		TransIds             []uint32   `json:"trans_ids"`
		AssetIds             []int32    `json:"asset_ids"`
		MiningSetups         [][]uint64 `json:"mining_setups"`
		MarauderSpawnRuleIds [][]uint64 `json:"marauder_spawn_rule_ids"`
		IsDeepSpace          []bool     `json:"is_deep_space"`

		// Count: 3292
		SourceIds      []NodeId `json:"source_ids"`
		DestIds        []NodeId `json:"dest_ids"`
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

