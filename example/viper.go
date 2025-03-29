package example

import (
	"fmt"
	"os"
	"path/filepath"
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
