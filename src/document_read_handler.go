package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DocumentReadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serverInfo.setCounter()

	var doc document
	doc.ID = p.ByName("documentNumber")

	_, err := validate(doc.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := doc.read(); err != nil {
		log.Println(err.Error())
		http.Error(w, "Error on get document", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(doc); err != nil {
		log.Println("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
