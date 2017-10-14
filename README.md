# daumcrawler

Crawls on Daum search engine on various types of contents.

## Installation
```bash
go get -u github.com/hyunchel/daumcrawler
```

## Usage
```go
import "github.com/hyunchel/daumcrawler"

var (
    kakaoappkey = "some kakao app key"
    ...
)

func main() {
    keyword := "hello!"
    daumcrawler.Run(kakaoappkey, keyword)
}
