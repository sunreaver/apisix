package plugins

import "encoding/json"

type Pluginer interface {
	Key() string
}

type Plugins []Pluginer

// 将数组形式转换为map形式
func (p Plugins) MarshalJSON() ([]byte, error) {
	var subsmap = make(map[string]any, len(p))
	for _, v := range p {
		subsmap[v.Key()] = v
	}
	return json.Marshal(subsmap)
}
