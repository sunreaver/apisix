package transformation

/*
	{
	    "scheme": "https",
	    "regex_uri": [
	        "/(.*)",
	        "/$1"
	    ],
	    "host": "www.aops.com",
	    "method": "COPY",
	    "headers": {
	        "content-type1": "application/json",
	        "content-type2": "application/json"
	    }
	}
*/

// ProxyRewrite 请求改写
type ProxyRewrite struct {
	Scheme   string         `json:"scheme,omitempty"`    // 协议改写
	RegexUri []string       `json:"regex_uri,omitempty"` // 路径改写->正则改写
	Host     string         `json:"host,omitempty"`      // 域名改写
	Method   string         `json:"method,omitempty"`    // Method改写
	Headers  map[string]any `json:"headers,omitempty"`   // 请求头改写
}

func (_ *ProxyRewrite) Key() string {
	return "proxy-rewrite"
}
