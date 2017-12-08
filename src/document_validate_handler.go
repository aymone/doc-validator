package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ValidateResponse struct {
	ID      string `json:"id"`
	IsValid bool   `json:"isValid"`
	Variety string `json:"type"`
}

func DocumentValidateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serverInfo.setCounter()

	docID := p.ByName("id")
	c, err := validate(docID)
	if err != nil {
		log.Println("Error on validate responses:")
		log.Println(err)
		http.Error(w, "Error on validate responses", http.StatusBadRequest)
		return
	}

	v := ValidateResponse{
		ID:      c.input,
		Variety: c.variety,
		IsValid: true,
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error on encode responses:")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
