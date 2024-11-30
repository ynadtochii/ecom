package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
  return json.NewEncoder(w).Encode(payload)
}

func RespondError(w http.ResponseWriter, statusCode int, err interface{}) {
	var response map[string]string

	switch v := err.(type) {
	case error:
		response = map[string]string{"error": v.Error()}
	case string:
		response = map[string]string{"message": v}
	default:
		response = map[string]string{"error": "Unknown error type"}
	}

	RespondJSON(w, statusCode, response)
}

func DecodeJSONBody(r *http.Request, dst interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	defer r.Body.Close()

	// Limit the size of the request body (optional, for safety)
	const maxBodySize = 1 << 20 // 1 MB
	limitedReader := io.LimitReader(r.Body, maxBodySize)

	// Decode the JSON into the provided destination struct
	decoder := json.NewDecoder(limitedReader)
	decoder.DisallowUnknownFields() // Prevent unknown fields in the payload

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	// Ensure there are no extra bytes in the request body
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.New("request body must contain a single JSON object")
	}

	return nil
}
