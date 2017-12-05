package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DocumentDeleteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serverInfo.setCounter()
	docID := p.ByName("documentNumber")

	if err := getClient().C("documents").RemoveId(docID); err != nil {
		log.Println(err.Error())
		http.Error(w, "Error on delete document", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
