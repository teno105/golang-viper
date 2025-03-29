package example

import (
	"fmt"
	"golang-viper/infra"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// 프로젝트 루트 찾기
func Example1() {
	rootPath := infra.GetProjectRoot()

	// 실행 중인 프로그램의 경로 가져오기
	viper.SetConfigName("example_1")                     // example1.yaml
	viper.SetConfigType("yaml")                          // 파일 타입 지정
	viper.AddConfigPath(filepath.Join(rootPath, "data")) // 현재 디렉토리에서 설정 파일 찾기

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("설정 파일 읽기 실패: %v", err)
	}

	// 값 가져오기
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")

	fmt.Printf("DB Host: %s, Port: %d\n", host, port)
}
