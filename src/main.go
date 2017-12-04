package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var serverInfo *ServerInfo

func main() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[Document-validator]")
	serverInfo = &ServerInfo{}
	serverInfo.init()

	log.Printf("Server starting at %s", serverInfo.getStartedAt())

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

func router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/status", ServerStatusHandler)
	router.POST("/documents/validate", DocumentValidateHandler)
	router.POST("/documents", DocumentCreateHandler)
	router.PUT("/documents/:documentNumber/blacklist/:status", DocumentBlacklistHandler)
	router.GET("/documents/:documentNumber", DocumentReadHandler)
	// router.DELETE("/documents/:documentNumber", DocumentDeleteHandler)
	return router
}
