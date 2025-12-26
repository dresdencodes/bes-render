package config 

import (
	"github.com/go-resty/resty/v2"
)

// PocketAuth represents the authentication response from PocketBase
type PocketAuth struct {
	Token  		string 		`json:"token"`
	Record struct {
		ID             		string 		`json:"id"`
		CollectionID   		string 		`json:"collectionId"`
		CollectionName 		string 		`json:"collectionName"`
		Username       		string 		`json:"username"`
		Email          		string 		`json:"email"`
		Created        		string 		`json:"created"`
		Updated        		string 		`json:"updated"`
	} 	`json:"record"`
}


var client = resty.New()

func PocketAuthorize(pocketIP string) (*PocketAuth, error) {
	
	// decode pocket auth
	var pocketAuth PocketAuth

	// client resty
	_, err := client.R().
		SetBody(map[string]string{
			"identity":"bes-chromie@servers.io",
			"password":"syquest123",
		}).
		SetResult(&pocketAuth).
		Post(pocketIP + "/api/collections/users/auth-with-password")
	return &pocketAuth, err

}