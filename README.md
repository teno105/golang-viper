# golang-apis

`golang-apis` 는 Golang으로 작성된 json 파일들을 gin으로 받아오는 프로젝트입니다.<br/>
게임이 실행될 때, 접속할 서버정보와 공지사항, 개인정보 정책 등을 얻어옵니다.

## 프로젝트 폴더 구조
```plaintext
golang-apis/
│
├── cmd/
│   └── golang-apis/
│        └── main.go
│
├── infra/
│   ├── db.go
│   └── file.go
│
├── internal/
│   └── models/
│        ├── in_game_board.go
│        ├── latest_policy.go
│        ├── maintenance.go
│        ├── notice.go
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
db.go : gorm 을 통해서 DB에 있는 데이터를 가져오는 함수가 정의 되어있습니다. (구현 예정)

### internal/models 폴더 설명
GameData 가 가지는 Data를 struct 형태로 정의되어 있습니다.

### data 폴더 설명
각 GameId 별로 게임서비스에 필요한 Data를 json 파일형태로 보관합니다.

## 주요 기능

### file.go
```go
func getProjectRoot() string {
	dir, err := os.Getwd()  // 현재 폴더 경로 확인 (/전체경로/cmd/golang-apis)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil { // 경로상에 go.mod 파일이 존재하는지 확인
			return dir
		}
		parent := filepath.Dir(dir) // 부모 디렉토리로 경로 변경
		if parent == dir {
			break
		}
		dir = parent
	}

	log.Fatal("go.mod 파일을 찾을 수 없습니다.")
	return ""
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
