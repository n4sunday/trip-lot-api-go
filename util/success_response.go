package util

type successResponse struct {
	Data interface{} `json:"data"`
}

type successMessage struct {
	Message interface{} `json:"message"`
}

func NewSuccessResponse(data interface{}) successResponse {
	return successResponse{
		Data: data,
	}
}

func NewSuccessMessage(message interface{}) successMessage {
	return successMessage{
		Message: message,
	}
}
