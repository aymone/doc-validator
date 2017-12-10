package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// DocumentReadAllHandler will handle read all requests.
// This handler accepts requests with filter or sorters.
// Errors can be returned if document not found or cannot encode response.
func DocumentReadAllHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serverInfo.setCounter()

	var docs []document

	filters := bson.M{}
	if filter := r.URL.Query().Get("filter"); filter != "" {
		filters["_id"] = bson.RegEx{Pattern: filter, Options: ""}
	}

	var sorters []string
	if sorter := r.URL.Query().Get("sort"); sorter != "" {
		if sorter == "id" {
			sorter = "_id"
		}

		sorters = append(sorters, sorter)
	}

	if err := getClient().C("documents").Find(filters).Sort(sorters...).All(&docs); err != nil {
		log.Error(err.Error())
		http.Error(w, "Error on get documents", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(docs); err != nil {
		log.Error("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
