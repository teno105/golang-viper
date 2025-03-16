package models

// LatestPolicy JSON 구조체
type LatestPolicy struct {
	Privacy   []PolicyData `json:"privacy"`
	Terms     []PolicyData `json:"terms"`
	NightPush []NightPush  `json:"night_push"`
}

type PolicyData struct {
	Language  string `json:"language"`
	Version   int    `json:"version"`
	StartDate string `json:"start_date"`
	Url       string `json:"url"`
	UrlForTxt string `json:"url_for_txt"`
}

type NightPush struct {
	Language string `json:"language"`
	Version  int    `json:"version"`
	Body     string `json:"body"`
}
