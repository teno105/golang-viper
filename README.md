# golang-viper

`go-viper` 는 Go 애플리케이션에서 설정을 쉽게 관리할 수 있도록 도와주는 라이브러리입니다.<br/>
JSON, YAML, TOML, ENV 등 다양한 설정 파일 형식을 지원하며, 환경 변수, 플래그, 기본값 설정 등도 함께 사용할 수 있습니다.

## 프로젝트 폴더 구조
```plaintext
golang-viper/
│
├── cmd/
│   └── golang-viper/
│        └── main.go
│
├── example/
│   ├── ex_1.go
│   ├── xxxx.go
│   ├── xxxx.go
│   ├── xxxx.go
│   ├── xxxx.go
│   └── file.go
│
├── internal/
│   └── models/
│        ├── maintenance.go
│        ├── store_link.go
│        └── version_infos.go
│
├── data/
│   └── 11/
│        ├── in_game_board.json
│        ├── latest_policy.json
│        ├── maintenance.json
│        ├── notice.json
│        ├── store_link.json
│        └── version_infos.go
│
├── go.mod
├── Makefile
└── README.md
```

### infra 폴더 설명
file.go : /Data 에서 struct GameData 의 Type에 맞는 데이터를 가져오는 함수가 정의 되어있습니다.<br/>

### internal/models 폴더 설명
GameData 가 가지는 Data를 struct 형태로 정의되어 있습니다.

### data 폴더 설명
각 GameId 별로 게임서비스에 필요한 Data를 json 파일형태로 보관합니다.

## 주요 기능

### ex_1.go
```go
import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 설정 파일 이름과 확장자 지정
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")   // 파일 타입 지정
	viper.AddConfigPath(".")      // 현재 디렉토리에서 설정 파일 찾기

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("설정 파일 읽기 실패: %v", err)
	}

	// 값 가져오기
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")

	fmt.Printf("DB Host: %s, Port: %d\n", host, port)
}
```
go.mod 파일을 기준으로 프로젝트 RootPath를 찾습니다.<br/>
이렇게 해야 지정된 파일을 어느 환경에서도 파일을 접근할 수 있습니다.
```go
// GameData 구조체
type GameData struct {
	InGameBoard  models.InGameBoard  `json:"in_game_board"`
	LatestPolicy models.LatestPolicy `json:"latest_policy"`
	VersionInfos models.VersionInfos `json:"version_infos"`
	Notice       models.Notice       `json:"notice"`
	Maintenance  models.Maintenance  `json:"maintenance"`
	StoreLink    models.StoreLink    `json:"store_link"`
}
```
```go
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
```
JSON 파일을 읽어 구조체로 변환하는 함수입니다.<br/>
아래와 같이 특정 경로에 있는 파일을 gameData 객체로 값을 가져옵니다.
```go
// cmd/main.go
err := infra.LoadJSONFile("data/11/in_game_board", &gameData);
```

### main.go
```go
r := gin.Default()

// GET /v2/init_data/games/:id 엔드포인트
r.GET("/v2/init_data/games/:id", func(c *gin.Context) {
    gameID := c.Param("id")           // URL에서 game ID 가져오기
    data, err := loadGameData(gameID) // 해당 ID의 데이터를 로드
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Game data not found"})
        return
    }
    c.JSON(http.StatusOK, data)
})
```
gin을 사용해서 특정 게임id 를 기준으로 GameData를 가져오는 Get 핸들러를 추가합니다.

## 실행화면
실행을 하면 응답상태를 확인할 수 있습니다.
![Image](https://github.com/user-attachments/assets/abf50340-ec98-4afe-97ca-fca05239a38e)
