package apisix

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sunreaver/tomlanalysis/timesize"
)

type Config struct {
	Addr    string            `toml:"addr" json:"addr" yaml:"addr"`
	ApiKey  string            `toml:"api_key" json:"api_key" yaml:"api_key"`
	Timeout timesize.Duration `toml:"timeout" json:"timeout" yaml:"timeout"`
}

func NewClient(cfg Config) *ApiSixClient {
	return &ApiSixClient{
		client: &http.Client{
			Timeout: cfg.Timeout.Duration(),
			Transport: &http.Transport{
				MaxIdleConns:        20,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     cfg.Timeout.Duration(),
				DisableKeepAlives:   false,
			},
		},
		cfg: cfg,
	}
}

type Sourcer interface {
	SourcePath() string
}

type ApiSixClient struct {
	client *http.Client
	cfg    Config
}

// 新建资源
func (c *ApiSixClient) Create(ctx context.Context, source Sourcer) (*Response, error) {
	return c.send(ctx, http.MethodPut, source)
}

// 更新资源
func (c *ApiSixClient) Update(ctx context.Context, source Sourcer) (*Response, error) {
	return c.send(ctx, http.MethodPatch, source)
}

// 删除资源
func (c *ApiSixClient) Delete(ctx context.Context, source Sourcer) (*Response, error) {
	return c.send(ctx, http.MethodDelete, source)
}

func (c *ApiSixClient) send(ctx context.Context, method string, source Sourcer) (*Response, error) {
	if source == nil || c == nil || c.client == nil {
		return nil, errors.New("invalid obj")
	}
	var err error
	var body []byte
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		body, err = json.Marshal(source)
		if err != nil {
			return nil, errors.Wrap(err, "source marshal")
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, c.cfg.Addr+source.SourcePath(), bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}
	req.Header.Set("X-API-KEY", c.cfg.ApiKey)
	respBody, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client do")
	}
	defer respBody.Body.Close()
	if http.StatusOK > respBody.StatusCode || respBody.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(respBody.Body)
		return nil, errors.Errorf("response status code: %d; body: %v", respBody.StatusCode, string(body))
	}
	var resp Response
	if err := json.NewDecoder(respBody.Body).Decode(&resp); err != nil {
		return nil, errors.Wrap(err, "response decode")
	} else if resp.Code != 0 {
		return &resp, errors.Errorf("response code: %d, message: %s", resp.Code, resp.Message)
	}
	return &resp, nil
}
