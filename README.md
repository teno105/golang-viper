# golang-viper

`go-viper` 는 Go 애플리케이션에서 설정을 쉽게 관리할 수 있도록 도와주는 라이브러리입니다.<br/>
JSON, YAML, TOML, ENV 등 다양한 설정 파일 형식을 지원하며, 환경 변수, 플래그, 기본값 설정 등도 함께 사용할 수 있습니다.

## 프로젝트 폴더
### 1. 폴더 구조
```plaintext
golang-viper/
│
├── cmd/
│   └── golang-viper/
│        └── main.go
│
├── infra/
│   └── file.go
│
├── data/
│   ├── read_toml.toml
│   └── 12/
│        ├── maintenance.yml
│        ├── store_link.yml
│        └── version_infos.yml
│
├── example/
│   ├── example_1.go
│   ├── example_2.go
│   ├── example_3.go
│   ├── example_4.go
│   ├── example_5.go
│   └── example_6.go
│
├── internal/
│   └── models/
│        └── result_data.go
│
├── go.mod
├── Makefile
└── README.md
```

### 2. 폴더 설명

| 폴더 | 설명 |
| --- | --- |
| `infra` | RootPath를 가져오는 함수가 정의 |
| `data` | 예제에 필요한 Data를 toml, yaml 파일형태로 보관 |
| `example` | 예제 소스 모음 |
| `internal/models` | GameData 가 가지는 Data를 struct 형태로 정의 |

# 예제
## 1. 기본값 설정 (viper.SetDefault)
기본값을 설정하면 설정 파일이나 환경 변수에서 값을 찾을 수 없을 때 사용됩니다.

```go
// example_1.go
func Example1() {
	viper.SetDefault("app.name", "TenoApp")
	appName := viper.GetString("app.name")
	fmt.Println("App Name:", appName)
}
```
### 결과화면
```
App Name: TenoApp
```

## 2. 설정 파일 로드 (viper.ReadInConfig)
Viper를 사용하여 JSON, TOML, YAML, HCL 설정 파일을 로드할 수 있습니다.
이번 예제에는 TOML 로드합니다.<br/>

### data/read_toml.toml
```toml
[server]
port = 8080
mode = "release"

[database]
user = "root"
password = "password123"
host = "localhost"
port = 3306
```

```go
// example_2.go
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
	fmt.Printf("Server Port: %d\n", port)
}
```

`.` Viper는 구분된 키 경로를 전달하여 중첩된 필드에 액세스할 수 있습니다 .
```go
GetString("database.user")	// (returns "root")
```
### 결과화면
```
DB User: root, Password: password123
Server Port: 8080
```

## 3. 환경 변수 사용 (viper.AutomaticEnv)
Viper는 환경 변수도 지원합니다.

```go
// example_3.go
func Example3() {
	// 환경 변수 사용 설정
	viper.AutomaticEnv()

	// 환경 변수에서 특정 값 읽기
	os.Setenv("APP_PORT", "8080") // 예제용 환경 변수 설정
	port := viper.GetInt("APP_PORT")

	fmt.Printf("App Port: %d\n", port)
}
```
### 결과화면
```
App Port: 8080
```


## 4. 플래그(Flag)와 함께 사용 (viper.BindPFlags)
Viper는 플래그에 바인딩하는 기능이 있습니다.<br/>
BindEnv처럼, 바인딩 메서드가 호출될 때 값이 설정되지 않고, 액세스할 때 설정됩니다. 즉, init()함수에서도 원하는 만큼 일찍 바인딩할 수 있습니다.<br/>

```go
// example_4.go
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
```
### 결과화면
```
go run .\cmd\golang-viper\main.go --flagname 111
Flag value: 111
```

## 5. Unmarshal 사용 (viper.Unmarshal)
viper.Unmarshal은 Viper에서 로드한 설정 데이터를 구조체로 변환하는 기능을 제공합니다.<br/>
이를 수행하는 방법은 두 가지가 있습니다.
```go
Unmarshal(rawVal any) : error
UnmarshalKey(key string, rawVal any) : error	// 특정 키에 대해서만 변환
```
```go
// example_5.go
package example

import (
	"fmt"
	"golang-viper/infra"
	"golang-viper/internal/models"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// YAML 파일을 읽어 구조체로 변환하는 함수
func loadFile(filePath string, target interface{}) error {
	// 실행 중인 프로그램의 경로 가져오기
	rootPath := infra.GetProjectRoot()
	if rootPath == "" {
		return fmt.Errorf("작업 디렉토리를 가져오는 중 오류 발생 %s", filePath)
	}

	path := filepath.Join(rootPath, filePath)
	viper.SetConfigFile(path)

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("파일이 존재하지 않음: %s %v", filePath, err)
		return err
	}

	// 설정 파일을 구조체로 변환
	if err := viper.Unmarshal(target); err != nil {
		return err
	}

	return nil
}

func Example5() {

	var resultData models.ResultData

	// version_infos.yml 읽기
	if err := loadFile("data/12/version_infos.yml", &resultData); err != nil {
		log.Printf("[WARN] version_infos.yml 파일로딩을 실패: %v", err)
	}

	// store_link.yml 읽기
	if err := loadFile("data/12/store_link.yml", &resultData); err != nil {
		log.Printf("[WARN] store_link.yml 파일로딩을 실패: %v", err)
	}

	// maintenance.yml 읽기
	if err := loadFile("data/12/maintenance.yml", &resultData); err != nil {
		log.Printf("[WARN] maintenance.yml 파일로딩을 실패: %v", err)
	}

	fmt.Printf("Version StoreType: %s, VersionNo: %s\n", resultData.VersionInfos.VersionInfo.StoreType, resultData.VersionInfos.VersionInfo.VersionNo)
	fmt.Printf("Store PlatformType: %s, Url: %s\n", resultData.StoreLink.PlatformType, resultData.StoreLink.StoreUrl)
	fmt.Printf("Maintenance Message: %s\n", resultData.Maintenance.Message[0].Body)
}
```
```go
// result_data.go
package models

// ResultData JSON 구조체
type ResultData struct {
	VersionInfos VersionInfos `mapstructure:"version_infos" json:"version_infos"`
	Maintenance  Maintenance  `mapstructure:"maintenance" json:"maintenance"`
	StoreLink    StoreLink    `mapstructure:"store_link" json:"store_link"`
}

// VersionInfos JSON 구조체
type VersionInfos struct {
	VersionInfo VersionInfo `mapstructure:"version_info" json:"version_info"`
}

type VersionInfo struct {
	StoreType      string `mapstructure:"store_type" json:"store_type"`
	VersionNo      string `mapstructure:"version_no" json:"version_no"`
	// ...
}

// Maintenance JSON 구조체
type Maintenance struct {
	Message []MaintenanceMessage `mapstructure:"message" json:"message"`
}

type MaintenanceMessage struct {
	Body        string `mapstructure:"body" json:"body"`
	// ...
}

// StoreLink JSON 구조체
type StoreLink struct {
	PlatformType string `mapstructure:"platform_type" json:"platform_type"`
	StoreUrl     string `mapstructure:"store_url" json:"store_url"`
}
```


### 결과화면
```
Version StoreType: Android, VersionNo: 1.3.55
Store PlatformType: Android, Url: https://play.google.com/store/apps/dev?id=6650600561956737529
Maintenance Message: 점검 테스트 내용
```

## 6. 실시간 감시(Hot Reload) (viper.WatchConfig)
Viper는 애플리케이션이 실행 중에 구성 파일을 실시간으로 읽을 수 있는 기능을 지원합니다.<br/>

```go
// example_6.go
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
```
### 결과화면
http://localhost:9095/info
를 호출
```
# 변경전
{
  "store_link": {
    "platform_type": "Android",
    "store_url": "https://play.google.com/store/apps/dev?id=6650600561956737529"
  }
}
```
```
# 변경후
{
  "store_link": {
    "platform_type": "iOS",
    "store_url": "https://play.google.com/store/apps/dev?id=6650600561956737529"
  }
}
```

# Viper 사용시 주의할 점

### 1. Viper의 설정 우선순위 이해하기

Viper는 여러 가지 방법으로 설정을 로드할 수 있으며, 다음과 같은 우선순위를 가집니다.

1. Explcit Set: viper.Set()으로 직접 설정한 값
2. Flag: pflag를 사용하여 설정한 값
3. Environment Variables: 환경 변수 값
4. Configuration File: 설정 파일에서 읽은 값
5. Default: viper.SetDefault()로 설정한 기본값

### 2. 환경 변수 사용 시 키 대소문자 주의

Viper는 기본적으로 환경 변수 키를 대문자로 변환하여 처리합니다. <br/>
환경 변수를 사용할 때는 viper.AutomaticEnv()를 호출한 후, viper.SetEnvKeyReplacer()를 사용하여 예상한 키값을 정확히 매핑해야 합니다.
```go
viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
```
이렇게 하면 config.db.host와 같은 키가 CONFIG_DB_HOST 환경 변수와 매칭됩니다.

### 3. 설정 파일 경로 및 포맷 확인

Viper는 설정 파일을 자동으로 탐색하지만, 파일을 찾지 못하면 오류 없이 진행되므로 예상과 다르게 동작할 수 있습니다. <br/>
viper.ReadInConfig() 호출 후, viper.ConfigFileUsed()로 실제 로드된 파일을 확인하는 것이 좋습니다.
```go
if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("설정 파일을 읽는 데 실패했습니다: %v", err)
}
fmt.Println("사용된 설정 파일:", viper.ConfigFileUsed())
```
또한, 설정 파일 포맷(JSON, YAML, TOML 등)이 올바르게 지정되었는지도 확인해야 합니다.

### 4. 기본값 설정 후 오버라이드 확인

viper.SetDefault()를 사용하여 기본값을 설정할 수 있지만, 설정 파일이나 환경 변수가 존재하면 덮어씌워집니다. <br/>
값이 정상적으로 반영되는지 확인하려면 viper.Get()로 가져온 값을 출력해 보세요.
```go
viper.SetDefault("server.port", 8080)
fmt.Println("Server Port:", viper.GetInt("server.port"))
```
### 5. 자동 리로드 기능 사용 시 주의

Viper는 viper.WatchConfig()를 사용하여 설정 파일 변경을 감지하고 자동으로 반영할 수 있습니다. <br/>
하지만, 설정이 변경될 때 이를 처리하는 핸들러를 반드시 등록해야 합니다.
```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("설정 파일이 변경되었습니다:", e.Name)
})
```
이 기능을 사용할 때, 설정 파일이 예상치 못하게 변경될 경우 프로그램이 예기치 않게 동작할 수 있으므로 주의가 필요합니다.

### 6. 중첩된 구조체와 Unmarshal 사용

Viper는 Unmarshal()을 사용하여 구조체로 매핑할 수 있습니다. 하지만, 구조체 태그(mapstructure 태그)를 정확하게 설정하지 않으면 올바르게 매핑되지 않습니다.
```go
type Config struct {
    Server struct {
        Port int `mapstructure:"port"`
    } `mapstructure:"server"`
}

var config Config
if err := viper.Unmarshal(&config); err != nil {
    log.Fatalf("구성 파일을 구조체로 변환하는 데 실패했습니다: %v", err)
}
```

### 7. 동적 키 사용 시 주의

Viper에서 viper.GetStringMap()을 사용하여 동적으로 맵 데이터를 가져올 때, <br/>
반환된 값이 map[string]interface{} 타입이므로 타입 변환에 유의해야 합니다.

```go
configMap := viper.GetStringMap("database")
for key, value := range configMap {
    fmt.Printf("Key: %s, Value: %v\n", key, value)
}
```
