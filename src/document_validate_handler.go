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

func DocumentValidateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serverInfo.setCounter()

	var v ValidateRequest

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&v)
	if err != nil {
		log.Println("Error on decode request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, err := validate(v.DocumentNumber)
	if err != nil {
		log.Println("Error on validate responses:")
		log.Println(err)
		http.Error(w, "Error on validate responses", http.StatusBadRequest)
		return
	}

	v.Variety = c.variety
	v.IsValid = true

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error on encode responses:")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
