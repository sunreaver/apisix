package security

/*
{
  "allow_origins": "*",
  "allow_methods": "*",
  "allow_headers": "*",
  "expose_headers": "*",
  "max_age": 5,
  "allow_credential": false,
  "allow_origins_by_regex": [],
  "allow_origins_by_metadata": []
}
*/

// cors 插件可以让你轻松地为服务端启用 CORS（Cross-Origin Resource Sharing，跨域资源共享）的返回头。
// https://apisix.apache.org/zh/docs/apisix/3.8/plugins/cors/
type CROS struct {
	AllowOrigins    string `json:"allow_origins,omitempty"`
	AllowMethods    string `json:"allow_methods,omitempty"`
	AllowHeaders    string `json:"allow_headers,omitempty"`
	ExposeHeaders   string `json:"expose_headers,omitempty"`
	MaxAge          int    `json:"max_age,omitempty"`
	AllowCredential bool   `json:"allow_credential,omitempty"`
}

func (_ *CROS) Key() string {
	return "cros"
}
