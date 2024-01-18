package util

import "time"

type Response struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Data      string    `json:"response"`
	Done      bool      `json:"done"`
}

func NewResponse(model string, createdAt time.Time, data string, done bool) *Response {
	return &Response{
		Model:     model,
		CreatedAt: createdAt,
		Data:      data,
		Done:      done,
	}
}

// func (r *Response) GetModel() string {
// 	return r.model
// }

// func (r *Response) GetCreateAt() time.Time {
// 	return r.createdAt
// }

// func (r *Response) GetData() string {
// 	return r.data
// }

// func (r *Response) GetDone() bool {
// 	return r.done
// }
