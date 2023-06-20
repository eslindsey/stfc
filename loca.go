package stfc

import (
	"encoding/json"
	"io/ioutil"
)

// Loca is a map of Ripper-style ID_key: text values.
type Loca map[string]string

// An entry from a Ripper-style JSON localisation file.
type LocaEntry struct {
	Id   string `json:"id"`
	Text string `json:"text"`
	Key  string `json:"key"`
}

// LocaFromFile converts a Ripper-style JSON localisation file to a map of
// ID_key: text values.
func LocaFromFile(file string, l *Loca) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var locas []LocaEntry
	err = json.Unmarshal(b, &locas)
	if err != nil {
		return err
	}
	*l = map[string]string{}
	for _, loca := range locas {
		(*l)[loca.Id + "_" + loca.Key] = loca.Text
	}
	return nil
}

