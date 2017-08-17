package main

import (
	"net/http"
	"encoding/json"
)

func Contact(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cvs); err != nil {
		panic(err)
	}

}