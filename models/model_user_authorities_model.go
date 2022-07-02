package models

type UserAuthoritiesModel struct {

	// チーム管理権限。チームの設定変更とメンバー管理が行える。
	TeamAdministrator bool `json:"team_administrator,omitempty"`

	// 書類管理権限。チーム内で送受信された書類を閲覧できる。
	TeamDocumentAdministrator bool `json:"team_document_administrator,omitempty"`

	// 承認者。チームの承認機能が有効な場合、書類の宛先に承認者として追加できる。
	DocumentApprover bool `json:"document_approver,omitempty"`

	// 親展書類管理権限。チーム内で送受信された親展書類を閲覧できる。
	PrivateDocumentAdministrator bool `json:"private_document_administrator,omitempty"`
}
