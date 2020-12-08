package api

import (
	"encoding/json"
	"go-user/internal/model"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, r model.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	return json.NewEncoder(w).Encode(r)
}
