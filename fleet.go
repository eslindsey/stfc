package stfc

import (
	"fmt"
)

type FleetId uint64

const (
	None  FleetId = 0
	ShipA         = 1936358958491318796
	ShipB         = 1936358958516484622
	ShipC         = 1936358958524873231
	ShipD         = 1936358958499707405
	ShipE         = 1936358958533261841
	ShipF         = 0 // unknown
	ShipG         = 0 // unknown
	ShipH         = 0 // unknown
)

func (f FleetId) String() string {
	switch f {
	case ShipA: return "Ship A"
	case ShipB: return "Ship B"
	case ShipC: return "Ship C"
	case ShipD: return "Ship D"
	case ShipE: return "Ship E"
	//case ShipF: return "Ship F"
	//case ShipG: return "Ship G"
	//case ShipH: return "Ship H"
	case None:  return "No Ship"
	default:    return fmt.Sprintf("FleetId(%d)", f)
	}
}

