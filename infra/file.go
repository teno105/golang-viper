package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// go.mod 파일을 기준으로 프로젝트 루트 찾기
func getProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	log.Fatal("go.mod 파일을 찾을 수 없습니다.")
	return ""
}

// JSON 파일을 읽어 구조체로 변환하는 함수
func LoadJSONFile(filePath string, target interface{}) error {
	rootPath := getProjectRoot()
	if rootPath == "" {
		err := fmt.Errorf("go.mod 파일을 찾을 수 없습니다.")
		fmt.Println("작업 디렉토리를 가져오는 중 오류 발생:", err)
		return err
	}

	jsonPath := filepath.Join(rootPath, filePath)

	// 파일 경로와 파일 존재 여부 확인
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		log.Printf("[WARN] 파일이 존재하지 않음: %s", filePath)
		return err
	}

	file, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &target)
}
