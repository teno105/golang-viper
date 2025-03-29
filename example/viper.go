package example

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// 프로젝트 루트 찾기
func findProjectRoot() (string, error) {
	// 실행 중인 프로그램의 경로 가져오기
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// 실행 경로에서 상위 디렉토리로 이동하면서 go.mod를 찾음
	dir := filepath.Dir(executablePath)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// 상위 디렉토리로 이동
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("프로젝트 루트를 찾을 수 없음")
}

// YAML 파일을 읽어 구조체로 변환하는 함수
func LoadYamlFile(filePath string, target interface{}) error {
	root, err := findProjectRoot()
	if err != nil {
		err := fmt.Errorf("go.mod 파일을 찾을 수 없습니다.")
		fmt.Println("작업 디렉토리를 가져오는 중 오류 발생:", err)
		return err
	}

	path := filepath.Join(root, filePath)
	viper.SetConfigFile(path)

	//viper.AddConfigPath(".")
	//viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("[WARN] 파일이 존재하지 않음: %s", filePath)
		return err
	}

	if err := viper.Unmarshal(target); err != nil {
		return err
	}

	return nil
}
