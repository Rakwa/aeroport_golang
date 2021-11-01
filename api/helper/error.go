package helper

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

/*
	Return error in json
*/
func GetError(err error, writer http.ResponseWriter, statusCode int) {

	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   statusCode,
	}

	message, _ := json.Marshal(response)
	writer.WriteHeader(response.StatusCode)
	writer.Write(message)
}
