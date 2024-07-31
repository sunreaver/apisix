package traffic

/*
{
	"rate": 10,
	"burst": 20,
	"key_type": "var",
	"key": "remote_addr",
	"rejected_code": 503,
	"rejected_msg": "503",
	"nodelay": false,
	"allow_degradation": false
}
*/

type KeyTypeEnum string

const (
	KeyTypeVar    KeyTypeEnum = "var"
	KeyTypeVarCom KeyTypeEnum = "var_combination"
)

type KeyEnum string

/*
["remote_addr", "server_addr", "http_x_real_ip", "http_x_forwarded_for", "consumer_name"]
*/
const (
	KeyRemoteAddr   KeyEnum = "remote_addr"
	KeyServerAddr   KeyEnum = "server_addr"
	KeyXRealIP      KeyEnum = "http_x_real_ip"
	KeyForwardedFor KeyEnum = "http_x_forwarded_for"
	KeyConsumerName KeyEnum = "consumer_name"
)

// limit-req 插件使用漏桶算法限制单个客户端对服务的请求速率。
type LimitReq struct {
	Rate  uint `json:"rate,omitempty"`  // 指定的请求速率（以秒为单位），请求速率超过 rate 但没有超过（rate + burst）的请求会被延时处理。
	Burst uint `json:"burst,omitempty"` // 请求速率超过（rate + burst）的请求会被直接拒绝。

	KeyType KeyTypeEnum `json:"key_type,omitempty"` // 要使用的用户指定 key 的类型。

	KeyName          KeyEnum `json:"key,omitempty"`               // 用来做请求计数的依据，当前接受的 key 有：remote_addr（客户端 IP 地址），server_addr（服务端 IP 地址）, 请求头中的 X-Forwarded-For 或 X-Real-IP，consumer_name（Consumer 的 username）。
	RejectedCode     int     `json:"rejected_code,omitempty"`     // 当超过阈值的请求被拒绝时，返回的 HTTP 状态码。
	RejectedMsg      string  `json:"rejected_msg,omitempty"`      // 当超过阈值的请求被拒绝时，返回的响应体。
	NoDelay          bool    `json:"nodelay,omitempty"`           // 当设置为 true 时，请求速率超过 rate 但没有超过（rate + burst）的请求不会加上延迟；当设置为 false，则会加上延迟。
	AllowDegradation bool    `json:"allow_degradation,omitempty"` // 当设置为 true 时，如果限速插件功能临时不可用，将会自动允许请求继续。
}

func (_ *LimitReq) Key() string {
	return "limit-req"
}
