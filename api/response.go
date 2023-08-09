package api

type Response struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
}

func NewResponse(statusCode int, data interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Data:       data,
	}
}
