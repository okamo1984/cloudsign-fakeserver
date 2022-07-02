package models

type DocumentListModel struct {

	// 該当する書類の総数
	Total int64 `json:"total,omitempty" faker:"oneof: 5"`

	// 指定されたページ番号。ページあたりの書類数は最大100件。
	Page int64 `json:"page,omitempty" faker:"oneof: 1"`

	// 書類データ
	Documents []DocumentModel `json:"documents,omitempty"`
}
