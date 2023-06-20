package stfc

type StaticShip struct {
	Id                 uint         `json:"id"`
	ArtId              uint         `json:"art_id"`
	LocaId             uint         `json:"loca_id"`
	MaxTier            uint         `json:"max_tier"`
	Rarity             uint         `json:"rarity"`
	Grade              uint         `json:"grade"`
	ScrapLevel         int          `json:"scrap_level"`
	BuildTimeInSeconds uint         `json:"build_time_in_seconds"`
	Faction            uint         `json:"faction"`
	BlueprintsRequired uint         `json:"blueprints_required"`
	HullType           uint         `json:"hull_type"`
	MaxLevel           uint         `json:"max_level"`
	BuildCost          ResourceList `json:"build_cost"`
	RepairCost         ResourceList `json:"repair_cost"`
	RepairTime         uint         `json:"repair_time"`
	BuildRequirements  []struct {
		RequirementType  string  `json:"requirement_type"`
		RequirementId    uint    `json:"requirement_id"`
		RequirementLevel uint    `json:"requirement_level"`
		PowerGain        Unknown `json:"power_gain"`
	} `json:"build_requirements"`
	CrewSlots []struct {
		Slots       uint `json:"slots"`
		UnlockLevel uint `json:"unlock_level"`
	} `json:"crew_slots"`
	Levels []struct {
		Level  uint    `json:"level"`
		Xp     int     `json:"xp"`
		Shield float64 `json:"shield"`
		Health float64 `json:"health"`
	} `json:"levels"`
	Ability struct {
		Id                uint      `json:"id"`
		ValueIsPercentage bool      `json:"value_is_percentage"`
		Values            []Unknown `json:"values"`
		ArtId             uint      `json:"art_id"`
		ShowPercentage    bool      `json:"show_percentage"`
		ValueType         uint      `json:"value_type"`
		Flag              uint      `json:"flag"`
	} `json:"ability"`
	BaseScrap ResourceList `json:"base_scrap"`
}

type ResourceList []struct {
	ResourceId uint `json:"resource_id"`
	Amount     uint `json:"amount"`
}
