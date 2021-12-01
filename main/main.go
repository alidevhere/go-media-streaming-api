package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	uploadDir         string = "/home/ali/go/src/Media_Streaming_API_Swipe_Shop/main/media/uploads"
	videoRenderingDir string = "/home/ali/go/src/Media_Streaming_API_Swipe_Shop/main/media/videos"
	Addr              string = ":8080"
)

func main() {

	// === ROUTERS ===
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload-video", UploadFile).Methods("POST")
	router.HandleFunc("/video/{id:[0-9]+}", StreamM3U8)
	router.HandleFunc("/video/{id:[0-9]+}/{segNo:index[0-9]+.ts}", StreamTS)
	// Router for serving complete directory of HLS files
	router.HandleFunc("/video/{id:[0-9]+}/stream", addHeaders())

	// ===	Serving 	===
	fmt.Printf("Starting server on %v\n", Addr)
	server := &http.Server{Addr: Addr, Handler: router}
	server.ListenAndServe()

}
