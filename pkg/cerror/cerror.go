package cerror

import (
	"encoding/json"
	"net/http"
)

type CError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func New(msg string, code int) *CError {
	return &CError{
		Msg:  msg,
		Code: code,
	}
}

func (c *CError) ResHttp(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(c.Code)
	json.NewEncoder(w).Encode(c)
}

func (c *CError) Error() string {
	return c.Msg
}
