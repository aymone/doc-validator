package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

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
		log.Println(err.Error())
		http.Error(w, "Error on get documents", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(docs); err != nil {
		log.Println("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
