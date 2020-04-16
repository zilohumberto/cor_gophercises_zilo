// https://courses.calhoun.io/lessons/les_goph_04
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zilohumberto/cor_gophercises_zilo/exercise_2_URL_shortener/urlshort"
)

var (
	yamlFile = flag.String("input", "example.yml", "yaml example urls")
)

func main() {
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := readFile()
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	check(err)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() string {
	data, err := ioutil.ReadFile(*yamlFile)
	check(err)
	return string(data)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
