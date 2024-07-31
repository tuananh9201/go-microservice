package main

import (
	"errors"
	"log"
	"net/http"

	common "github.com/tuananh9201/commons"
	pb "github.com/tuananh9201/commons/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandlerCreateOrder)
}

func (h *handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	// gateway
	customerID := r.PathValue("customerID")
	var items []*pb.ItemsWithQWuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, "gateway: ReadJSON: "+err.Error())
		return
	}
	log.Println("gateway: HandlerCreateOrder")
	err := validateItems(items)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, "gateway: validateItems: "+err.Error())
		return
	}
	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "gateway: CreateOrder: "+err.Error())
		return
	}
	common.WriteJSON(w, http.StatusCreated, o)
}

func validateItems(items []*pb.ItemsWithQWuantity) error {
	if len(items) == 0 {
		return errors.New("items must have at least one item")
	}
	for _, item := range items {
		if item.ID == "" {
			return errors.New("product_id is required")
		}
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
	}
	return nil
}
