package stfc

import (
	"bytes"
	"encoding/json"
)

type CoursesSetFleetWarpCourseRequest struct {
	TargetActionId int         `json:"target_action_id"`
	FleetId        uint64      `json:"fleet_id"`
	TargetNode     uint64      `json:"target_node"`
	TargetX        float64     `json:"target_x"`  // should be within playfield; see [TranslateCoords]
	TargetY        float64     `json:"target_y"`  // should be within playfield; see [TranslateCoords]
	TargetAction   int         `json:"target_action"`
	ClientWarpPath []uint64    `json:"client_warp_path"`
	ClientWarpCost interface{} `json:"client_warp_cost"`
	IsInstantWarp  bool        `json:"is_instant_warp"`
}

type CoursesSetFleetWarpCourseResponse struct {
	MyDeployedFleets MyDeployedFleets `json:"my_deployed_fleets"`
	StaticUpdate     Unknown          `json:"static_update"`
	VisitedSystems   []bool           `json:"visited_systems"`
}

// CoursesSetFleetWarpCourse initiates a warp course. The destination X and Y
// should be within the playfield. See [TranslateCoords] and [RandomCoords].
func (s *Session) CoursesSetFleetWarpCourse(request *CoursesSetFleetWarpCourseRequest) (*CoursesSetFleetWarpCourseResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, err := s.Post("/courses/set_fleet_warp_course", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	// Decode proto & JSON, return
	var response CoursesSetFleetWarpCourseResponse
	if err := getMessage1JSON(body, &response); err != nil {
		return nil, err
	}
	s.MyDeployedFleets = response.MyDeployedFleets
	// TODO: Ensure requested ship is actually deployed and warping?
	return &response, nil
}

