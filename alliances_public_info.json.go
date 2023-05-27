package stfc

type AlliancesPublicInfoWrapperJson map[string]*AlliancesPublicInfoJson

type AlliancesPublicInfoJson map[string]struct {
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

