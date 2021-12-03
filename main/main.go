package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Load Server Configurations
var Configurations = loadServerConfig("serverConfig.json")

func main() {

	// === ROUTERS ===
	router := mux.NewRouter()
	router.HandleFunc("/upload-video", UploadFile).Methods("POST")
	// ===	Serving 	===
	fmt.Printf("Starting server on %v\n", Configurations.Addr)
	server := &http.Server{Addr: Configurations.Addr, Handler: router}
	server.ListenAndServe()

}
