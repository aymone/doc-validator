package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
		// Errors from go-mgo client not return status codes :/
		if strings.Contains(err.Error(), "E11000 duplicate") {
			http.Error(w, "Duplicated document number", http.StatusConflict)
			return
		}

		http.Error(w, "Error on save new document number", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(doc); err != nil {
		log.Println("Error on encode response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
