package models

// DocumentAttributeUpsertRequest - 書類情報作成・更新用 オブジェクト
type DocumentAttributeUpsertRequest struct {

	// 管理用タイトル
	Title string `json:"title,omitempty"`

	// 契約相手の名称
	Counterparty string `json:"counterparty,omitempty"`

	// 契約締結日 nullを許容する
	ContractAt string `json:"contract_at,omitempty"`

	// 契約開始日 nullを許容する
	ValidityStartAt string `json:"validity_start_at,omitempty"`

	// 契約終了日 nullを許容する
	ValidityEndAt string `json:"validity_end_at,omitempty"`

	// 解約通知期限 nullを許容する
	ValidityEndNoticeAt string `json:"validity_end_notice_at,omitempty"`

	// 自動更新の有無:   * 0 - 指定なし   * 1 - あり   * 2 - なし
	AutoUpdate int32 `json:"auto_update,omitempty"`

	// 管理番号
	LocalId string `json:"local_id,omitempty"`

	// 取引金額 nullを許容する
	Amount int32 `json:"amount,omitempty"`

	Options []DocumentAttributeUpsertRequestOptions `json:"options,omitempty"`
}
