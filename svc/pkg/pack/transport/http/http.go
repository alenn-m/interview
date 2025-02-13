package http

import (
	"net/http"

	"github.com/alenn-m/interview/svc/pkg/pack/service"
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

	r.Chi.Route("/packs", func(r chi.Router) {
		r.Post("/", h.Create)       // POST /packs - Create a new pack
		r.Get("/", h.List)          // GET /packs - List all packs
		r.Get("/{id}", h.Get)       // GET /packs/{id} - Get a specific pack
		r.Put("/{id}", h.Update)    // PUT /packs/{id} - Update a specific pack
		r.Delete("/{id}", h.Delete) // DELETE /packs/{id} - Delete a specific pack
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	pack, err := h.decode.DecodeCreate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.svc.Create(r.Context(), pack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := resp.EncodeResponse(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resp.EncodeResponse(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := h.decode.DecodeID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.svc.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resp.EncodeResponse(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	pack, err := h.decode.DecodeUpdate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pack.ID, err = h.decode.DecodeID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.svc.Update(r.Context(), pack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resp.EncodeResponse(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.decode.DecodeID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.svc.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
