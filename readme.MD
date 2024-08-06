# APISIX Golang apis/sdk

## 使用示例

```golang
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/sunreaver/apisix"
    "github.com/sunreaver/apisix/v3apisix"
    "github.com/sunreaver/apisix/v3apisix/plugins"
    "github.com/sunreaver/apisix/v3apisix/plugins/traffic"
    "github.com/sunreaver/apisix/v3apisix/plugins/transformation"
    "github.com/sunreaver/tomlanalysis/timesize"
)

func main() {
    // TODO: Implement your test cases here.
    upstream := v3apisix.Upstream{
        ID:   "my_test_" + "1234567890",
        Name: "a",
        Nodes: []v3apisix.UpstreamNode{
            {
                Host:   "127.0.0.1",
                Port:   80,
                Weight: 100,
            },
            {
                Host:   "127.0.0.1",
                Port:   81,
                Weight: 100,
            },
        },
        Scheme:       "http",
        Retries:      3,
        RetryTimeout: 20,
        KeepalivePool: &v3apisix.UpstreamKeepalivePool{
            IdleTimeout: 30,
        },
    }
    s := v3apisix.Service{
        ID:              "my_test_" + "1234567890",
        Name:            "a",
        Desc:            "test-a-1",
        Upstream:        upstream,
        EnableWebsocket: false,
    }

    client := apisix.NewClient(apisix.Config{
        Addr:    "https://apisix-api.com",
        ApiKey:  "edd1c9f034335f136f87ad84b625c8f1",
        Timeout: timesize.Duration(time.Second * 5),
    })

    resp, err := client.Create(context.TODO(), &s)
    if err != nil {
        fmt.Println(err)
        panic(err.Error())
    } else {
        fmt.Printf("%+v\n", resp)
    }

    t := v3apisix.Route{
        ID:   "my_test_" + "1234567890",
        Name: "a",
        Uri:  "/mypath/bac/*",
        Host: "myrouter.com",
        Plugins: plugins.Plugins{
            &traffic.LimitReq{
                Rate:    10,
                Burst:   2,
                KeyName: "remote_addr",
            },
            &transformation.ProxyRewrite{
                Scheme: "http",
                RegexUri: []string{
                    "/mypath/bac/(.*)", "/bac/$1",
                },
            },
        },
        Status:    0,
        ServiceId: resp.ID(),
    }

    resp, err = client.Create(context.TODO(), &t)
    if err != nil {
        fmt.Println(err)
        panic(err.Error())
    } else {
        fmt.Printf("%+v\n", resp)
    }

    resp, err = client.Delete(context.TODO(), &t)
    if err != nil {
        fmt.Println(err)
        panic(err.Error())
    } else {
        fmt.Printf("%+v\n", resp)
    }
    resp, err = client.Delete(context.TODO(), &s)
    if err != nil {
        fmt.Println(err)
        panic(err.Error())
    } else {
        fmt.Printf("%+v\n", resp)
    }
}
```