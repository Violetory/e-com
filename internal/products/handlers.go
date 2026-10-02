package products

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Violetory/e-com/internal/json"
	"github.com/jackc/pgx/v5"
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

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// 取出ID，转化为 Int64
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "无效的产品ID", http.StatusBadRequest)
		return
	}

	// 查询商品
	product, err := h.service.GetProductByID(r.Context(), id)

	// 查询成功，但商品不存在
	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "商品不存在", http.StatusNotFound)
		return
	}

	// 其他查询错误
	if err != nil {
		log.Println("获取产品失败:", err)
		http.Error(w, "获取产品失败", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
