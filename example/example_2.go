package example

import (
	"fmt"
	"golang-viper/infra"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

func Example2() {
	// 실행 중인 프로그램의 경로 가져오기
	rootPath := infra.GetProjectRoot()

	viper.SetConfigName("read_toml")                     // 설정 파일 이름: "read_toml"
	viper.SetConfigType("toml")                          // Config's format: "json", "toml", "yaml", "yml"
	viper.AddConfigPath(filepath.Join(rootPath, "data")) // data 디렉토리에서 설정 파일 찾기

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("설정 파일 읽기 실패: %v", err)
		return
	}

	// 값 가져오기
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	port := viper.GetInt("server.port")

	fmt.Printf("DB User: %s, Password: %s\n", user, password)
	fmt.Printf("Server Post: %d\n", port)
}
