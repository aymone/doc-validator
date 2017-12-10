package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// DocumentDeleteHandler will handle delete requests.
// Return error if document not found.
func DocumentDeleteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serverInfo.setCounter()
	docID := p.ByName("id")

	if err := getClient().C("documents").RemoveId(docID); err != nil {
		log.Error(err.Error())
		http.Error(w, "Error on delete document", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
