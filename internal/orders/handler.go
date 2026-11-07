package orders

import (
	"log"
	"net/http"

	"github.com/nnaxim/go-backend-study/internal/json"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var tempOrder createdOrderParams

	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder, err := h.service.PlaceOrder(r.Context(), tempOrder)
	if err != nil {
		log.Println(err)
		if err == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err == ErrProductNoStock {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		return
	}

	json.Write(w, http.StatusCreated, createOrder)
}