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
