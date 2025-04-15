package utils

import "net/http"

type JSONHelper interface {
	Read(http.ResponseWriter, *http.Request, any) error
	Write(http.ResponseWriter, int, any) error
	Error(http.ResponseWriter, int, any) error
}
