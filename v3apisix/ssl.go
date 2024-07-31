package v3apisix

/*
	{
	    "id": "1",          # id
	    "cert": "cert",     # 证书
	    "key": "key",       # 私钥
	    "snis": ["t.com"]   # HTTPS 握手时客户端发送的 SNI
	}
*/

// SSL 证书管理
type SSL struct {
	ID   string   `json:"id"`   // 证书ID
	Cert string   `json:"cert"` // 证书内容
	Key  string   `json:"key"`  // 私钥内容
	SNIS []string `json:"snis"` // HTTPS 握手时客户端发送的 SNI
}

func (s *SSL) SourcePath() string {
	return "/apisix/admin/ssls/" + s.ID
}
