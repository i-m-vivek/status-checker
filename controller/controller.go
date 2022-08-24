package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/i-m-vivek/status-checker/cache"
	"github.com/i-m-vivek/status-checker/util"
)

// handler for /
func WelcomeApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, Welcome to the Website Checker</h1>")

}

// handler for /POST/websites
func PostWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Hit: /POST/websites")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send list of websites to check."})
		return
	}

	json.NewDecoder(r.Body).Decode(&cache.Websites)
	cache.Websites.Status = make(map[string]string)

	json.NewEncoder(w).Encode(map[string][]string{"websites": cache.Websites.List})
	cache.Websites.Status = util.WebsiteChecker(cache.Websites.List)

	fmt.Printf("List of websites updated to %v Successfully.\n\n", cache.Websites.List)

}

// handler for /GET/websites?name=websitename(optional)
func GetWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("API Hit: /GET/websites\n\n")
	w.Header().Set("Content-Type", "application/json")

	websiteName := r.URL.Query().Get("name")

	if websiteName != "" {
		resp := make(map[string]string)
		if cache.Websites.Status[websiteName] == "" {
			resp[websiteName] = "Not in database. Please register it by a POST request first."
		} else {
			resp[websiteName] = cache.Websites.Status[websiteName]
		}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			fmt.Println("Error:: in GetWebsiteHandler:", err)
			return
		}
		return
	}

	err := json.NewEncoder(w).Encode(cache.Websites.Status)
	if err != nil {
		fmt.Println("Error:: in GetWebsiteHandler:", err)
		return
	}
}
