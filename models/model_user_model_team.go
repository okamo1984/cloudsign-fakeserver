package models

// UserModelTeam - ユーザーが所属するチーム
type UserModelTeam struct {

	// チームの ID
	Id string `json:"id,omitempty"`

	// チームの名前
	Name string `json:"name,omitempty"`
}
