package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Load Server Configurations
var Configurations serverConfig

func main() {

	Configurations = loadServerConfig("/home/ali/go/src/Media_Streaming_API_Swipe_Shop/serverConfig.json")

	// === ROUTERS ===
	router := mux.NewRouter()
	router.HandleFunc("/upload-video", UploadFile).Methods("POST")

	// ===	Serving 	===
	fmt.Printf("Starting server on %v\n", Configurations.Addr)
	server := &http.Server{Addr: Configurations.Addr, Handler: router}
	err := server.ListenAndServe()
	fmt.Print(err.Error())

}
