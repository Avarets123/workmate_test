package utils

import (
	"encoding/json"
	"net/http"
	"strings"
	"workmate/pkg/cerror"
)

func ValidateMethod(w http.ResponseWriter, req *http.Request, method string) bool {
	if !strings.EqualFold(req.Method, method) {
		e := cerror.New("METHOD_NOT_AVAILABLE", http.StatusMethodNotAllowed)
		e.ResHttp(w)
		return false
	}

	return true
}

func ValidateReq[T any](w http.ResponseWriter, req *http.Request, method string, reqData T) bool {

	err := json.NewDecoder(req.Body).Decode(reqData)
	if err != nil {
		e := cerror.New("Invalid json", http.StatusUnprocessableEntity)
		e.ResHttp(w)
		return false
	}

	return true
}
