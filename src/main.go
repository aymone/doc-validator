package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Fatal("Undefined env: PORT")
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), requestHandler(router()))
	if err != nil {
		log.Fatal("Error on starting falcon server")
	}
}

func requestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Accept", "text/json")
		w.Header().Set("Accept-Charset", "utf-8")
		w.Header().Set("Accept-Encoding", "gzip, deflate")
		h.ServeHTTP(w, r)
	})
}

func PingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\n"))
}

func router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ping", PingHandler)
	return router
}
