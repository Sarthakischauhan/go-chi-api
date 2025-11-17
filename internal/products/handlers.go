package products

import (
	// "encoding/json"
	"log"
	"net/http"

	"github.com/Sarthakischauhan/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	err := h.service.GetProducts(r.Context())
	products := []string{"Sarthak", "Chauhan"}
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.Write(w, http.StatusAccepted, products)

}
