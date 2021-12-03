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
	router := mux.NewRouter()
	router.HandleFunc("/upload-video", UploadFile).Methods("POST")

	router.HandleFunc("/video/{id:[0-9]+}/stream/", StreamHandler).Methods("GET")
	router.HandleFunc("/video/{id:[0-9]+}/stream/{segNo:index[0-9]+.ts}", StreamHandler).Methods("GET")

	//router.HandleFunc("/video/{id:[0-9]+}/stream/", StreamM3U8)
	//router.HandleFunc("/video/{id:[0-9]+}/{segNo:index[0-9]+.ts}", StreamTS)

	//router.HandleFunc("/video/{id:[0-9]+}/stream", RenderVideoHLS())

	//
	//router.HandleFunc("/video/stream/{id:[0-9]+}", StreamDir)
	// Router for serving complete directory of HLS files

	// ===	Serving 	===
	fmt.Printf("Starting server on %v\n", Addr)
	server := &http.Server{Addr: Addr, Handler: router}
	server.ListenAndServe()

}
