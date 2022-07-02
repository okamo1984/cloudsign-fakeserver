package models

type ErrorModel struct {

	// エラーの種類
	Error string `json:"error,omitempty" faker:"word"`

	// エラーメッセージ
	Message string `json:"message,omitempty" faker:"sentence"`
}
