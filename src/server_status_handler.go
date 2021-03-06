package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ServerStatus has the response body for ServerStatusHandler
type ServerStatus struct {
	StartedAt string `json:"startedAt"`
	Uptime    string `json:"uptime"`
	Requests  int    `json:"requests"`
}

// ServerStatusHandler handle server info, can return errors on encode response
func ServerStatusHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serverInfo.setCounter()

	status := ServerStatus{
		StartedAt: serverInfo.getStartedAt(),
		Uptime:    serverInfo.getUptime(),
		Requests:  serverInfo.getCounter(),
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Error("Error on encode responses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
