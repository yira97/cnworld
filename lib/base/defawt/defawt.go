package defawt

type RespLayout struct {
	Status  string      `json:"status"`
	Content interface{} `json:"content"`
	Detail  string      `json:"detail"` // debug
	Token   string      `json:"token"`
}