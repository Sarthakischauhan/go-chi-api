package orders

import (
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

func (h *handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var ordersPayload createOrderParams

	ctx := r.Context()
	err := json.Read(r, &ordersPayload)

	if err != nil {
		log.Println("Couldn't read the request payload")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderCreated, err := h.service.createOrder(ctx, ordersPayload)

	if err != nil {
		log.Println("Couldn't read the request payload")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusAccepted, orderCreated)

}
