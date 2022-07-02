package models

type AccessTokenModel struct {

	// アクセストークンの値
	AccessToken string `json:"access_token,omitempty"`

	// アクセストークンが有効な残り秒数
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// Bearer
	TokenType string `json:"token_type,omitempty"`
}
