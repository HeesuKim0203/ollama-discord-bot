package util

type Request struct {
	Message string `json:"message"`
	Model   string `json:"model"`
}

func NewRequst(message string, model string) *Request {
	return &Request{
		Message: message,
		Model:   model,
	}
}
