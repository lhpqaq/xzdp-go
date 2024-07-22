package xzdp

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
type FailResponse struct {
	Success bool   `json:"success"`
	ErrmMsg string `json:"errorMsg"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}

func NewFailureResponse(message string) *FailResponse {
	return &FailResponse{
		Success: false,
		ErrmMsg: message,
	}
}
