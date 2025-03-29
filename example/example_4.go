package example

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Example4() {
	// 커맨드라인 플래그 설정
	pflag.Int("flagname", 1234, "help message for flagname")
	pflag.Parse()

	// Viper에 플래그 바인딩
	viper.BindPFlags(pflag.CommandLine)

	// 값 가져오기
	i := viper.GetInt("flagname") // retrieve value from viper
	fmt.Println("Flag value:", i) // use the variable to avoid the error
}
