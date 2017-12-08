package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DocumentBlacklistHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serverInfo.setCounter()

	docID := p.ByName("id")
	docStatus := p.ByName("status")

	_, err := validate(docID)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var doc document
	findErr := getClient().C("documents").FindId(docID).One(&doc)
	if findErr != nil {
		log.Error(findErr)
		http.Error(w, findErr.Error(), http.StatusNotFound)
		return
	}

	statusErr := doc.setStatus(docStatus)
	if statusErr != nil {
		log.Error(statusErr)
		http.Error(w, statusErr.Error(), http.StatusBadRequest)
		return
	}

	updateErr := getClient().C("documents").UpdateId(docID, &doc)
	if updateErr != nil {
		log.Error(updateErr)
		http.Error(w, updateErr.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
