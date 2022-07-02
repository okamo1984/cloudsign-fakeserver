package models

type FileModel struct {

	// ファイルの ID
	Id string `json:"id,omitempty" faker:"hyphenTruncatedUUID"`

	// ファイルのタイトル
	Name string `json:"name,omitempty" faker:"lang=jpn,word"`

	// 書類内のファイルの順序
	Order int64 `json:"order,omitempty" faker:"boundary_start=0, boundary_end=3"`

	// ファイルのページ数
	TotalPages int64 `json:"total_pages,omitempty" faker:"boundary_start=0, boundary_end=3"`

	// 入力項目の一覧
	Widgets []WidgetModel `json:"widgets,omitempty"`
}
