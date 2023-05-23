package stfc

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func init() {
}

type Session struct {
	LoginResponse interface{}
}

func Login(username, password string) (*Session, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://cdn-nv3-live.startrek.digitgaming.com/accounts/v1/accounts/login/windows", nil)
	req.Header.Add("User-Agent",       "UnityPlayer/2020.3.18f1-digit-multiple-fixes-build (UnityWebRequest/1.0, libcurl/7.75.0-DEV)")
	req.Header.Add("X-TRANSACTION-ID", "ca98560c-e47c-4af5-b859-e35764181733")   // TODO: Randomize
	req.Header.Add("X-PRIME-VERSION",  "1.000.31110")
	req.Header.Add("X-Suppress-Codes", "1")
	req.Header.Add("X-Api-Key",        "FCX2QsbxHjSP52B")
	req.Header.Add("X-PRIME-SYNC",     "0")
	req.Header.Add("X-Unity-Version",  "2020.3.18f1-digit-multiple-fixes-build")
	req.Form = url.Values{
		"auth_provider":   {""},
		"auth_token":      {""},
		"ad_hoc_username": {username},
		"ad_hoc_password": {password},
		"email":           {""},
		"password":        {""},
		"channel":         {"digit_WindowsPlayer"},
		"partner_id":      {"8abbfd6b-90c5-44bd-9d90-28d72cb203ff"},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret Session
	err = json.Unmarshal(body, &ret.LoginResponse)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

