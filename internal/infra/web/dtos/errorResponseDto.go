package web

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseDto struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, code int, errMsg string) {
	w.WriteHeader(code)
	errorResponse := ErrorResponseDto{
		Code:  code,
		Error: errMsg,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
