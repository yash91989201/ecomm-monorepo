package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yash91989201/ecomm-monorepo/common/clients"
	"github.com/yash91989201/ecomm-monorepo/common/pb"
	"github.com/yash91989201/ecomm-monorepo/common/types"
)

type handler struct {
	ctx             context.Context
	inventoryClient *clients.InventoryClient
}

func NewHandler(ctx context.Context, client *clients.InventoryClient) *handler {
	return &handler{
		ctx:             ctx,
		inventoryClient: client,
	}
}

func (h *handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var p *types.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	product, err := h.inventoryClient.CreateProduct(h.ctx, types.ToPBProductReq(p))
	if err != nil {
		log.Print(err)
		http.Error(w, "error creating product", http.StatusInternalServerError)
		return
	}

	res := types.ToProduct(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) getProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	product, err := h.inventoryClient.GetProduct(h.ctx, &pb.ProductReq{Id: id})
	if err != nil {
		http.Error(w, "error getting product", http.StatusInternalServerError)
		return
	}

	res := types.ToProduct(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	product, err := h.inventoryClient.GetProduct(h.ctx, &pb.ProductReq{Id: id})
	if err != nil {
		http.Error(w, "error getting product", http.StatusInternalServerError)
		return
	}

	res := types.ToProduct(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
