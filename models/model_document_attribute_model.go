package models

import (
	"time"
)

type DocumentAttributeModel struct {

	// 書類のID
	DocumentId string `json:"document_id,omitempty" faker:"hyphenTruncatedUUID"`

	// 管理用タイトル
	Title string `json:"title,omitempty" faker:"lang=jpn,len=50"`

	// 契約相手の名称
	Counterparty string `json:"counterparty,omitempty" faker:"lang=jpn,len=15"`

	// 契約締結日 nullを許容する
	ContractAt string `json:"contract_at,omitempty" faker:"rfc3339"`

	// 契約開始日 nullを許容する
	ValidityStartAt string `json:"validity_start_at,omitempty" faker:"rfc3339"`

	// 契約終了日 nullを許容する
	ValidityEndAt string `json:"validity_end_at,omitempty" faker:"rfc3339"`

	// 解約通知期限 nullを許容する
	ValidityEndNoticeAt string `json:"validity_end_notice_at,omitempty" faker:"rfc3339"`

	// 自動更新の有無:   * 0 - 指定なし   * 1 - あり   * 2 - なし
	AutoUpdate int32 `json:"auto_update,omitempty" faker:"oneof: 0, 1, 2"`

	// 管理番号
	LocalId string `json:"local_id,omitempty"`

	// 取引金額 nullを許容する
	Amount int32 `json:"amount,omitempty" faker:"amount"`

	// 作成日時
	CreatedAt time.Time `json:"created_at,omitempty"`

	// 更新日時
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// ユーザー定義の項目
	Options []DocumentAttributeModelOptions `json:"options,omitempty" faker:"len=1"`
}
