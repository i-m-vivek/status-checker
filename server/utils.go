package server

import (
	"fmt"
	"net/http"
	"sync"
)

type StatusChecker interface {
	Check(myurl string) (status bool, err error)
}
type HttpChecker struct {
}

// Check if website is UP or DOWN
func (h HttpChecker) Check(myurl string) (status string,
	err error) {

	r, e := http.Get(myurl)

	if err != nil {
		err = e
		return
	}
	if r == nil {
		status = "DOWN"
		return
	}
	if r.StatusCode == 200 {
		status = "UP"
		return
	}

	status = "DOWN"
	return
}

// helper func that takes a web url check and update status in given map of website
func checkHelper(myurl string, wg *sync.WaitGroup, website_status *map[string]string) {
	defer wg.Done()

	var hchecker HttpChecker
	status, err := hchecker.Check(myurl)

	if err != nil {
		fmt.Printf("**ERROR** website: %v error: %v\n\n", myurl, err)
	}

	(*website_status)[myurl] = status

}

// takes a list of website url and return a map with their status
func WebsiteChecker(website_list []string) map[string]string {
	var wg sync.WaitGroup

	website_status := make(map[string]string)

	for _, myurl := range website_list {
		wg.Add(1)
		go checkHelper(myurl, &wg, &website_status)
	}

	wg.Wait()

	fmt.Printf("Website status updated successfully\n")
	return website_status
}
