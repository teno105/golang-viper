package models

// Notice JSON 구조체
type Notice struct {
	LocalList []LanguageData `json:"local_list"`
}

type LanguageData struct {
	Language string `json:"language,omitempty"` // "name" 대신 "language" 사용
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type NoticeRoot struct {
	Notice Notice `json:"notice"`
}
