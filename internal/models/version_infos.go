package models

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
