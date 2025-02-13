package resp

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	utilErrors "github.com/alenn-m/interview/svc/util/errors"
)

func ReturnError(w http.ResponseWriter, err error, code int) {
	if errors.Is(err, sql.ErrNoRows) {
		code = http.StatusNotFound
		err = errors.New("not found")
	} else if errors.As(err, &utilErrors.ErrValidation{}) {
		code = http.StatusBadRequest
	}

	result, _ := json.Marshal(err.Error())

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
