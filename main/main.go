package main

//COMAND FFMPEG
//ffmpeg -i 32.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// configure the songs directory name and port
	const songsDir = "media/videos"
	const port = 8080

	// add a handler for the song files
	//http.Handle("/", addHeaders(http.FileServer(http.Dir(songsDir))))
	//fmt.Printf("Starting server on %v\n", port)
	//log.Printf("Serving %s on HTTP port: %v\n", songsDir, port)
	//serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload-video", UploadFile).Methods("POST")
	router.HandleFunc("/video/{id}", addHeaders(http.FileServer(http.Dir(songsDir)))).Methods("GET")
	fmt.Println("Listening...")
	server := &http.Server{Addr: ":8080", Handler: router}
	server.ListenAndServe()

}
