package v3apisix

import (
	"github.com/sunreaver/apisix/v3apisix/plugins"
	"github.com/sunreaver/apisix/v3apisix/vars"
)

type RouteEnable int8

const (
	RouteEnableTrue  = RouteEnable(1)
	RouteEnableFalse = RouteEnable(0)
)

type Route struct {
	ID      string          `json:"id"`                // 路由 ID。必备字段，Create/Update/Delete 时需要指定。
	Name    string          `json:"name,omitempty"`    // 路由名称。
	Desc    string          `json:"desc,omitempty"`    // 路由描述信息。
	Uri     string          `json:"uri"`               // 匹配规则，路径，支持正则和前缀匹配。
	Plugins plugins.Plugins `json:"plugins,omitempty"` // 插件列表，可以配置多个插件。
	// Script 与 Plugin 不兼容，并且 Script 优先执行 Script，这意味着配置 Script 后，Route 上配置的 Plugin 将不被执行。
	// 理论上，在 Script 中可以编写任意 Lua 代码，你也可以直接调用已有的插件以复用已有的代码。
	Script      string      `json:"script,omitempty"`
	ServiceId   int         `json:"service_id,omitempty"`   // 需要绑定的 Service id
	Host        string      `json:"host,omitempty"`         // 匹配规则，域名，比如 foo.com；也支持泛域名，比如 *.foo.com。
	RemoteAddrs []string    `json:"remote_addrs,omitempty"` // 匹配规则，表示允许有多个不同 IP 地址，符合其中任意一个即可。
	Methods     []string    `json:"methods,omitempty"`      // 匹配规则，请求方法，比如 GET、POST 等。
	Priority    int         `json:"priority,omitempty"`     // 匹配规则，优先级，值越小优先级越高。
	Vars        vars.Vars   `json:"vars,omitempty"`         // 匹配规则，支持通过请求头，请求参数、Cookie 进行路由匹配。
	Status      RouteEnable `json:"status,omitempty"`       // 路由状态，1 表示启用，0 表示禁用。
}

func (r Route) SourcePath() string {
	return "/apisix/admin/routes/" + r.ID
}
