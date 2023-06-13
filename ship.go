package stfc

import (
	"errors"
)

type Ship struct {
	Session  *Session
	FleetId  uint64
	Location NodeId
}

var (
	ErrNoPathFound = errors.New("no path found")
)

func (s *Ship) WarpTo(target NodeId, x, y float32, instant bool) error {
	g, err := s.Session.Galaxy()
	if err != nil {
		return err
	}
	path, _ := g.Shortest(s.Location, target)
	// TODO: Eliminate paths based on warp range, discovered systems
	_, err = s.Session.CoursesSetFleetWarpCourse(&CoursesSetFleetWarpCourseRequest{
		TargetActionId: 0,
		FleetId:        s.FleetId,
		TargetNode:     target,
		TargetX:        x,
		TargetY:        y,
		TargetAction:   -1,
		ClientWarpPath: path,
		IsInstantWarp:  instant,
	})
	return err
}

