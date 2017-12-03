package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ValidateRequest struct {
	DocumentNumber string `json:"documentNumber"`
	IsValid        bool   `json:"isValid"`
	Variety        string `json:"type"`
}

func ValidateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serverInfo.setCounter()

	var v ValidateRequest

	d := json.NewDecoder(r.Body)
	err := d.Decode(&v)
	if err != nil {
		log.Println("Error on decode request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	c, err := validate(v.DocumentNumber)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	v.Variety = c.variety
	v.IsValid = true

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Fatal("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
