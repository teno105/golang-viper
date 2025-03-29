package infra

import (
	"path/filepath"
	"runtime"
)

// 프로젝트 루트 찾기
func GetProjectRoot() string {
	// 실행 중인 프로그램의 경로 가져오기
	_, filename, _, _ := runtime.Caller(0)
	filename = filepath.Join(filename, "..")
	return filepath.Dir(filename)
}
