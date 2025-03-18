package models

// LatestPolicy JSON 구조체
type LatestPolicy struct {
	Privacy   []PolicyData `mapstructure:"privacy" json:"privacy"`
	Terms     []PolicyData `mapstructure:"terms" json:"terms"`
	NightPush []NightPush  `mapstructure:"night_push" json:"night_push"`
}

type PolicyData struct {
	Language  string `mapstructure:"language" json:"language"`
	Version   int    `mapstructure:"version" json:"version"`
	StartDate string `mapstructure:"start_date" json:"start_date"`
	Url       string `mapstructure:"url" json:"url"`
	UrlForTxt string `mapstructure:"url_for_txt" json:"url_for_txt"`
}

type NightPush struct {
	Language string `mapstructure:"language" json:"language"`
	Version  int    `mapstructure:"version" json:"version"`
	Body     string `mapstructure:"body" json:"body"`
}
