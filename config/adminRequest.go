package config

import "net/http"

const cacheSizeKey = "cache_size_key"

var configSizeKey = map[string]interface{}{
	cacheSizeKey: 50,
}

type adminRequest struct {}

func (ar *adminRequest) GetConfigSizeKey(options map[string]interface{}) {
	clientID := configSizeKey["client_id"].(string)
	clientSecret := configSizeKey["client_secret"].(string)

	if clientID == "" || clientSecret == "" {
		http.Error(rw, "Client ID or Client Secret not set", http.StatusInternalServerError)
	}
}


type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (ar *adminRequest) GenerateToken(w http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

	},
},