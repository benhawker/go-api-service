package friendships

import (
	"encoding/json"
	"net/http"

	// "github.com/aws/aws-sdk-go/service/dynamodb"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Show(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal("{}")
	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h *Handler) Create(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal("{}")
	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h *Handler) Common(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal("{}")
	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}


