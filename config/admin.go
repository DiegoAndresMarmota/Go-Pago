package config

import "errors"

type SettingsConfiguration struct {
	ClientID      string
	ClientSecret  string
	AccessToken   string
	RefreshToken  string
	PlatformID    string
	CorporationID string
	IntegratorID  string
	CacheMaxSize  int
	BaseURL       string
	Sandbox       bool
	UserAgent     string
	ProductID     string
	TrackingID    string
}

// type configurations struct {
// 	configurations SettingsConfiguration
// }

func (st *SettingsConfiguration) generateAccessToken() (string, error) {
	if st.AccessToken != "" {
		return st.AccessToken, nil
	}
	if st.ClientID == "" || st.ClientSecret == "" {
		return "", errors.New("Must set client_id and client_secret")
	}
	var data map[string]interface{}
	accessToken, ok := data["access_token"].(string)
	if !ok {
		return "", errors.New("Invalid response format for access token")
	}
	st.AccessToken = accessToken
	return accessToken, nil
}


func (st *SettingsConfiguration) clone(options map[string]interface{}) map[string]interface{} {
	if options == nil {
		return make(map[string]interface{})
	}
	cloneMap := make(map[string]interface{})
	for k, v := range options {
		cloneMap[k] = v
	}
	return cloneMap
}
