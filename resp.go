package apisix

import "path"

type Response struct {
	Key  string         `json:"key"`
	Data map[string]any `json:"value"`
}

func (r Response) ID() string {
	id := path.Base(r.Key)
	if id == "." || id == "/" {
		return ""
	}
	return id
}
