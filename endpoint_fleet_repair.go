package stfc

import (
	"bytes"
	"encoding/json"
)

type FleetRepairRequest struct {
	FleetId    uint64  `json:"fleet_id"`
}

func (s *Session) FleetRepair(request *FleetRepairRequest) ([]byte, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	_, err = s.Post("/fleet/repair", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	// TODO: decode proto & JSON, return
	return nil, nil
/*
*/
}

