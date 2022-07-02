package models

type UserModel struct {

	// ユーザーの ID
	Id string `json:"id,omitempty"`

	// ユーザーのメールアドレス
	Email string `json:"email,omitempty"`

	// ユーザーの名前
	Name string `json:"name,omitempty"`

	// ユーザーの組織の名称
	Organization string `json:"organization,omitempty"`

	Team UserModelTeam `json:"team,omitempty"`

	Authorities UserModelAuthorities `json:"authorities,omitempty"`
}
