package main

import (
	"encoding/json"
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
		log.Error("Error on validate responses:")
		log.Error(err)
		http.Error(w, "Error on validate responses", http.StatusBadRequest)
		return
	}

	v := ValidateResponse{
		ID:      c.input,
		Variety: c.variety,
		IsValid: true,
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Error("Error on encode responses:")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
