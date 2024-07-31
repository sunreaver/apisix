package general

/*
{
	"http_to_https": true,
	"uri": "$uri/index.html",
	"ret_code": 302,
	"encode_uri": false,
	"append_query_string": false
}
*/

// redirect 插件可用于配置 URI 重定向。
// http_to_https、uri 和 regex_uri 只能配置其中一个属性。
// http_to_https、和 append_query_string 只能配置其中一个属性。
type Redirect struct {
	HttpToHttps       bool   `json:"http_to_https,omitempty"`       // 当设置为 true 并且请求是 HTTP 时，它将被重定向具有相同 URI 和 301 状态码的 HTTPS，原 URI 的查询字符串也将包含在 Location 头中。
	Uri               string `json:"uri"`                           // 要重定向到的 URI，可以包含 NGINX 变量。例如：/test/index.htm，$uri/index.html，${uri}/index.html，https://example.com/foo/bar。如果你引入了一个不存在的变量，它不会报错，而是将其视为一个空变量。
	RetCode           int    `json:"ret_code,omitempty"`            // HTTP 响应码; 默认302
	EncodeUri         bool   `json:"encode_uri,omitempty"`          // 当设置为 true 时，对返回的 Location Header 按照 RFC3986 的编码格式进行编码。
	AppendQueryString bool   `json:"append_query_string,omitempty"` // 当设置为 true 时，将原始请求中的查询字符串添加到 Location Header。如果已配置 uri 或 regex_uri 已经包含查询字符串，则请求中的查询字符串将附加一个&。如果你已经处理过查询字符串（例如，使用 NGINX 变量 $request_uri），请不要再使用该参数以避免重复。
}

func (_ *Redirect) Key() string {
	return "redirect"
}
