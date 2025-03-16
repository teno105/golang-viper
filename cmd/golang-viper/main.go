package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"golang-viper/infra"
	"golang-viper/internal/models"
)

// GameData 구조체
type GameData struct {
	InGameBoard  models.InGameBoard  `json:"in_game_board"`
	LatestPolicy models.LatestPolicy `json:"latest_policy"`
	VersionInfos models.VersionInfos `json:"version_infos"`
	Notice       models.Notice       `json:"notice"`
	Maintenance  models.Maintenance  `json:"maintenance"`
	StoreLink    models.StoreLink    `json:"store_link"`
}

// 특정 게임 ID에 해당하는 데이터를 로드하는 함수
func loadGameData(gameID string) (*GameData, error) {
	inGameBoardPath := filepath.Join("data", gameID, "in_game_board.json")
	latestPolicyPath := filepath.Join("data", gameID, "latest_policy.json")
	versionInfosPath := filepath.Join("data", gameID, "version_infos.json")
	noticePath := filepath.Join("data", gameID, "notice.json")
	maintenancePath := filepath.Join("data", gameID, "maintenance.json")
	storeLinkPath := filepath.Join("data", gameID, "store_link.json")

	var gameData GameData

	// in_game_board.yml 읽기
	if err := infra.LoadYamlFile(inGameBoardPath, &gameData); err != nil {
		log.Printf("[WARN] in_game_board.json 파일로딩을 실패: %v", err)
		return nil, err
	}

	// version_infos.yml 읽기
	if err := infra.LoadYamlFile(versionInfosPath, &gameData); err != nil {
		log.Printf("[WARN] version_infos.json 파일로딩을 실패 %v", err)
		return nil, err
	}

	// latest_policy.yml 읽기
	if err := infra.LoadYamlFile(latestPolicyPath, &gameData); err != nil {
		log.Printf("[WARN] latest_policy.json 파일로딩을 실패: %v", err)
		return nil, err
	}

	// notice.yml 읽기
	if err := infra.LoadYamlFile(noticePath, &gameData); err != nil {
		log.Printf("[WARN] notice.json 파일로딩을 실패: %v", err)
		return nil, err
	}

	// maintenance.yml 읽기
	if err := infra.LoadYamlFile(maintenancePath, &gameData); err != nil {
		log.Printf("[WARN] maintenance.json 파일로딩을 실패: %v", err)
		return nil, err
	}

	// store_link.yml 읽기
	if err := infra.LoadYamlFile(storeLinkPath, &gameData); err != nil {
		log.Printf("[WARN] store_link.json 파일로딩을 실패 %v", err)
		return nil, err
	}

	return &gameData, nil
}

func main() {

	r := gin.Default()

	// GET /v2/init_data/games/:id 엔드포인트
	r.GET("/v2/init_data/games/:id", func(c *gin.Context) {
		gameID := c.Param("id")           // URL에서 game ID 가져오기
		data, err := loadGameData(gameID) // 해당 ID의 데이터를 로드
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game data not found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	// 서버 실행
	fmt.Println("Server is running on port 9095...")
	r.Run(":9095")
}
