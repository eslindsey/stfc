package stfc

import (
	"bytes"
	"encoding/json"
)

type CoursesSetFleetWarpCourseRequest struct {
	TargetActionId int         `json:"target_action_id"`
	FleetId        uint64      `json:"fleet_id"`
	TargetNode     NodeId      `json:"target_node"`
	TargetX        float32     `json:"target_x"`
	TargetY        float32     `json:"target_y"`
	TargetAction   int         `json:"target_action"`
	ClientWarpPath []NodeId    `json:"client_warp_path"`
	ClientWarpCost interface{} `json:"client_warp_cost"`
	IsInstantWarp  bool        `json:"is_instant_warp"`
}

func (s *Session) CoursesSetFleetWarpCourse(request *CoursesSetFleetWarpCourseRequest) ([]byte, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	_, err = s.Post("/courses/set_fleet_warp_course", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	// TODO: decode proto & JSON, return
	return nil, nil
}


