package stfc

type AccountsLogin struct {
	Code            int `json:"code"`
	InstanceAccount struct {
		AccountId          string `json:"account_id"`
		MasterId           string `json:"master_id"`
		ProductCode        string `json:"product_code"`
		InstanceId         int    `json:"instance_id"`
		Name               string `json:"name"`
		Status             int    `json:"status"`
		Created            string `json:"created"` // TODO: Parse to time.Time
		SourcePartnerId    string `json:"source_partner_id"`
		SourceChannel      string `json:"source_channel"`
		SourceReferrer     string `json:"source_referrer"`
		SourceCampaign     string `json:"source_campaign"`
		Language           string `json:"language"`
		ArchiveState       int    `json:"archive_state"`
		ArchiveKey         string `json:"archive_key"`
		InstanceIdPrevious int    `json:"instance_id_previous"`
		InstanceIdCurrent  int    `json:"instance_id_current"`
		CountryCode        string `json:"country_code"`
	} `json:"instance_account"`
	SessionId         string      `json:"session_id"`
	InstanceSessionId string      `json:"instance_session_id"`
	AdHocUsername     string      `json:"ad_hoc_username"`
	AdHocPassword     string      `json:"ad_hoc_password"`
	TOS               interface{} `json:"tos"`
	HTTPCode          int         `json:"http_code"`
	Version           string      `json:"version"`
}

