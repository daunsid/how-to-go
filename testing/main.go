package main

import (
	"encoding/json"
	"net/http"
)

func handlGetFoo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	type response struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	json.
	w.Write([]byte("FOO"))
}
