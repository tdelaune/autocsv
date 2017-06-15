package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var file, separator, searchField, neededFields string
var feed *Feed

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&file, "file", "", "File datasource")
	flag.StringVar(&separator, "separator", ";", "csv field separator")
	flag.StringVar(&searchField, "search-field", "", "Searchable field")
	flag.StringVar(&neededFields, "needed-fields", "", "Expected fields")

	flag.Parse()

	if file == "" || searchField == "" || neededFields == "" {
		fmt.Println("options required !")
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func main() {
	feed = New(file, separator, searchField, neededFields)
	feed.Parse()

	go http.HandleFunc("/autocomplete", requestHandler)

	fmt.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("q")

	if keyword == "" {
		http.Error(w, "Not found", http.StatusUnprocessableEntity)
		return
	}

	results := feed.Search(keyword)

	json, err := json.Marshal(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
