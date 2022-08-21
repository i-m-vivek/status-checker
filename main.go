package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/i-m-vivek/status-checker/cache"
	"github.com/i-m-vivek/status-checker/router"
	"github.com/i-m-vivek/status-checker/util"
)

func main() {

	fmt.Printf("** Welcome to the Website Checker **\n\nstarting the server...\n\n")
	r := router.Router()

	http.Handle("/", r)

	go StartChecker()
	log.Fatal(http.ListenAndServe(":3000", r))

}

// Starts the checker and update in every 60 seconds
func StartChecker() {
	for {
		cache.Websites.Status = util.WebsiteChecker(cache.Websites.List)
		time.Sleep(60 * time.Second)
	}

}
