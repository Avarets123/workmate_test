package tasks

import (
	"net/http"
	"workmate/pkg/cerror"
)

func (t *TaskCreateReq) Validate() *cerror.CError {

	if t.Name == "" {
		return &cerror.CError{
			Msg:  "Field 'name' must not be empty!",
			Code: http.StatusUnprocessableEntity,
		}
	}

	return nil

}
