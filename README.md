# uaparse
ua parse

```go
const (
	tests = "Mozilla/5.0 (Linux; Android 7.1.2; vivo X9 Build/N2G47H; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044506 Mobile Safari/537.36 MMWEBID/134 MicroMessenger/7.0.3.1400(0x2700033C) Process/tools NetType/4G Language/zh_CN"
)
func main() {
	println(ua.ParseDevice(tests).Brand != "vivo")
	println(ua.ParseUA(tests).Family != "Safari")
}
```
