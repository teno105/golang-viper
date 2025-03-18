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

type ResponseData struct {
	Code       string            `json:"code"`
	ResultData models.ResultData `json:"result_data"`
}

// 특정 게임 ID에 해당하는 데이터를 로드하는 함수
func loadGameData(gameID string) (*models.ResultData, error) {
	inGameBoardPath := filepath.Join("data", gameID, "in_game_board.yml")
	latestPolicyPath := filepath.Join("data", gameID, "latest_policy.yml")
	versionInfosPath := filepath.Join("data", gameID, "version_infos.yml")
	noticePath := filepath.Join("data", gameID, "notice.yml")
	maintenancePath := filepath.Join("data", gameID, "maintenance.yml")
	storeLinkPath := filepath.Join("data", gameID, "store_link.yml")

	var resultData models.ResultData

	// in_game_board.yml 읽기
	if err := infra.LoadYamlFile(inGameBoardPath, &resultData); err != nil {
		log.Printf("[WARN] in_game_board.yml 파일로딩을 실패: %v", err)
		return nil, err
	}

	// version_infos.yml 읽기
	if err := infra.LoadYamlFile(versionInfosPath, &resultData); err != nil {
		log.Printf("[WARN] version_infos.yml 파일로딩을 실패 %v", err)
		return nil, err
	}

	// latest_policy.yml 읽기
	if err := infra.LoadYamlFile(latestPolicyPath, &resultData); err != nil {
		log.Printf("[WARN] latest_policy.yml 파일로딩을 실패: %v", err)
		return nil, err
	}

	// notice.yml 읽기
	if err := infra.LoadYamlFile(noticePath, &resultData); err != nil {
		log.Printf("[WARN] notice.yml 파일로딩을 실패: %v", err)
		return nil, err
	}

	// maintenance.yml 읽기
	if err := infra.LoadYamlFile(maintenancePath, &resultData); err != nil {
		log.Printf("[WARN] maintenance.yml 파일로딩을 실패: %v", err)
		return nil, err
	}

	// store_link.yml 읽기
	if err := infra.LoadYamlFile(storeLinkPath, &resultData); err != nil {
		log.Printf("[WARN] store_link.yml 파일로딩을 실패 %v", err)
		return nil, err
	}

	return &resultData, nil
}

func main() {

	r := gin.Default()

	// GET /v2/init_data/games/:id 엔드포인트
	r.GET("/v2/init_data/games/:id", func(c *gin.Context) {
		gameID := c.Param("id")                 // URL에서 game ID 가져오기
		resultData, err := loadGameData(gameID) // 해당 ID의 데이터를 로드
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game data not found"})
			return
		}

		// 결과 데이터를 JSON으로 반환
		res := ResponseData{
			Code:       "0",
			ResultData: *resultData,
		}
		c.JSON(http.StatusOK, res)
	})

	// 서버 실행
	fmt.Println("Server is running on port 9095...")
	r.Run(":9095")
}
