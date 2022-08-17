package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Website struct {
	List   []string          `json:"websites"`
	Status map[string]string `json:"-"`
}

var websites Website

func Start() {

	fmt.Printf("**Welcome to the Website Checker**\n\nstarting the server...\n\n")
	r := mux.NewRouter()

	r.HandleFunc("/", welcomeApiHander)
	r.HandleFunc("/POST/websites", postWebsiteHandler).Methods("POST")
	r.HandleFunc("/GET/websites", getWebsiteHandler).Methods("GET")

	http.Handle("/", r)

	go StartChecker()
	log.Fatal(http.ListenAndServe(":3000", r))

}

// Starts the checker and update in every 60 seconds
func StartChecker() {
	for {
		websites.Status = WebsiteChecker(websites.List)
		time.Sleep(60 * time.Second)
	}

}

// handler for /
func welcomeApiHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, Welcome to the Website Checker</h1>")

}

// handler for /POST/websites
func postWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Hit: /POST/websites")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send list websites to check.")
		return
	}

	json.NewDecoder(r.Body).Decode(&websites)
	websites.Status = make(map[string]string)

	json.NewEncoder(w).Encode(fmt.Sprintf("Websites in checking: %v", websites.List))
	websites.Status = WebsiteChecker(websites.List)

	fmt.Printf("List of websites updated to %v Successfully.\n\n", websites.List)

}

// handler for /GET/websites?name=websitename(optional)
func getWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("API Hit: /GET/websites\n\n")
	w.Header().Set("Content-Type", "application/json")

	website_name := r.URL.Query().Get("name")

	if website_name != "" {
		resp := make(map[string]string)
		if websites.Status[website_name] == "" {
			resp[website_name] = "Not in database. Please register it by a POST request first."
		} else {
			resp[website_name] = websites.Status[website_name]
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(websites.Status)
}
