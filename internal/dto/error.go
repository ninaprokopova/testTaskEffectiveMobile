package dto

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
}

func NewErrorResponse(code, message string) (resp ErrorResponse) {
	resp = ErrorResponse{}
	resp.Error.Code = code
	resp.Error.Message = message
	return resp
}
