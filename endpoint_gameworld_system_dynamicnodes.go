package stfc

import (
	"bytes"
	"encoding/json"
)

type GameWorldSystemDynamicNodesResponse struct {
	ParentSystem          Unknown                   `json:"parent_system"`
	DeployedFleets        map[string]*DeployedFleet `json:"deployed_fleets"`
	PlayerContainer       Unknown                   `json:"player_container"`
	AllianceContainer     Unknown                   `json:"alliance_container"`
	MarauderQuickScanData Unknown                   `json:"marauder_quick_scan_data"`
	MiningSlots           Unknown                   `json:"mining_slots"`
	DockingPoints         Unknown                   `json:"docking_points"`
}

type MarauderQuickScanData struct {
	TargetId      string          `json:"target_id"`
	TargetFleetId uint64          `json:"target_fleet_id"`
	OffenseRating float32         `json:"offense_rating"`
	DefenseRating float32         `json:"defense_rating"`
	OfficerRating float32         `json:"officer_rating"`
	Resources     map[string]uint `json:"resources"`
	FleetType     uint            `json:"fleet_type"`
	ChestRarity   uint            `json:"chest_rarity"`
	ChestIdRefs   struct {
		ArtFileReference Unknown `json:"art_file_reference"`
		ArtId            uint    `json:"art_id"`
		LocaId           uint    `json:"loca_id"`
	} `json:"chest_id_refs"`
	CargoSpace     int       `json:"cargo_space"` // observed -1 value
	FactionId      uint64    `json:"faction_id"`
	FleetOfficers  []Unknown `json:"fleet_officers"`
	BridgeOfficers []Unknown `json:"bridge_officers"`
	IdRefs         struct {
		ArtFileReference Unknown `json:"art_file_reference"`
		ArtId            uint    `json:"art_id"`
		LocaId           uint    `json:"loca_id"`
	} `json:"id_refs"`
	ShipLevels                    map[string]uint `json:"ship_levels"`
	Strength                      float64         `json:"strength"`
	ParticipatingInArmadaAttackId uint            `json:"participating_in_armada_attack_id"`
	TargetedByArmadaAttackIds     []Unknown       `json:"targeted_by_armada_attack_ids"`
}

func (s *Session) GameWorldSystemDynamicNodes(systemId uint64) (*GameWorldSystemDynamicNodesResponse, error) {
	b, err := json.Marshal(map[string]uint64{"system_id": systemId})
	if err != nil {
		return nil, err
	}
	body, err := s.Post("/game_world/system/dynamic_nodes", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	// Decode protobuf & JSON as generic
	var response GameWorldSystemDynamicNodesResponse
	if err := getMessage1JSON(body, &response); err != nil {
		return nil, err
	}
	// TODO: decode proto & JSON, return
	return &response, nil
}
