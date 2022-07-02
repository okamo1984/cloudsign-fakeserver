package models

type ReporteeModel struct {

	// 共有先の ID
	Id string `json:"id,omitempty" faker:"hyphenTruncatedUUID"`

	// 共有先のメールアドレス
	Email string `json:"email,omitempty" faker:"email"`

	// 共有先の名前
	Name string `json:"name,omitempty" faker:"lang=jpn,name"`

	// 共有先の会社名など
	Organization string `json:"organization,omitempty" faker:"lang=jpn,name"`
}
