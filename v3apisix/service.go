package v3apisix

type Service struct {
	ID              string   `json:"id,omitempty"`   // service ID。必备字段，Create/Update/Delete 时需要指定。
	Name            string   `json:"name,omitempty"` // 服务名称。
	Desc            string   `json:"desc,omitempty"` // 服务描述。
	EnableWebsocket bool     `json:"enable_websocket,omitempty"`
	Upstream        Upstream `json:"upstream"`
}

func (s Service) SourcePath() string {
	return "/apisix/admin/services/" + s.ID
}
