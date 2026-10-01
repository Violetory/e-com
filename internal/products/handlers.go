package products

import (
	"github.com/Violetory/e-com/internal/json"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println("获取产品列表失败:", err)
		http.Error(w, "获取产品列表失败", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}
