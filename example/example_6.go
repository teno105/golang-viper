package example

import (
	"fmt"
	"golang-viper/internal/models"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 절대 경로의 파일을 읽는 함수
func loadHandler(path string, target interface{}) error {
	viper.SetConfigFile(path)

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("파일이 존재하지 않음: %s %v", path, err)
		return err
	}

	// 설정 파일을 구조체로 변환
	if err := viper.Unmarshal(target); err != nil {
		return err
	}

	return nil
}

func Example6() {

	type Config struct {
		StoreLink models.StoreLink `mapstructure:"store_link" json:"store_link"`
	}

	var storeLink Config

	// version_infos.yml 읽기
	if err := loadFile("data/12/store_link.yml", &storeLink); err != nil {
		log.Printf("[WARN] version_infos.yml 파일로딩을 실패: %v", err)
	}

	fmt.Printf("Platform: %s\n", storeLink.StoreLink.PlatformType)
	fmt.Printf("StoreUrl: %s\n", storeLink.StoreLink.StoreUrl)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 변경된 파일의 경로를 출력
		fmt.Println("Config file changed:", e.Name)

		// e.Name이 절대경로이기 때문에 절대 경로로 파일을 읽어야 함
		if err := loadHandler(e.Name, &storeLink); err != nil {
			log.Printf("[WARN] watchConfig 파일로딩을 실패: %v", err)
		}
	})

	// 파일이 변경이 되었는지 확인하기 위해 gin 을 사용함
	r := gin.Default()

	r.GET("/info", func(c *gin.Context) {

		// viper에서 설정 값 가져오기
		c.JSON(http.StatusOK, storeLink)
	})
	r.Run(":9095")
}
