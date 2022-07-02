package models

type ReporteeModel struct {

	// 共有先の ID
	Id string `json:"id,omitempty"`

	// 共有先のメールアドレス
	Email string `json:"email,omitempty"`

	// 共有先の名前
	Name string `json:"name,omitempty"`

	// 共有先の会社名など
	Organization string `json:"organization,omitempty"`
}
