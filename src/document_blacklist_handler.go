package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DocumentBlacklistHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serverInfo.setCounter()

	docID := p.ByName("documentNumber")
	docStatus := p.ByName("status")

	_, err := validate(docID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var doc document
	if err := doc.blacklist(docID, docStatus); err != nil {
		log.Println(err.Error())
		http.Error(w, "Error on update document", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(doc); err != nil {
		log.Println("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
