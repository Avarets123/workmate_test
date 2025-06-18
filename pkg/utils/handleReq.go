package utils

import (
	"encoding/json"
	"net/http"
	"workmate/pkg/cerror"
)

func ValidateReq[T any](w http.ResponseWriter, req *http.Request, method string, reqData T) bool {

	err := json.NewDecoder(req.Body).Decode(reqData)
	if err != nil {
		e := cerror.New(err.Error(), http.StatusUnprocessableEntity)
		e.ResHttp(w)
		return false
	}

	return true
}
