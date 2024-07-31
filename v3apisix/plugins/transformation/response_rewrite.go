package plugins

import "github.com/sunreaver/apisix/v3apisix/vars"

/*
{
	"status_code": 400,
	"body": "{}",
	"body_base64": false,
	"headers": {
		"Content-Type": "application/json",
		"X-Custom-Header": "custom value"
	},
	"vars": [
		["status_code", "==", "200"]
	]
}
*/

// response-rewrite 插件支持修改上游服务或 APISIX 返回的 Body 和 Header 信息。
type ResponseRewrite struct {
	StatusCode int               `json:"status_code,omitempty"`
	Body       string            `json:"body,omitempty"`
	Base64     bool              `json:"body_base64,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Vars       vars.Vars         `json:"vars,omitempty"`
}

func (_ *ResponseRewrite) Key() string {
	return "response-rewrite"
}
