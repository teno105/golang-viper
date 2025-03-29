package example

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Example3() {
	// 환경 변수 사용 설정
	viper.AutomaticEnv()

	// 환경 변수에서 특정 값 읽기
	os.Setenv("APP_PORT", "8080") // 예제용 환경 변수 설정
	port := viper.GetInt("APP_PORT")

	fmt.Printf("App Port: %d\n", port)
}
