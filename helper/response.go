package helper

type DataResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseFormat(code int, msg string, data interface{}) (int, interface{}) {
	res := &DataResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}

	return code, res
}
