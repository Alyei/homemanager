package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

// Config contains the basic configuration options for the server.
type Config struct {
	TLS struct {
		Pkey string `json:"pkey"`
		Cert string `json:"cert"`
	}
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	config := loadConfig()

	log.Output(1, "Starting server.")

	// Setting up routes with mux.
	r := mux.NewRouter()
	r.HandleFunc("/", exampleHandler)

	// Starting the listener.
	err := http.ListenAndServeTLS(config.Port, config.TLS.Cert, config.TLS.Pkey, r)
	if err != nil {
		log.Fatal(err)
	}
}

// ExampleHandler from https://golangcode.com/basic-https-server-with-certificate/
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{ "status":"ok" }`)
}

// Loads the config from ./config.json
func loadConfig() Config {
	var config Config

	configFile, err := os.Open("config.json")
	defer configFile.Close()

	if err != nil {
		log.Fatal("Could not open config:", err.Error())
	}

	//Parse config and store it in the struct
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}
