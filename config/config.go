package config

import (
	"errors"
)

var (
	clientId, clientSecret, platformId, corporationId, integratorId, accessToken, refreshToken string
	)
var productId = []string{""}
var schema = "https"
var host = "api.gopago.com"
var userAgent = []string{"GoPago"}

type Configurations struct {
	Sandbox           	bool
	ShowPromiseError  	bool
	CacheMaxSize      	uint16 //CacheMaxSize = 50
	ClientID 			uint16
	ClientSecret 		uint16
	AccessToken 		string
	ProductId 			string
	TrackingId 			string
	UserAgent 			string
}

//Añadir ClientID, ClientSecret, AccessToken y ... en ruta admin

func Configure(configs Configurations) {
	if configs.ClientID == "" && configs.ClientSecret == "" && configs.AccessToken == "" {
		panic(errors.New("You must provide a method of authentication (ClientId & ClientSecret or AccessToken)"))
	}
	if (configs.ClientId != "" && configs.ClientSecret == "") ||
		(configs.ClientId == "" && configs.ClientSecret != "") {
		panic(errors.New("You must provide both ClientId and ClientSecret"))
	}

	clientId = configs.ClientId
	clientSecret = configs.ClientSecret

	accessToken = configs.AccessToken
	platformId = configs.PlatformId
	corporationId = configs.CorporationId
	integratorId = configs.IntegratorId

	Configurations.Sandbox = configs.Sandbox
	Configurations.ShowPromiseError = configs.ShowPromiseError

	if !Configurations.ShowPromiseError {
	
	}
}


func GetClientId() string {
	return clientId
}


func GetClientSecret() string {
	return clientSecret
}


func GetPlatformId() string {
	return platformId
}


func GetCorporationId() string {
	return corporationId
}


func GetIntegratorId() string {
	return integratorId
}


func SetAccessToken(token string) {
	accessToken = token
}


func GetAccessToken() string {
	return accessToken
}


func SetRefreshToken(token string) {
	refreshToken = token
}


func GetRefreshToken() string {
	return refreshToken
}


func GetBaseUrl() string {
	return schema + "://" + host
}


func GetProductId() string {
	return ProductId
}


func GetTrackingId() string {
	return TrackingID
}


func GetUserAgent() string {
	return userAgent
}
