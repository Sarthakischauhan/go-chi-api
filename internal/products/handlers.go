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
	products, err := h.service.GetProducts(r.Context())
	if err != nil {
		log.Println("Error fetching product records:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusAccepted, products)
}

func (h *handler) AddProductsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var productBody createProductParams

	err := json.Read(r, &productBody)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := h.service.CreateProduct(r.Context(), productBody)

	if err != nil {
		log.Println("Error while adding")
		json.Write(w, http.StatusInternalServerError, createdProduct)
	}

	json.Write(w, http.StatusAccepted, createdProduct)
}

// Create handlers for different services here
