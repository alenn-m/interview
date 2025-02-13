package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alenn-m/interview/svc/pkg/pack/entity"
	"github.com/go-chi/chi/v5"
)

type DecodeEncoder interface {
	DecodeCreate(r *http.Request) (*entity.Pack, error)
	DecodeUpdate(r *http.Request) (*entity.Pack, error)
	DecodeID(r *http.Request) (int, error)
}

type decodeEncoder struct{}

func NewDecodeEncoder() DecodeEncoder {
	return &decodeEncoder{}
}

func (h *decodeEncoder) DecodeCreate(r *http.Request) (*entity.Pack, error) {
	var pack entity.Pack
	if err := json.NewDecoder(r.Body).Decode(&pack); err != nil {
		return nil, err
	}
	return &pack, nil
}

func (h *decodeEncoder) DecodeUpdate(r *http.Request) (*entity.Pack, error) {
	var pack entity.Pack
	if err := json.NewDecoder(r.Body).Decode(&pack); err != nil {
		return nil, err
	}
	return &pack, nil
}

func (h *decodeEncoder) DecodeID(r *http.Request) (int, error) {
	id := chi.URLParam(r, "id")
	return strconv.Atoi(id)
}
