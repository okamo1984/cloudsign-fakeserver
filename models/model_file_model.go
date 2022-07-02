package models

type FileModel struct {

	// ファイルの ID
	Id string `json:"id,omitempty"`

	// ファイルのタイトル
	Name string `json:"name,omitempty"`

	// 書類内のファイルの順序
	Order int64 `json:"order,omitempty"`

	// ファイルのページ数
	TotalPages int64 `json:"total_pages,omitempty"`

	// 入力項目の一覧
	Widgets []WidgetModel `json:"widgets,omitempty"`
}
