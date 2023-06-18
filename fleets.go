package stfc

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type DrydockId uint

const (
	None  DrydockId = 0
	ShipA           = 17
	ShipB           = 18
	ShipC           = 19
	ShipD           = 20
	ShipE           = 21
	ShipF           = 22
	ShipG           = 44
	ShipH           = 45
)

func (d DrydockId) String() string {
	switch d {
	case ShipA: return "Ship A"
	case ShipB: return "Ship B"
	case ShipC: return "Ship C"
	case ShipD: return "Ship D"
	case ShipE: return "Ship E"
	//case ShipF: return "Ship F"
	//case ShipG: return "Ship G"
	//case ShipH: return "Ship H"
	case None:  return "No Ship"
	default:    return fmt.Sprintf("DrydockId(%d)", d)
	}
}

var (
	AllDrydocks = []DrydockId{}

	ErrFleetNotFound = errors.New("fleet not found")
)

type FleetRaw struct {
	ShipIds             []uint             `json:"ship_ids"`
	Name                string             `json:"name"`
	DrydockId           DrydockId          `json:"drydock_id"`
	LastRecall          string             `json:"last_recall"` // TODO: time.Time
	Officers            []uint             `json:"officers"`
	Stats               map[string]float32 `json:"stats"`
	RepairTime          uint               `json:"repair_time"`
	RepairCost          Unknown            `json:"repair_cost"`
	PrecalculatedRepair bool               `json:"precalculated_repair"`
}

type Fleet struct {
	*FleetRaw
	Id       uint64 `json:"-"`
	StringId string `json:"-"`
}

func (s *Session) Fleet(id DrydockId) (*Fleet, error) {
	if s.Sync2Response == nil {
		return nil, ErrNotSynced
	}
	for k, v := range s.Sync2Response.Fleets {
		if v.DrydockId != id {
			continue
		}
		var err error
		ret := &Fleet{StringId: k, FleetRaw: v}
		ret.Id, err = strconv.ParseUint(k, 10, 64)
		if err != nil {
			return nil, err
		}
		return ret, nil
	}
	return nil, ErrFleetNotFound
}

func (f *Fleet) WarpTo(target uint64, x, y float64, instant bool) error {
	if s.Sync2Response == nil {
		return ErrNotSynced
	}
	g, err := s.Galaxy()
	if err != nil {
		return err
	}
	// See if fleet is deployed, otherwise traveling from home
	from := s.Sync2Response.Starbase.Location.System
	if f, ok := s.MyDeployedFleets[f.StringId]; ok {
		from = f.NodeAddress.System
	}
	// Calculate shortest path
	path, _ := g.Shortest(from, target)
	_, err = s.CoursesSetFleetWarpCourse(&CoursesSetFleetWarpCourseRequest{
		TargetActionId: 0,
		FleetId:        f.Id,
		TargetNode:     target,
		TargetX:        x,
		TargetY:        y,
		TargetAction:   -1,
		ClientWarpPath: path,
		IsInstantWarp:  instant,
	})
	return err
}

func (s *Session) populateDrydocks() {
	for _, v := range s.Sync2Response.Fleets {
		AllDrydocks = append(AllDrydocks, v.DrydockId)
	}
	sort.Slice(AllDrydocks, func(i, j int) bool {
		return AllDrydocks[i] < AllDrydocks[j]
	})
}

