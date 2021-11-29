package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

//ERROR HANDLE IF FILE NOT FOUND
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// 1- Extract file from FORM
	file, header, err := r.FormFile("video")
	product_id := r.Form.Get("product-id")
	description := r.Form.Get("description")
	tags := r.Form.Get("tags")

	fmt.Println(product_id, description, tags)
	//fmt.Printf("product ID: %v\ndescription: %v\ntags: %v", product_id, description, tags)
	if err != nil {
		fmt.Println("Upload file key not found:", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer file.Close()

	fmt.Printf("file name: %v", header.Filename)
	fmt.Printf("file Size: %v", header.Size)
	fmt.Printf("file Header: %v", header.Header)

	// 2- Create temp file on server
	outFile, PathErr := os.CreateTemp("./media/uploads", "upload-*.mp4")
	if PathErr != nil {
		fmt.Println("Temporary file creation path not found", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer outFile.Close()

	// 3- COPY File to server
	io.Copy(outFile, file)
	fmt.Print(outFile.Name())
	w.WriteHeader(http.StatusCreated)

}

func ProcessUploadFile(vid video) {

}
