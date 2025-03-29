package models

// ResultData JSON 구조체
type ResultData struct {
	VersionInfos VersionInfos `mapstructure:"version_infos" json:"version_infos"`
	Maintenance  Maintenance  `mapstructure:"maintenance" json:"maintenance"`
	StoreLink    StoreLink    `mapstructure:"store_link" json:"store_link"`
}

// VersionInfos JSON 구조체
type VersionInfos struct {
	VersionInfo VersionInfo `mapstructure:"version_info" json:"version_info"`
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
