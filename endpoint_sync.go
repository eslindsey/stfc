package stfc

import (
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
)

var (
	ErrNotSynced = errors.New("sync2 has not been performed yet")
)

func (s *Session) Sync(n int) (*SyncJSON, error) {
	body, err := s.Post("/sync", []Header{{"X-Prime-Sync", fmt.Sprintf("%d", n)}}, nil)
	if err != nil {
		return nil, err
	}
	var sync Sync
	if err := proto.Unmarshal(body, &sync); err != nil {
		return nil, err
	}
	if sync.Payload == nil {
		return nil, ErrSyncMissingPayload
	}
	dest := &SyncJSON{}
	if n == 2 {
		if s.Sync2Response == nil {
			s.Sync2Response = &SyncJSON{}
		}
		dest = s.Sync2Response
	}
	if err := json.Unmarshal([]byte(sync.Payload.Json), dest); err != nil {
		return nil, err
	}
	if dest.MyDeployedFleets != nil {
		s.MyDeployedFleets = dest.MyDeployedFleets
	}
	if n == 2 {
		s.populateVisited()
		s.populateDrydocks()
	}
	return dest, nil
}

