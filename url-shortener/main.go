package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/handler"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := readYAMLData("urls.yml")
	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readYAMLData(ymlFileName string) []byte {
	yamlData, err := os.ReadFile(ymlFileName)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to open the YAML file: %s\n", err))
	}

	return yamlData
}
