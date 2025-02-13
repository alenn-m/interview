package http

import (
	"encoding/json"
	"net/http"

	"github.com/alenn-m/interview/svc/pkg/order/entity"
)

type DecodeEncoder interface {
	DecodeCreate(r *http.Request) (*entity.Request, error)
}

type decodeEncoder struct{}

func NewDecodeEncoder() DecodeEncoder {
	return &decodeEncoder{}
}

func (h *decodeEncoder) DecodeCreate(r *http.Request) (*entity.Request, error) {
	var request entity.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return &request, nil
}
