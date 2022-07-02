package models

type ParticipantModel struct {

	// 宛先ID
	Id string `json:"id,omitempty" faker:"hyphenTruncatedUUID"`

	// 宛先のメールアドレス
	Email string `json:"email,omitempty" faker:"email"`

	// 宛先の名前
	Name string `json:"name,omitempty" faker:"lang=jpn,name,len=10"`

	// 宛先の会社名
	Organization string `json:"organization,omitempty" faker:"lang=jpn,name,len=15"`

	// 書類内の宛先の順序（送信者は 0）
	Order int64 `json:"order,omitempty" faker:"boundary_start=0, boundary_end=3"`

	// 宛先の状態:  * 0 - アクセス不可  * 2 - 下書き  * 3 - 配信待ち  * 4 - 確認待ち  * 6 - 送信済み  * 7 - 確認済み  * 8 - 押印または入力済み  * 9 - 却下  * 10 - キャンセル  * 12 - 署名中
	Status int32 `json:"status,omitempty" faker:"oneof: 0, 2, 3, 4, 6, 7, 8, 9, 10, 12"`

	// 宛先に設定されているアクセスコード。APIを使用しているユーザーが値を設定した場合のみレスポンスに含まれる。
	AccessCode string `json:"access_code,omitempty" faker:"password"`

	// 宛先の言語設定。ja（日本語）、en（英語）、zh-CHS（簡体字）、zh-CHT（繁体字）のいずれか。
	LanguageCode string `json:"language_code,omitempty" faker:"oneof: ja en zh-CHS"`

	// 各受信者による書類の同意/却下日時（RFC3339準拠）
	ProcessedAt string `json:"processed_at,omitempty" faker:"rfc3339"`

	// URL有効期限（RFC3339準拠）。statusが4の場合のみレスポンスに含まれる。
	AccessExpiresAt string `json:"access_expires_at,omitempty" faker:"rfc3339"`
}
