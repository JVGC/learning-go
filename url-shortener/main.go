package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/handler"
)

func main() {
	ymlFileFlag := flag.String("yml", "urls.yml", "YAML file containing the URL mappings")
	flag.Parse()
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := readFileData(*ymlFileFlag)
	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	json := readFileData("urls.json")

	jsonHandler, err := handler.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFileData(fileName string) []byte {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to open the %s file: %s\n", fileName, err))
	}

	return fileData
}
