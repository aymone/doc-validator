package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func DocumentCreateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serverInfo.setCounter()

	var doc document

	d := json.NewDecoder(r.Body)
	err := d.Decode(&doc)
	if err != nil {
		log.Println("Error on decode request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	c, err := validate(doc.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAt := time.Now()

	doc.ID = c.input
	doc.Variety = c.variety
	doc.CreatedAt = createdAt
	doc.UpdatedAt = createdAt

	if err := doc.create(); err != nil {
		log.Println(err.Error())
		http.Error(w, "Error on save new document", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(doc); err != nil {
		log.Println("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
