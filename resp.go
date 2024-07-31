package apisix

type Response struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      map[string]any `json:"data"`
	RequestID string         `json:"request_id"`
}

func (r Response) ID() string {
	if r.Data == nil {
		return ""
	}
	idAny, ok := r.Data["id"]
	if !ok {
		return ""
	}
	id, _ := idAny.(string)
	return id
}
