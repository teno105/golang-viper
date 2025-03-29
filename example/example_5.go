package example

import (
	"fmt"
	"golang-viper/infra"
	"golang-viper/internal/models"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// YAML 파일을 읽어 구조체로 변환하는 함수
func loadFile(filePath string, target interface{}) error {
	// 실행 중인 프로그램의 경로 가져오기
	rootPath := infra.GetProjectRoot()
	if rootPath == "" {
		return fmt.Errorf("작업 디렉토리를 가져오는 중 오류 발생 %s", filePath)
	}

	path := filepath.Join(rootPath, filePath)
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("파일이 존재하지 않음: %s %v", filePath, err)
		return err
	}

	// 설정 파일을 구조체로 변환
	if err := viper.Unmarshal(target); err != nil {
		return err
	}

	return nil
}

func Example5() {

	var resultData models.ResultData

	// version_infos.yml 읽기
	if err := loadFile("data/12/version_infos.yml", &resultData); err != nil {
		log.Printf("[WARN] version_infos.yml 파일로딩을 실패: %v", err)
	}

	// store_link.yml 읽기
	if err := loadFile("data/12/store_link.yml", &resultData); err != nil {
		log.Printf("[WARN] store_link.yml 파일로딩을 실패: %v", err)
	}

	// maintenance.yml 읽기
	if err := loadFile("data/12/maintenance.yml", &resultData); err != nil {
		log.Printf("[WARN] maintenance.yml 파일로딩을 실패: %v", err)
	}

	fmt.Printf("Version StoreType: %s, VersionNo: %s\n", resultData.VersionInfos.VersionInfo.StoreType, resultData.VersionInfos.VersionInfo.VersionNo)
	fmt.Printf("Store PlatformType: %s, Url: %s\n", resultData.StoreLink.PlatformType, resultData.StoreLink.StoreUrl)
	fmt.Printf("Maintenance Message: %s\n", resultData.Maintenance.Message[0].Body)
}
