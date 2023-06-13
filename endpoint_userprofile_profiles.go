package stfc

import (
	"bytes"
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

func (s *Session) UserProfileProfiles(userIds []string) ([]*Profile, error) {
	b, err := json.Marshal(map[string][]string{"user_ids": userIds})
	if err != nil {
		return nil, err
	}
	body, err := s.Post("/user_profile/profiles", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	var profiles Profiles
	if err := proto.Unmarshal(body, &profiles); err != nil {
		return nil, err
	}
	return profiles.Payload.Payload2.Profile, nil
}

