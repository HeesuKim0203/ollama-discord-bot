package util

import "time"

type Response struct {
	model     string    `json:"model"`
	createdAt time.Time `json:"created_at"`
	data      string    `json:"response"`
	done      bool      `json:"done"`
}

func NewResponse(model string, createdAt time.Time, data string, done bool) *Response {
	return &Response{
		model:     model,
		createdAt: createdAt,
		data:      data,
		done:      done,
	}
}

func (r *Response) GetModel() string {
	return r.model
}

func (r *Response) GetCreateAt() time.Time {
	return r.createdAt
}

func (r *Response) GetData() string {
	return r.data
}

func (r *Response) GetDone() bool {
	return r.done
}
