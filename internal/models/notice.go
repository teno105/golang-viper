package models

// Notice JSON 구조체
type Notice struct {
	LocalList []LanguageData `mapstructure:"local_list" json:"local_list"`
}

type LanguageData struct {
	Language string `mapstructure:"language" json:"language,omitempty"`
	Title    string `mapstructure:"title" json:"title"`
	Body     string `mapstructureon:"body" json:"body"`
}
