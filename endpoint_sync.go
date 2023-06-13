package stfc

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
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
	var dest *SyncJSON
	if n == 2 {
		dest = s.Sync2Response
	}
	dest = &SyncJSON{}
	if err := json.Unmarshal([]byte(sync.Payload.Json), dest); err != nil {
		return nil, err
	}
	if n == 2 {
		s.populateVisited()
	}
	return dest, nil
}

