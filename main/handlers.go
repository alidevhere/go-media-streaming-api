package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func addHeaders(h http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

/*
// addHeaders will act as middleware to give us CORS support
func addHeaders() http.HandlerFunc {
	fmt.Println("view vide called")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		id := mux.Vars(r)["id"]
		h := http.FileServer(http.Dir(videoRenderingDir + "/upload-" + id))
		h.ServeHTTP(w, r)
	}
}
*/
//ERROR HANDLE IF FILE OR DECSCRIPTION NOT FOUND
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// 1- Extract file from FORM
	file, header, err := r.FormFile("video")
	product_id := r.Form.Get("product-id")
	description := r.Form.Get("description")
	tags := r.Form.Get("tags")

	//fmt.Println(product_id, description, tags)
	//fmt.Printf("product ID: %v\ndescription: %v\ntags: %v", product_id, description, tags)
	if err != nil {
		fmt.Println("Upload file key not found:", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer file.Close()

	//fmt.Printf("file name: %v", header.Filename)
	//fmt.Printf("file Size: %v", header.Size)
	//fmt.Printf("file Header: %v", header.Header)

	// 2- Create temp file on server
	outFile, PathErr := os.CreateTemp(uploadDir, "upload-*.mp4")
	if PathErr != nil {
		fmt.Println("Temporary file creation path not found", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer outFile.Close()

	// 3- COPY File to server
	io.Copy(outFile, file)
	fmt.Print(outFile.Name())

	go ProcessUploadFile(video{Product_id: product_id, Description: description, Tags: tags, Path: outFile.Name(), FileSize: header.Size})
	w.WriteHeader(http.StatusCreated)

}

func ProcessUploadFile(vid video) {
	//	Process file
	//	save its meta data inti DB
	//"ffmpeg -i 32.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8"
	//fmt.Print(vid.Path)
	dirName := strings.TrimSuffix(filepath.Base(vid.Path), filepath.Ext(vid.Path))
	newDir := videoRenderingDir + "/" + dirName
	_, dirCreationErr := exec.Command("mkdir", newDir).Output()
	if dirCreationErr != nil {
		log.Fatal("Directory for new video upload not created", dirCreationErr.Error())
	}
	out, err := exec.Command("ffmpeg", "-i", vid.Path, "-profile:v", "baseline", "-level", "3.0", "-s", "640x360", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", newDir+"/index.m3u8").Output()

	fmt.Print(out)
	if err != nil {
		log.Fatal(err)
	}
}

func renderVideo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("INNN...")
	id := mux.Vars(r)["id"]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	h := http.FileServer(http.Dir(videoRenderingDir + "/upload-" + id))
	h.ServeHTTP(w, r)
}
