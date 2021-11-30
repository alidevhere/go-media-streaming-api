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
	outFile, PathErr := os.CreateTemp(uploadDir, "*.mp4")
	if PathErr != nil {
		fmt.Println("Temporary file creation path not found", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer outFile.Close()

	// 3- COPY File to server
	io.Copy(outFile, file)
	//fmt.Print(outFile.Name())

	go ProcessUploadFile(video{Product_id: product_id, Description: description, Tags: tags, Path: outFile.Name(), FileSize: header.Size})
	w.WriteHeader(http.StatusCreated)

}

func ProcessUploadFile(vid video) {
	//	Process file
	//	save its meta data inti DB
	//"ffmpeg -i 32.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8"
	//fmt.Print(vid.Path)
	dirName := strings.TrimSuffix(filepath.Base(vid.Path), filepath.Ext(vid.Path))
	//fmt.Println(dirName)
	//dirName = dirName[7:]
	//fmt.Println(dirName)

	//fmt.Println("ONLY DIR NAME", dirName, "===")
	newDir := videoRenderingDir + "/" + dirName
	//fmt.Println("NEW DIR:", newDir, ":::END")
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

//224728450
func StreamM3U8(w http.ResponseWriter, r *http.Request) {
	id, idErr := mux.Vars(r)["id"]
	if !idErr {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//0-Base dir for .m3u8 and .ts files of id
	m3u8Path := fmt.Sprintf("%s/%s/index.m3u8", videoRenderingDir, id)
	http.ServeFile(w, r, m3u8Path)
}

func StreamTS(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, idErr := vars["id"]
	segNo, segErr := vars["segNo"]
	if !idErr || !segErr {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Printf("id:%s   seg: %s", id, segNo)
	tsPath := fmt.Sprintf("%s/%s/%s", videoRenderingDir, id, segNo)
	http.ServeFile(w, r, tsPath)
}
