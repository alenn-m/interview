package http

import (
	"net/http"

	"github.com/alenn-m/interview/svc/pkg/order/service"
	"github.com/alenn-m/interview/svc/util/resp"
	"github.com/alenn-m/interview/svc/util/router"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc    service.Service
	decode DecodeEncoder
}

func Register(r router.Router, svc service.Service, dec DecodeEncoder) {
	h := &Handler{
		svc:    svc,
		decode: dec,
	}

	r.Chi.Route("/order", func(r chi.Router) {
		r.Post("/create", h.Create)
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	request, err := h.decode.DecodeCreate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := h.svc.Create(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resp.EncodeResponse(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
