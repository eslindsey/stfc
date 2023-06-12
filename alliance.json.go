package stfc

import (
	"fmt"
)

type Rank uint

const (
	RankUnknown   Rank = 0
	RankAdmiral        = 4186399962
	RankCommodore      = 1869972635
	RankPremier        = 3423897132
	RankOperative      = 4071078643
	RankAgent          = 1967264300
)

func (r Rank) String() string {
	switch r {
	case RankAdmiral:
		return "Admiral"
	case RankCommodore:
		return "Commodore"
	case RankPremier:
		return "Premier"
	case RankOperative:
		return "Operative"
	case RankAgent:
		return "Agent"
	default:
		return fmt.Sprintf("Rank(%d)", r)
	}
}

type Relationship uint

const (
	RelationshipUnknown Relationship = 0
)

func (r Relationship) String() string {
	switch r {
	default:
		return fmt.Sprintf("Relationship(%d)", r)
	}
}

type AllianceJson struct {
	Contributions map[string]uint `json:"alliance_contributions"`
	Alliance      struct {
		*AlliancePublicInfoJson
		Announcement string                   `json:"announcement"`
		HelpedJobs   map[string][]interface{} `json:"helped_jobs"` // map value is probably [requestor_id string, unknown uint, time string, [helper_id string, ...]]
	} `json:"alliance"`
	Members map[string]struct {
		Rank      Rank   `json:"rank"`
		CreatedAt string `json:"created_at"`
	} `json:"alliance_members"`
	Notifications []struct {
		UUID             string  `json:"UUID"`
		CreatedTimestamp float32 `json:"created_timestamp"`
		ExpiryTimestamp  float32 `json:"expiry_timestamp"`
		ProducerType     uint    `json:"producer_type"`
		Params           struct {
			AllianceID               string `json:"alliance_id"`
			AllianceNotificationType uint   `json:"alliance_notification_type"`
			RequestingUserID         string `json:"requesting_user_id"`
			JobType                  uint   `json:"job_type"`
		} `json:"params"`
	} `json:"alliance_notifications"`
	Diplomacy      map[string]Relationship `json:"alliance_diplomacy_relationships"`
	MemberActivity struct{}                `json:"alliance_member_activity"`
}

type AlliancesPublicInfoWrapperJson map[string]*AlliancesPublicInfoJson

type AlliancesPublicInfoJson map[string]*AlliancePublicInfoJson

type AlliancePublicInfoJson struct {
	ID            uint64 `json:"id"`
	Name          string `json:"name"`
	Slogan        string `json:"slogan"`
	Tag           string `json:"tag"`
	Public        bool   `json:"public"`
	Emblem        uint32 `json:"emblem"`
	Level         uint32 `json:"level"`
	GameworldID   uint32 `json:"gameworld_id"`
	MemberCount   uint32 `json:"member_count"`
	MilitaryMight uint32 `json:"military_might"`
}

