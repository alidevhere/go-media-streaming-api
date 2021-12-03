package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type video struct {
	Video_id    string
	Product_id  string
	Description string
	Tags        string
	Path        string
	FileSize    int64
}

type serverConfig struct {
	Addr      string `json:"Addr"`
	UploadDir string `json:"uploadDir"`
	RenderDir string `json:"renderDir"`
}

func loadServerConfig(path string) serverConfig {
	var Configurations serverConfig

	file, fileErr := ioutil.ReadFile(path)

	if fileErr != nil {
		log.Fatal(fileErr.Error())
	}

	json.Unmarshal(file, &Configurations)
	return Configurations
}
