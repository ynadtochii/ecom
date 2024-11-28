package health

import (
	"net/http"

	utils "github.com/ynadtochii/ecom/util"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Health(router *http.ServeMux) {
	router.HandleFunc("/health", h.health)

}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Ok"})
}
