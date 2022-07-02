package models

type DocumentAttributeModelOptions struct {

	// ユーザー定義の項目の番号
	Order int32 `json:"order,omitempty"`

	// ユーザー定義の項目の値
	Content string `json:"content,omitempty"`
}
