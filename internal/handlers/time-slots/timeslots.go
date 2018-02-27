package timeslots

import (
	"encoding/json"
	"net/http"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal("{}")
	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}