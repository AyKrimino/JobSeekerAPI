package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func EncodeStringSliceToJSON(s []string) ([]byte, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("failed to encode string slice to JSON: %w", err)
	}
	return jsonData, nil 
}

func DecodeJSONTOStringSlice(jsonData []byte) ([]string, error) {
	var result []string

	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON data to string slice: %w", err)
	}
	return result, nil
}
