package utils

import (
	"encoding/json"
	"net/http"
)

type JSON struct{}

func (j JSON) Read(w http.ResponseWriter, r *http.Request, data any) error {
	body := http.MaxBytesReader(w, r.Body, int64(1024*1024))

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func (j JSON) Write(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func (j *JSON) Error(w http.ResponseWriter, status int, data any) error {
	return j.Write(w, status, data)
}
