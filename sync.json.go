package stfc

type UnknownObject struct{}
type UnknownType   interface{}

type SyncJSON struct {
	MyDeployedFleets UnknownObject `json:"my_deployed_fleets"`
	StaticUpdate     UnknownType   `json:"static_update"`
	Starbase         struct {
		UserID   string `json:"user_id"`
		Location struct {
			Galaxy uint `json:"galaxy"`
			System uint `json:"system"`
			Planet uint `json:"planet"`
		} `json:"location"`
		Destination UnknownType `json:"destination"`
		PeaceShield struct {
			ID                    uint   `json:"id"`
			Target                uint   `json:"target"`
			TriggeredOn           string `json:"triggered_on"` // TODO: time.Time
			ExpiryTime            string `json:"expiry_time"`  // TODO: time.Time
			PeaceShieldResourceID int    `json:"peace_shield_resource_id"`
		} `json:"peace_shield"`
		State             uint        `json:"state"`
		LastRelocation    *string     `json:"last_relocation"` // TODO: nil time.Time
		BattleID          UnknownType `json:"battle_id"`
		CeasefireBrokenAt *string     `json:"ceasefire_broken_at"` // TODO: nil time.Time
		Coordinates       struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"coordinates"`
	} `json:"starbase"`
	StarbaseModules map[string]struct {
		ID    uint `json:"id"`
		Level uint `json:"level"`
	} `json:"starbase_modules"`
	Ships map[string]struct {
		ID              uint    `json:"id"`
		HullID          uint    `json:"hull_id"`
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
		ID             uint   `json:"id"`
		HullID         uint   `json:"hull_id"`
		Damage         uint   `json:"damage"`
		ShieldDamage   uint   `json:"shield_damage"`
		LastUpdateTime string `json:"last_update_time"` // TODO: time.Time
		Components     []int  `json:"components"`
	} `json:"defenses"`
	VisitedSystems []bool `json:"visited_systems"`
	MyShieldState  struct {
		ID                    uint   `json:"id"`
		Target                uint   `json:"target"`
		TriggeredOn           string `json:"triggered_on"` // TODO: time.Time
		ExpiryTime            string `json:"expiry_time"`  // TODO: time.Time
		PeaceShieldResourceID int    `json:"peace_shield_resource_id"`
	} `json:"my_shield_state"`
	UserHistory     map[string]bool `json:"user_history"`
	FactionStanding map[string]struct {
		Standing     uint `json:"standing"`
		IsDiscovered bool `json:"is_discovered"`
	} `json:"faction_standing"`
	MySkillData         UnknownObject `json:"my_skill_data"`
	BattleResultHeaders []UnknownType `json:"battle_result_headers"`
	Fleets              map[string]struct {
		ShipIDs             []uint             `json:"ship_ids"`
		Name                string             `json:"name"`
		DrydockID           uint               `json:"drydock_id"`
		LastRecall          string             `json:"last_recall"` // TODO: time.Time
		Officers            []uint             `json:"officers"`
		Stats               map[string]float32 `json:"stats"`
		RepairTime          uint               `json:"repair_time"`
		RepairCost          UnknownObject      `json:"repair_cost"`
		PrecalculatedRepair bool               `json:"precalculated_repair"`
	} `json:"fleets"`
	FulfilledConnectionRequirements []bool `json:"fulfilled_connection_requirements"`
}
