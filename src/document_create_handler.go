package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

// DocumentCreateHandler will handle create requests.
// Before try to create any document,
// validator will be called and return error if document number is not valid.
// Theres also error codes related with item create.
func DocumentCreateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serverInfo.setCounter()

	var docR documentRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&docR)
	if err != nil {
		log.Error("Error on decode request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	c, err := validate(docR.ID)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAt := time.Now()
	doc := document{
		ID:        c.input,
		Variety:   c.variety,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	if err := doc.create(); err != nil {
		log.Error(err.Error())
		// Errors from go-mgo client not return status codes :/
		if strings.Contains(err.Error(), "E11000 duplicate") {
			http.Error(w, "Duplicated document number", http.StatusConflict)
			return
		}

		http.Error(w, "Error on save new document number", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(doc); err != nil {
		log.Error("Error on encode response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d *document) create() error {
	return getClient().C("documents").Insert(&d)
}
