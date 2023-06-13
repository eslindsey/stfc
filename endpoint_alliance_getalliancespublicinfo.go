package stfc

import (
	"encoding/json"
	"bytes"

	"google.golang.org/protobuf/proto"
)

func (s *Session) AllianceGetAlliancesPublicInfo(allianceIds []uint64, reqType Message1Type) ([]byte, error) {
	b, err := json.Marshal(&AllianceRequest{AllianceIDs: allianceIds})
	if err != nil {
		return nil, err
	}
	body, err := s.Post("/alliance/get_alliances_public_info", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	var alliance AllianceEndpoint
	if err := proto.Unmarshal(body, &alliance); err != nil {
		return nil, err
	}
	for _, d := range alliance.Details {
		if Message1Type(d.Type) == reqType {
			return d.Details, nil
		}
	}
	return nil, ErrTypeNotFound
}

