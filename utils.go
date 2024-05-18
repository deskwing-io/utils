package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ReadBody reads and unmarshals the request body into the given interface.
func ReadBody(r *http.Request, req interface{}) error {
	dataRaw, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataRaw, &req)
	if err != nil {
		return err
	}
	return nil
}

// RespondWithJSON writes a JSON response with the given status code and payload.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

// RespondWithError writes a JSON response with the given error.
func RespondWithError(w http.ResponseWriter, err Error) {
	RespondWithJSON(
		w, err.HttpCode, map[string]string{
			"message": err.Message,
			"code":    err.Code,
		},
	)
}

// Error represents an HTTP error with a message, code, and HTTP status code.
type Error struct {
	Message  string
	Code     string
	HttpCode int
}

// NewError creates a new Error with the given message, code, and HTTP status code.
func NewError(msg, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}
