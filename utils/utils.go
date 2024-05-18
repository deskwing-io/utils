package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

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

func RespondWithError(w http.ResponseWriter, err Error) {
	RespondWithJSON(
		w, err.HttpCode, map[string]string{
			"message": err.Message,
			"code":    err.Code,
		},
	)
}

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}
