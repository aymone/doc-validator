package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// DocumentReadHandler will handle read one requests.
// Errors can be returned if document not found or cannot encode response.
func DocumentReadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serverInfo.setCounter()

	docID := p.ByName("id")

	_, err := validate(docID)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var doc document
	if err := getClient().C("documents").FindId(docID).One(&doc); err != nil {
		log.Error(err.Error())
		http.Error(w, "Error on get document", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(doc); err != nil {
		log.Error("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
