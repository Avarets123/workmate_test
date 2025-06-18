package tasks

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"workmate/pkg/cerror"
	"workmate/pkg/utils"

	"github.com/gorilla/mux"
)

type handler struct {
	taskService *Service
}

func ApplyHandler(router *mux.Router, service *Service) {

	h := handler{
		taskService: service,
	}

	router.HandleFunc("/tasks", h.create).Methods(http.MethodPost)
	router.HandleFunc("/tasks", h.listing).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}/cancel", h.cancel).Methods(http.MethodPatch)
	router.HandleFunc("/tasks/{id}", h.findOne).Methods(http.MethodGet)

}

func (h *handler) findOne(w http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		cerr := cerror.New("id must be passed!", http.StatusUnprocessableEntity)
		cerr.ResHttp(w)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		cerr := cerror.New(err.Error(), http.StatusUnprocessableEntity)
		cerr.ResHttp(w)
		return
	}

	task, cerr := h.taskService.FindOne(idInt)
	if cerr != nil {
		cerr.ResHttp(w)
		return
	}

	h.responseHttp(w, http.StatusOK, task)

}

func (h *handler) cancel(w http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		cerr := cerror.New("id must be passed!", http.StatusUnprocessableEntity)
		cerr.ResHttp(w)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		cerr := cerror.New(err.Error(), http.StatusUnprocessableEntity)
		cerr.ResHttp(w)
		return
	}

	cerr := h.taskService.Cancel(idInt)
	if cerr != nil {
		cerr.ResHttp(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *handler) listing(w http.ResponseWriter, req *http.Request) {

	res := h.taskService.Listing()

	h.responseHttp(w, http.StatusOK, res)

}

func (h *handler) create(w http.ResponseWriter, req *http.Request) {
	reqData := TaskCreateReq{}

	isValid := utils.ValidateReq(w, req, http.MethodPost, &reqData)
	if !isValid {
		return
	}

	cerr := reqData.Validate()
	if cerr != nil {
		cerr.ResHttp(w)
		return
	}

	task := h.taskService.Create(reqData)

	h.responseHttp(w, http.StatusCreated, task)

}

func (h *handler) responseHttp(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if reflect.ValueOf(body).IsZero() {
		return
	}

	json.NewEncoder(w).Encode(body)

}
