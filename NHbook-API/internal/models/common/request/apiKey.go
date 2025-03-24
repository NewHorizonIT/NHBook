package request

type ApiKeyRequest struct {
	Key    string `json:"key"`
	Status uint   `json:"status"`
}
