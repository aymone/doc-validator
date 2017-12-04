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

	filter := r.URL.Query().Get("filter")
	filters := bson.M{"_id": bson.RegEx{Pattern: filter, Options: ""}}

	var docs []document
	if err := getClient().C("documents").Find(filters).All(&docs); err != nil {
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
