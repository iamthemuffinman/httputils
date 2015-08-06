package httputils

import (
	"encoding/json"
	"net/http"
)

// WriteJSONError writes a string as JSON encoded error
func WriteJSONError(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	WriteJSON(w, err, code)
}

// WriteJSON writes the given interface as JSON or returns an error
func WriteJSON(w http.ResponseWriter, v interface{}, code int) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(data)
	return nil
}
