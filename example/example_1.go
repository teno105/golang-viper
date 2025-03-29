package example

import (
	"fmt"

	"github.com/spf13/viper"
)

func Example1() {
	viper.SetDefault("app.name", "TenoApp")
	appName := viper.GetString("app.name")
	fmt.Println("App Name:", appName)
}
