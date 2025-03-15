package api

import (
	"encoding/json"
	"net/http"
)

func parseBody(r *http.Request, t any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&t)
}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
