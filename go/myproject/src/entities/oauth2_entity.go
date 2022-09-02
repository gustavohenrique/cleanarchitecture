package entities

type Oauth2Entity struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Token        string `json:"token,omitempty"`
}
