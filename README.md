# proxyfinderclient

## 사용법
```
import (
    "github.com/rlarlgh/proxyfinderclient"
)

func main() {
    client := proxyfinderclient.NewProxyFinderClient("localhost", 50053)
    // host, port, err := client.FindProxyServerStrict("https://google.co.kr")
	host, port, err := client.FindProxyServer()
    if err != nil {
        log.Println("** proxy server not found:", err)
    }
    proxyServer := fmt.Sprintf("%s:%d", host, port)
    log.Println("** 프록시 서버를 찾았습니다:", proxyServer)
}
```

