package stfc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

type Header struct {Key, Value string}

const (
	AppID        = "4af7c20b-7646-4fb7-b64f-ae0a8c51c1f1"
	UnityVersion = "2020.3.18f1-digit-multiple-fixes-build"
	UserAgent    = "UnityPlayer/" + UnityVersion + " (UnityWebRequest/1.0, libcurl/7.75.0-DEV)"
)

var (
	ErrNoSuccess          = errors.New("non-200 response code")
	ErrNilArgument        = errors.New("nil argument")
	ErrTypeNotFound       = errors.New("requested type not found")
	ErrNotImplemented     = errors.New("not implemented")
	ErrEmptyResponse      = errors.New("empty response")
	ErrSyncMissingPayload = errors.New("sync missing JSON payload")
)

func init() {
}

type AdhocCredentials struct {
	AdhocUsername string `json:"adhoc_username" bson:"adhoc_username"`
	AdhocPassword string `json:"adhoc_password" bson:"adhoc_password"`
}

type Session struct {
	LoginResponse AccountsLogin
	LiveHost      string
	Alive         bool
	galaxy        *Galaxy
}

type AllianceRequest struct {
	UserCurrentRank uint     `json:"user_current_rank"`
	AllianceID      uint64   `json:"alliance_id"`
	AllianceIDs     []uint64 `json:"alliance_ids"`
}

/*
 * GENERAL FUNCTIONS
 */

func ScopelyID(email, password string) (*AdhocCredentials, error) {
	return nil, ErrNotImplemented
}

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
	if len(body) < 1 {
		return nil, ErrEmptyResponse
	}
	var ret Session
	err = json.Unmarshal(body, &ret.LoginResponse)
	if err != nil {
		return nil, err
	}
	if ret.LoginResponse.HTTPCode != 200 {
		return nil, ErrNoSuccess
	}
	ret.LiveHost = fmt.Sprintf("https://live-%03d-web.startrek.digitgaming.com", ret.LoginResponse.InstanceAccount.InstanceIDCurrent)
	go ret.keepalive()
	return &ret, nil
}

/*
 * INSTANCE FUNCTIONS
 */

func (s *Session) Post(endpoint string, headers []Header, body io.Reader) ([]byte, error) {
	if s == nil {
		return nil, ErrNilArgument
	}
	var tee io.Reader
	var buf bytes.Buffer
	if body != nil {
		tee = io.TeeReader(body, &buf)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", s.LiveHost + endpoint, tee)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("X-Transaction-Id", "fd0ce62c-9843-439d-88b4-591ec2326d07")   // TODO: Randomize?
	req.Header.Set("X-Auth-Session-Id", s.LoginResponse.InstanceSessionID)
	req.Header.Set("X-Prime-Version", "1.000.31437")
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("X-Prime-Sync", "0")
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
		log.Printf("error code %d: %s", resp.StatusCode, resp.Status)
		log.Printf("request:\n%#v", req)
		if b, err := io.ReadAll(&buf); err == nil {
			log.Println("body (str):\n" + string(b))
			log.Println("body (hex):\n" + hex.EncodeToString(b))
		}
		return nil, ErrNoSuccess
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (s *Session) Sync(n int) (*SyncJSON, error) {
	body, err := s.Post("/sync", []Header{{"X-Prime-Sync", fmt.Sprintf("%d", n)}}, nil)
	if err != nil {
		return nil, err
	}
	var sync Sync
	if err := proto.Unmarshal(body, &sync); err != nil {
		return nil, err
	}
	if sync.Payload == nil {
		return nil, ErrSyncMissingPayload
	}
	var syncJson SyncJSON
	if err := json.Unmarshal([]byte(sync.Payload.Json), &syncJson); err != nil {
		return nil, err
	}
	return &syncJson, nil
}

func (s *Session) Profiles(userIds []string) ([]*Profile, error) {
	b, err := json.Marshal(map[string][]string{"user_ids": userIds})
	if err != nil {
		return nil, err
	}
	body, err := s.Post("/user_profile/profiles", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	var profiles Profiles
	if err := proto.Unmarshal(body, &profiles); err != nil {
		return nil, err
	}
	return profiles.Payload.Payload2.Profile, nil
}

//func (s *Session) AlliancesProto(allianceIds []uint64) ([]*AlliancesPublicInfo_AlliancePublicInfo, error) {
//	b, err := s.Alliances(allianceIds, 71)
//	if err != nil {
//		return nil, err
//	}
//	var details []*AlliancesPublicInfo_AlliancePublicInfo
//	if err = proto.Unmarshal(b, &details); err != nil {
//		return nil, err
//	}
//	return details, nil
//}

func (s *Session) AlliancesJson(allianceIds []uint64) (*AlliancesPublicInfoJson, error) {
	b, err := s.Alliances(allianceIds, 42)
	if err != nil {
		return nil, err
	}
	var details AlliancesPublicInfoWrapperJson
	if err = json.Unmarshal(b, &details); err != nil {
		return nil, err
	}
	if unwrapped, ok := details["alliances_info"]; ok {
		return unwrapped, nil
	}
	return nil, ErrTypeNotFound
}

func (s *Session) Alliances(allianceIds []uint64, requestedType uint32) ([]byte, error) {
	b, err := json.Marshal(&AllianceRequest{AllianceIDs: allianceIds})
	if err != nil {
		return nil, err
	}
	body, err := s.Post("/alliance/get_alliances_public_info", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	var alliance AllianceEndpoint
	if err := proto.Unmarshal(body, &alliance); err != nil {
		return nil, err
	}
	for _, d := range alliance.Details {
		if d.Type == requestedType {
			return d.Details, nil
		}
	}
	return nil, ErrTypeNotFound
}

func (s *Session) keepalive() {
	s.Alive = true
	for s.Alive {
		time.Sleep(time.Minute)
		_, err := s.Sync(1)   // TODO: Update state with received data
		if err != nil {
			log.Printf("Session %s keepalive failure: %s", s.LoginResponse.InstanceSessionID, err)
			s.Alive = false
		}
	}
}

/*
 * UTILITY FUNCTIONS
 */

// Lazy load the galaxy
func (s *Session) Galaxy() (*Galaxy, error) {
	if s.galaxy == nil {
		b, err := ioutil.ReadFile("static/nodes.json")
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, &s.galaxy)
		if err != nil {
			return nil, err
		}
	}
	return s.galaxy, nil
}

