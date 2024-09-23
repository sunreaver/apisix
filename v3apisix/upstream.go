package v3apisix

type LBType string // roundrobin,chash,ewma,least_conn

const (
	LBTypeRoundRobin LBType = "roundrobin"
	LBTypeChash      LBType = "chash"
	LBTypeEWMA       LBType = "ewma"
	LBTypeLeastConn  LBType = "least_conn"
)

type DiscoveryType string // eureka,nacos,consul,dns

const (
	DiscoveryTypeDNS    DiscoveryType = "dns"
	DiscoveryTypeEureka DiscoveryType = "eureka"
	DiscoveryTypeNacos  DiscoveryType = "nacos"
	DiscoveryTypeConsul DiscoveryType = "consul"
)

type UpstreamNode struct {
	Host   string `json:"host,omitempty"`   // 地址，可以是 IP 或域名
	Port   int    `json:"port,omitempty"`   // 端口
	Weight int    `json:"weight,omitempty"` // 权重
}

type UpstreamTimeout struct {
	Connect float32 `json:"connect,omitempty"` // 建立连接的超时时间，单位秒。
	Send    float32 `json:"send,omitempty"`    // 向后端发送请求的超时时间，单位秒。
	Read    float32 `json:"read,omitempty"`    // 从后端读取响应的超时时间，单位秒。
}

type UpstreamTLS struct {
	ClientCertID string `json:"client_cert_id,omitempty"` // 客户端证书ID，来源域名配置
}

type UpstreamKeepalivePool struct {
	Size        int `json:"size,omitempty"`         // 连接池大小，默认 320
	IdleTimeout int `json:"idle_timeout,omitempty"` // 空闲连接超时时间，单位秒，默认 60s
	Requests    int `json:"requests,omitempty"`     // 最大请求数，默认 1000
}

// 负载均衡上游.
// https://apisix.apache.org/zh/docs/apisix/3.8/admin-api/#upstream-uri
type Upstream struct {
	ID            string                 `json:"id,omitempty"`             // 负载均衡上游的唯一标识符。
	Name          string                 `json:"name,omitempty"`           // 上游名称。
	Desc          string                 `json:"desc,omitempty"`           // 上游描述。
	Type          LBType                 `json:"type,omitempty"`           // 负载均衡算法，默认值是roundrobin。
	Nodes         []UpstreamNode         `json:"nodes,omitempty"`          // [ {"host":"host", "port":80, "weight": 100} ]
	ServiceName   string                 `json:"service_name,omitempty"`   // 服务发现时使用的服务名
	DiscoveryType DiscoveryType          `json:"discovery_type,omitempty"` // 服务发现类型,与 service_name 配合使用。
	ChashKey      string                 `json:"key,omitempty"`            // 该选项只有类型是 chash 才有效。根据 key 来查找对应的节点 id，相同的 key 在同一个对象中，则返回相同 id。目前支持的 NGINX 内置变量有 uri, server_name, server_addr, request_uri, remote_port, remote_addr, query_string, host, hostname, arg_***，其中 arg_*** 是来自 URL 的请求参数
	HashOn        string                 `json:"hash_on,omitempty"`        // 用于 hash 的字段，hash_on 支持的类型有 vars（NGINX 内置变量），header（自定义 header），cookie，consumer，默认值为 vars。
	Checks        any                    `json:"checks,omitempty"`         // 健康检查配置，目前支持的选项有 active 和 passive。
	Retries       int                    `json:"retries"`                  // 使用 NGINX 重试机制将请求传递给下一个上游，默认启用重试机制且次数为后端可用的节点数量。如果指定了具体重试次数，它将覆盖默认值。当设置为 0 时，表示不启用重试机制。
	RetryTimeout  int                    `json:"retry_timeout"`            // 限制是否继续重试的时间，若之前的请求和重试请求花费太多时间就不再继续重试。当设置为 0 时，表示不启用重试超时机制。
	Timtout       *UpstreamTimeout       `json:"timeout,omitempty"`        // 设置连接、发送消息、接收消息的超时时间，以秒为单位。
	PassHost      string                 `json:"pass_host,omitempty"`      // 请求发给上游时的 host 设置选型。 [pass，node，rewrite] 之一，默认是 pass。pass: 将客户端的 host 透传给上游； node: 使用 upstream node 中配置的 host； rewrite: 使用配置项 upstream_host 的值。
	UpstreamHost  string                 `json:"upstream_host,omitempty"`  // 指定上游请求的 host，只在 pass_host 配置为 rewrite 时有效。
	Scheme        string                 `json:"scheme,omitempty"`         // 请求上游的协议，对于 7 层代理，可选值为 [http, https, grpc, grpcs]。对于 4 层代理，可选值为 [tcp, udp, tls]。默认值为 http
	TLS           *UpstreamTLS           `json:"tls,omitempty"`            // TLS 配置
	KeepalivePool *UpstreamKeepalivePool `json:"keepalive_pool,omitempty"` // 连接池配置
}

func (u *Upstream) SourcePath() string {
	return "/apisix/admin/upstreams/" + u.ID
}
