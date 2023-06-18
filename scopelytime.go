package stfc

import (
	"time"
)

const ScopelyTimeLayout = `"2006-01-02T15:04:05"`

type ScopelyTime time.Time

func (t *ScopelyTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == "" || str == "null" {
		return nil
	}
	tm, err := time.Parse(ScopelyTimeLayout, str)
	if err != nil {
		return err
	}
	*t = ScopelyTime(tm)
	return nil
}

func (t *ScopelyTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*t).Format(ScopelyTimeLayout)), nil
}

