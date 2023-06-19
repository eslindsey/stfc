package stfc

import (
	"bytes"
	"encoding/json"
)

type CoursesRecallFleetWarpRequest struct {
	ClientWarpCost Unknown     `json:"client_warp_cost"`
	ClientWarpPath []uint64    `json:"client_warp_path"`
	FleetId        uint64      `json:"fleet_id"`
	IsInstantWarp  bool        `json:"is_instant_warp"`
}

type CoursesRecallFleetWarpResponse struct {
	MyDeployedFleets MyDeployedFleets `json:"my_deployed_fleets"`
	StaticUpdate     Unknown          `json:"static_update"`
}

func (s *Session) CoursesRecallFleetWarp(request *CoursesRecallFleetWarpRequest) (*CoursesRecallFleetWarpResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, err := s.Post("/courses/recall_fleet_warp", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	// Decode proto & JSON, return
	var response CoursesRecallFleetWarpResponse
	if err := getMessage1JSON(body, &response); err != nil {
		return nil, err
	}
	s.MyDeployedFleets = response.MyDeployedFleets
	return &response, nil
}

