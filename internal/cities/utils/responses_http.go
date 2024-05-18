package utils

type ResponseJson struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BabResponse(msg string) ResponseJson {
	return ResponseJson{
		Message: msg,
		Success: false,
	}
}
func OkResponse(msg string) ResponseJson {
	return ResponseJson{
		Message: msg,
		Success: true,
		Data:    nil,
	}
}

func OkResponseData(msg string, data interface{}) ResponseJson {
	return ResponseJson{
		Message: msg,
		Success: true,
		Data:    data,
	}
}
