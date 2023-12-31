package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if val, doesExist := pathsToUrls[r.URL.Path]; doesExist {
			w.Header().Set("Content-Type", "application/json")
			http.Redirect(w, r, val, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]urlRedirect, error) {
	pathsToUrls := make([]urlRedirect, 0)

	err := yaml.Unmarshal(yml, &pathsToUrls)
	if err != nil {
		return nil, err
	}

	return pathsToUrls, nil
}

func buildMap(urlObjs []urlRedirect) map[string]string {
	pathMap := make(map[string]string)

	for _, urlRedirect := range urlObjs {
		pathMap[urlRedirect.Path] = urlRedirect.URL
	}

	return pathMap
}

type urlRedirect struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

////////////////////

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJson(jsonData)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJson)
	return MapHandler(pathMap, fallback), nil
}

func parseJson(jsonData []byte) ([]urlRedirect, error) {
	pathsToUrls := make([]urlRedirect, 0)

	err := json.Unmarshal(jsonData, &pathsToUrls)
	if err != nil {
		return nil, err
	}

	return pathsToUrls, nil
}
