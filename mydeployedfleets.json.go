package stfc

type MyDeployedFleets map[string]*DeployedFleet

type DeployedFleet struct {
	Uid                                  string                        `json:"uid"`
	State                                int                           `json:"state"`
	CourseId                             *uint64                       `json:"course_id"`
	NodeId                               uint64                        `json:"node_id"`
	NodeAddress                          NodeAddress                   `json:"node_address"`
	Type                                 uint                          `json:"type"`
	FleetId                              uint64                        `json:"fleet_id"`
	PursuedByNum                         int                           `json:"pursued_by_num"`
	PursuitTargetId                      Unknown                       `json:"pursuit_target_id"` // can be null
	IsActive                             bool                          `json:"is_active"`
	ShipIds                              []uint64                      `json:"ship_ids"`
	HullIds                              []uint                        `json:"hull_ids"`
	WarpTime                             *ScopelyTime                  `json:"warp_time"`
	WarpData                             []Unknown                     `json:"warp_data"` // looks like 0=warp_path, 1=target_x, 2=target_y, 3=warp_cost, 4=target_action, 5=is_instant_warp ?
	LatestCourseVectorX                  *float64                      `json:"latest_course_vector_x"`
	LatestCourseVectorY                  *float64                      `json:"latest_course_vector_y"`
	LastUpdateTime                       ScopelyTime                   `json:"last_update_time"`
	Attributes                           map[string]float32            `json:"attributes"` // keys range from "-7" to "-14"
	Stats                                map[string]float32            `json:"stats"`
	ShipAttributes                       map[string]map[string]float32 `json:"ship_attributes"`
	ShipStats                            map[string]map[string]float32 `json:"ship_stats"`
	InternalStats                        map[string]float32            `json:"internal_stats"`
	ActiveBuffs                          []*Buff                       `json:"active_buffs"`
	ShipTiers                            map[string]int                `json:"ship_tiers"`
	ShipLevels                           map[string]int                `json:"ship_levels"`
	ShipLevelPercentages                 map[string]float32            `json:"ship_level_percentages"`
	ShipComponents                       map[string][]int              `json:"ship_components"`   // found a -1 value
	StatusEffects                        Unknown                       `json:"status_effects"`
	ShipCosmetics                        Unknown                       `json:"ship_cosmetics"`
	BattleOpponentFleetId                Unknown                       `json:"battle_opponent_fleet_id"`
	BattleWon                            Unknown                       `json:"battle_won"`
	BattleStartTime                      ScopelyTime                   `json:"battle_start_time"`
	IsNoEffect                           bool                          `json:"is_no_effect"`
	IsCloaked                            bool                          `json:"is_cloaked"`
	IsSupported                          bool                          `json:"is_supported"`
	IsDebuffed                           bool                          `json:"is_debuffed"`
	IsWarShieldActivated                 bool                          `json:"is_war_shield_activated"`
	IsArmadaSupported                    bool                          `json:"is_armada_supported"`
	IsWeaponDamageActivated              bool                          `json:"is_weapon_damage_activated"`
	IsWeaponPenetrationActivated         bool                          `json:"is_weapon_penetration_activated"`
	IsWeaponShotsActivated               bool                          `json:"is_weapon_shots_activated"`
	IsCriticalDamageActivated            bool                          `json:"is_critical_damage_activated"`
	IsDetected                           bool                          `json:"is_detected"`
	IsSystemWideBuffed                   bool                          `json:"is_system_wide_buffed"`
	IsSystemWideSupremeBuffed            bool                          `json:"is_system_wide_supreme_buffed"`
	FleetGrade                           float32                       `json:"fleet_grade"`
	DefenseRating                        float32                       `json:"defense_rating"`
	OffenseRating                        float32                       `json:"offense_rating"`
	HealthRating                         float32                       `json:"health_rating"`
	OfficerRating                        float32                       `json:"officer_rating"`
	DeflectorRating                      float32                       `json:"deflector_rating"`
	SensorRating                         float32                       `json:"sensor_rating"`
	ShipHps                              map[string]float32            `json:"ship_hps"`
	ShipShieldHps                        map[string]float32            `json:"ship_shield_hps"`
	ShipDmg                              map[string]float32            `json:"ship_dmg"`
	ShipShieldDmg                        map[string]float32            `json:"ship_shield_dmg"`
	ImpulseSpeed                         float32                       `json:"impulse_speed"`
	WarpDistance                         float32                       `json:"warp_distance"`
	WarpSpeed                            float32                       `json:"warp_speed"`
	ShipShieldTotalRegenerationDurations map[string]float32            `json:"ship_shield_total_regeneration_durations"`
	CurrentCoords                        struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"current_coords"`
	IsMining      bool    `json:"is_mining"`
	CurrentCourse Unknown `json:"current_course"`
	IsDestroyed   bool    `json:"is_destroyed"`
	FleetData     struct {
		Cargo struct {
			Resources Unknown `json:"resources"`
		} `json:"cargo"`
		CargoMax         uint        `json:"cargo_max"`
		SafeCargo        uint        `json:"safe_cargo"`
		DockPointId      int         `json:"dock_point_id"`
		CrewData         []*CrewData `json:"crew_data"`
		MiningLevel      int         `json:"mining_level"`
		LastBattleId     Unknown     `json:"last_battle_id"`
		LastChestRollKey Unknown     `json:"last_chest_roll_key"`
		ArmadaAttackId   int         `json:"armada_attack_id"`
	} `json:"fleet_data"`
	DockPointData            Unknown   `json:"dock_point_data"`
	LastHailingFrequencySent Unknown   `json:"last_hailing_frequency_sent"`
	DetectedByAllianceIds    []Unknown `json:"detected_by_alliance_ids"`
}

type NodeAddress struct {
	Galaxy uint64  `json:"galaxy"`
	System uint64  `json:"system"`
	Planet *uint64 `json:"planet"` // can be null
}

type Buff struct {
	BuffId         uint               `json:"buff_id"`
	ActivatorId    int                `json:"activator_id"`   // found a -1 value
	ActivationTime ScopelyTime        `json:"activation_time"`
	ExpiryTime     ScopelyTime        `json:"expiry_time"`
	Attributes     map[string]float32 `json:"attributes"`
	Ranks          []int              `json:"ranks"`
}

type CrewData struct {
	Id    uint `json:"id"`
	Level int  `json:"level"`
	Rank  int  `json:"rank"`
}

