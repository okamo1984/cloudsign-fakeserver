package models

type WidgetModel struct {

	// 入力項目の ID
	Id string `json:"id,omitempty" faker:"hyphenTruncatedUUID"`

	// 入力項目の種類:  * 0 - 押印  * 1 - フリーテキスト  * 2 - チェックボックス
	WidgetType int32 `json:"widget_type,omitempty" faker:"oneof: 0, 1, 2"`

	// 入力が割り当てられている宛先の ID
	ParticipantId string `json:"participant_id,omitempty" faker:"hyphenTruncatedUUID"`

	// 入力対象のファイルの ID
	FileId string `json:"file_id,omitempty" faker:"hyphenTruncatedUUID"`

	// 入力項目が設定されている対象ファイルのページ番号
	Page int64 `json:"page,omitempty" faker:"boundary_start=0, boundary_end=3"`

	// 入力項目左上の対象ファイル・対象ページにおける設置位置の X 座標
	X int64 `json:"x,omitempty"`

	// 入力項目左上の対象ファイル・対象ページにおける設置位置の Y 座標
	Y int64 `json:"y,omitempty"`

	// 入力項目の幅
	W int64 `json:"w,omitempty"`

	// 入力項目の高さ
	H int64 `json:"h,omitempty"`

	// 入力項目に入力されたテキスト
	Text string `json:"text,omitempty" faker:"lang=jpn,sentence"`

	// 入力項目の状態:   * 0 - 未入力   * 1 - 入力済み
	Status int32 `json:"status,omitempty" faker:"oneof: 0, 1"`

	// 入力項目に入力されたラベル
	Label string `json:"label,omitempty" faker:"lang=jpn,word"`

	// widget_type が 2 (チェックボックス) の場合のみ含まれる、必須/任意の設定:   * true - 必須   * false - 任意
	Required bool `json:"required,omitempty"`
}
