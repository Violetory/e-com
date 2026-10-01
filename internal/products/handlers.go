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
	if err := h.service.ListProducts(r.Context()); err != nil {
		log.Println("获取产品列表失败:", err)
		http.Error(w, "获取产品列表失败", http.StatusInternalServerError)
		return
	}

	productList := []string{"Hello", "World"}
	json.Write(w, http.StatusOK, productList)
}
