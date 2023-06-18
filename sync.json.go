package stfc

type SyncJSON struct {
	MyDeployedFleets MyDeployedFleets `json:"my_deployed_fleets"`
	StaticUpdate     Unknown      `json:"static_update"`
	Starbase         struct {
		UserId   string `json:"user_id"`
		Location struct {
			Galaxy uint64 `json:"galaxy"`
			System uint64 `json:"system"`
			Planet uint64 `json:"planet"`
		} `json:"location"`
		Destination Unknown `json:"destination"`
		PeaceShield struct {
			Id                    uint   `json:"id"`
			Target                uint   `json:"target"`
			TriggeredOn           string `json:"triggered_on"` // TODO: time.Time
			ExpiryTime            string `json:"expiry_time"`  // TODO: time.Time
			PeaceShieldResourceId int    `json:"peace_shield_resource_id"`
		} `json:"peace_shield"`
		State             uint        `json:"state"`
		LastRelocation    *string     `json:"last_relocation"` // TODO: nil time.Time
		BattleId          Unknown `json:"battle_id"`
		CeasefireBrokenAt *string     `json:"ceasefire_broken_at"` // TODO: nil time.Time
		Coordinates       struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"coordinates"`
	} `json:"starbase"`
	StarbaseModules map[string]struct {
		Id    uint `json:"id"`
		Level uint `json:"level"`
	} `json:"starbase_modules"`
	Ships map[string]struct {
		Id              uint    `json:"id"`
		HullId          uint    `json:"hull_id"`
		MaxHP           float32 `json:"max_hp"`
		MaxShieldHP     float32 `json:"max_shield_hp"`
		Damage          float32 `json:"damage"`
		ShieldDamage    float32 `json:"shield_damage"`
		Level           uint    `json:"level"`
		LevelPercentage float32 `json:"level_percentage"`
		Tier            uint    `json:"tier"`
		LastUpdateTime  string  `json:"last_update_time"` // TODO: time.Time
		Cosmetics       struct {
		} `json:"cosmetics"`
		Components  []int `json:"components"`
		IsDestroyed bool  `json:"is_destroyed"`
	} `json:"ships"`
	Resources map[string]struct {
		CurrentAmount int `json:"current_amount"`
		UpperCapacity int `json:"upper_capacity"`
		LowerCapacity int `json:"lower_capacity"`
	} `json:"resources"`
	ResourceProducers map[string]struct {
		LastSnapshot  string  `json:"last_snapshot"` // TODO: time.Time
		LastHarvest   string  `json:"last_harvest"`  // TODO: time.Time
		CurrentAmount float32 `json:"current_amount"`
		IsActive      bool    `json:"is_active"`
		TotalAmount   float32 `json:"total_amount"`
	} `json:"resource_producers"`
	Defenses map[string]struct {
		Id             uint   `json:"id"`
		HullId         uint   `json:"hull_id"`
		Damage         uint   `json:"damage"`
		ShieldDamage   uint   `json:"shield_damage"`
		LastUpdateTime string `json:"last_update_time"` // TODO: time.Time
		Components     []int  `json:"components"`
	} `json:"defenses"`
	VisitedSystems []bool `json:"visited_systems"`
	MyShieldState  struct {
		Id                    uint   `json:"id"`
		Target                uint   `json:"target"`
		TriggeredOn           string `json:"triggered_on"` // TODO: time.Time
		ExpiryTime            string `json:"expiry_time"`  // TODO: time.Time
		PeaceShieldResourceId int    `json:"peace_shield_resource_id"`
	} `json:"my_shield_state"`
	UserHistory     map[string]bool `json:"user_history"`
	FactionStanding map[string]struct {
		Standing     uint `json:"standing"`
		IsDiscovered bool `json:"is_discovered"`
	} `json:"faction_standing"`
	MySkillData         Unknown `json:"my_skill_data"`
	BattleResultHeaders []Unknown `json:"battle_result_headers"`
	Fleets              map[string]*FleetRaw  `json:"fleets"`
	FulfilledConnectionRequirements []bool `json:"fulfilled_connection_requirements"`
}

