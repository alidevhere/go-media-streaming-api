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
)

//ERROR HANDLE IF FILE OR DECSCRIPTION NOT FOUND
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// 1- Extract file from FORM
	file, header, err := r.FormFile("video")
	product_id := r.Form.Get("product-id")
	description := r.Form.Get("description")
	tags := r.Form.Get("tags")

	if err != nil {
		fmt.Println("Upload file key not found:", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer file.Close()

	// 2- Create temp file on server
	outFile, PathErr := os.CreateTemp(Configurations.UploadDir, "*.mp4")
	if PathErr != nil {
		log.Fatal("Temporary file creation path not found", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer outFile.Close()

	// 3- COPY File to server
	io.Copy(outFile, file)

	go ProcessUploadFile(video{Product_id: product_id, Description: description, Tags: tags, Path: outFile.Name(), FileSize: header.Size})
	w.WriteHeader(http.StatusCreated)

}

func ProcessUploadFile(vid video) {
	dirName := strings.TrimSuffix(filepath.Base(vid.Path), filepath.Ext(vid.Path))
	newDir := Configurations.RenderDir + "/" + dirName
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
