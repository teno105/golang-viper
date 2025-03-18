package models

// ResultData JSON 구조체
type ResultData struct {
	InGameBoard  InGameBoard  `mapstructure:"in_game_board" json:"in_game_board"`
	LatestPolicy LatestPolicy `mapstructure:"latest_policy" json:"latest_policy"`
	VersionInfos VersionInfos `mapstructure:"version_infos" json:"version_infos"`
	Notice       Notice       `mapstructure:"notice" json:"notice"`
	Maintenance  Maintenance  `mapstructure:"maintenance" json:"maintenance"`
	StoreLink    StoreLink    `mapstructure:"store_link" json:"store_link"`
}

// InGameBoard JSON 구조체
type InGameBoard struct {
	Display bool   `mapstructure:"display" json:"display"`
	Url     string `mapstructure:"url" json:"url"`
}

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

// VersionInfos JSON 구조체
type VersionInfos struct {
	VersionInfo   VersionInfo         `mapstructure:"version_info" json:"version_info"`
	UpdateMessage []UpdateMessageData `mapstructure:"update_message" json:"update_message"`
}

type VersionInfo struct {
	StoreType      string `mapstructure:"store_type" json:"store_type"`
	VersionNo      string `mapstructure:"version_no" json:"version_no"`
	GameServerName string `mapstructure:"game_server_name" json:"game_server_name"`
	GameServerUrl  string `mapstructure:"game_server_url" json:"game_server_url"`
	VisiblePopup   bool   `mapstructure:"visible_popup" json:"visible_popup"`
	LatestVersion  string `mapstructure:"latest_version" json:"latest_version"`
	DynamicConfig  string `mapstructure:"dynamic_config" json:"dynamic_config"`
}

type UpdateMessageData struct {
	LanguageType string `mapstructure:"language_type" json:"language_type"`
	Title        string `mapstructure:"title" json:"title"`
	Body         string `mapstructure:"body" json:"body"`
}

// Notice JSON 구조체
type Notice struct {
	LocalList []LanguageData `mapstructure:"local_list" json:"local_list"`
}

type LanguageData struct {
	Language string `mapstructure:"language" json:"language,omitempty"`
	Title    string `mapstructure:"title" json:"title"`
	Body     string `mapstructureon:"body" json:"body"`
}

// Maintenance JSON 구조체
type Maintenance struct {
	Message []MaintenanceMessage `mapstructure:"message" json:"message"`
}

type MaintenanceMessage struct {
	Language    string `mapstructure:"language" json:"language,omitempty"`
	Title       string `mapstructure:"title" json:"title"`
	Body        string `mapstructure:"body" json:"body"`
	DetailedURL string `mapstructure:"detailed_url" json:"detailed_url"`
}

// StoreLink JSON 구조체
type StoreLink struct {
	PlatformType string `mapstructure:"platform_type" json:"platform_type"`
	StoreUrl     string `mapstructure:"store_url" json:"store_url"`
}
