package stfc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Header struct {Key, Value string}

const (
	AppID        = "4af7c20b-7646-4fb7-b64f-ae0a8c51c1f1"
	UnityVersion = "2020.3.18f1-digit-multiple-fixes-build"
	UserAgent    = "UnityPlayer/" + UnityVersion + " (UnityWebRequest/1.0, libcurl/7.75.0-DEV)"
)

var (
	ErrNoSuccess   = errors.New("non-200 response code")
	ErrNilArgument = errors.New("nil argument")
)

func init() {
}

type Session struct {
	LoginResponse AccountsLogin
	LiveHost      string
}

/*
 * GENERAL FUNCTIONS
 */

func Login(username, password string) (*Session, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://cdn-nv3-live.startrek.digitgaming.com/accounts/v1/accounts/login/windows", nil)
	req.Header.Set("User-Agent",       UserAgent)
	req.Header.Set("Accept",           "*/*")
	req.Header.Set("Accept-Encoding",  "deflate")
	req.Header.Set("X-TRANSACTION-ID", "ca98560c-e47c-4af5-b859-e35764181733")   // TODO: Randomize?
	req.Header.Set("X-PRIME-VERSION",  "1.000.31110")
	req.Header.Set("X-Suppress-Codes", "1")
	req.Header.Set("Content-Type",     "application/x-www-form-urlencoded")
	req.Header.Set("X-Api-Key",        "FCX2QsbxHjSP52B")
	req.Header.Set("X-PRIME-SYNC",     "0")
	req.Header.Set("X-Unity-Version",  UnityVersion)
	req.Header.Set("Connection",       "close")
	data := url.Values{
		"auth_provider":   {""},
		"auth_token":      {""},
		"ad_hoc_username": {username},
		"ad_hoc_password": {password},
		"email":           {""},
		"password":        {""},
		"channel":         {"digit_WindowsPlayer"},
		"partner_id":      {"8abbfd6b-90c5-44bd-9d90-28d72cb203ff"},
	}
	req.Body = io.NopCloser(strings.NewReader(data.Encode()))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrNoSuccess
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
	ret.LiveHost = fmt.Sprintf("https://live-%03d-web.startrek.digitgaming.com", ret.LoginResponse.InstanceAccount.InstanceIDCurrent)
	return &ret, nil
}

/*
 * INSTANCE FUNCTIONS
 */

func (s *Session) Post(endpoint string, headers []Header, body io.Reader) ([]byte, error) {
	if s == nil {
		return nil, ErrNilArgument
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", s.LiveHost + endpoint, body)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("X-Transaction-Id", "fd0ce62c-9843-439d-88b4-591ec2326d07")   // TODO: Randomize?
	req.Header.Set("X-Auth-Session-Id", s.LoginResponse.InstanceSessionID)
	req.Header.Set("X-Prime-Version", "1.000.31309")
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Accept", "application/x-protobuf")   // Ripper recommends application/json but testing doesn't show a difference in return value
	req.Header.Set("X-Unity-Version", UnityVersion)
	if headers != nil {
		for _, h := range headers {
			req.Header.Set(h.Key, h.Value)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrNoSuccess
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (s *Session) Sync(n int) ([]byte, error) {
	return s.Post("/sync", []Header{{"X-Prime-Sync", fmt.Sprintf("%d", n)}}, nil)
}

