package handler

type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorResponse struct {
	Errors []ErrorResponseDetail `json:"errors"`
}

type ErrorResponseDetail struct {
	Message string `json:"message"`
}

func BuildSuccessResponse(res interface{}) SuccessResponse {
	return SuccessResponse{Data: res}
}

func BuildSuccessMessage(msg string) SuccessResponse {
	return SuccessResponse{Message: msg}
}

func BuildErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Errors: []ErrorResponseDetail{
			{
				Message: err.Error(),
			},
		},
	}
}
