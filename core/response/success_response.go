package res

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(statusCode int, msg string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		StatusCode: statusCode,
		Message:    msg,
		Data:       data,
	}
}
