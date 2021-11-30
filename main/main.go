package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//COMAND FFMPEG
//ffmpeg -i 32.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8
//http://localhost:8080/32/index.m3u8

const (
	uploadDir         string = "/home/ali/go/src/Media_Streaming_API_Swipe_Shop/main/media/uploads"
	videoRenderingDir string = "/home/ali/go/src/Media_Streaming_API_Swipe_Shop/main/media/videos/upload-2146559729/index.m3u8"
	Addr              string = ":8080"
	//conversionCommand string = "ffmpeg -i 32.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8"
)

func main() {

	// === ROUTERS ===
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload-video", UploadFile).Methods("POST")
	router.HandleFunc("/", addHeaders(http.FileServer(http.Dir(videoRenderingDir)))).Methods("GET")
	//router.HandleFunc("/video/{id}/index.m3u8", renderVideo).Methods("GET")
	//router.HandleFunc("/video/{id}", addHeaders()).Methods("GET")

	// ===	Serving 	===
	fmt.Printf("Starting server on %v\n", Addr)
	server := &http.Server{Addr: Addr, Handler: router}
	server.ListenAndServe()

}
