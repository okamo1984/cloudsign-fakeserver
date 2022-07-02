package models

type DocumentModel struct {

	// 書類ID
	Id string `json:"id,omitempty" faker:"hyphenTruncatedUUID"`

	// 作成者ID
	UserId string `json:"user_id,omitempty" faker:"hyphenTruncatedUUID"`

	// 書類のタイトル
	Title string `json:"title,omitempty" faker:"lang=jpn,len=50"`

	// 契約相手の名称などのメモ。受信者には表示されない
	Note string `json:"note,omitempty" faker:"lang=jpn,len=50"`

	// 受信者に対するメッセージ
	Message string `json:"message,omitempty" faker:"lang=jpn,len=50"`

	// 書類の状態:  * 0 - 下書き  * 1 - 先方確認中  * 2 - 締結済  * 3 - 取消、または却下  * 4 - テンプレート  * 13 - インポート書類
	Status int64 `json:"status,omitempty" faker:"oneof 0, 1, 2, 3, 4, 13"`
 
	// 転送許可設定
	CanTransfer bool `json:"can_transfer,omitempty"`

	// 親展書類設定
	Private bool `json:"private,omitempty"`

	// 送信日時（RFC3339準拠）
	SentAt string `json:"sent_at,omitempty" faker:"rfc3339"`

	// 書類の最終処理日時（RFC3339準拠） documentModel.statusが以下のいずれかである場合は、最後の同意/却下日時。participants中のparticipantModel.processed_atの最新と等しい。  status:   * 1 - 先方確認中   * 2 - 締結済   * 3 - 取消、または却下   * 13 - インポート書類  documentModel.statusが上記のいずれでもない場合は、書類の最終更新日時。
	LastProcessedAt string `json:"last_processed_at,omitempty" faker:"rfc3339"`

	// 作成日時（RFC3339準拠）
	CreatedAt string `json:"created_at,omitempty" faker:"rfc3339"`

	// 更新日時（RFC3339準拠）
	UpdatedAt string `json:"updated_at,omitempty" faker:"rfc3339"`

	// 参加者の一覧（作成者を含む）
	Participants []ParticipantModel `json:"participants,omitempty" faker:"boundary_start=1, boundary_end=3"`

	// ファイルの一覧
	Files []FileModel `json:"files,omitempty" faker:"boundary_start=1, boundary_end=3"`

	// 共有先の一覧
	Reportees []ReporteeModel `json:"reportees,omitempty" faker:"boundary_start=1, boundary_end=3"`
}
