package traffic

/*
{
	"header_name": "X-Request-Id",
	"include_in_response": true,
}
*/

type RequestID struct {
	HeaderName        string `json:"header_name,omitempty"` // unique ID 的请求头的名称。默认值：X-Request-Id
	IncludeInResponse bool   `json:"include_in_response"`   // 当设置为 true 时，将 unique ID 加入返回头。
}

// request-id 插件通过 APISIX 为每一个请求代理添加 unique ID 用于追踪 API 请求。
// 默认采用UUID格式
func (_ *RequestID) Key() string {
	return "request-id"
}
