package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/ProductHealth/gose4"
)

// Validate the version of a running se4 endpoint
func main() {
	endpoint := flag.String("endpoint", "", "Full se4 status url")
	expectedVersion := flag.String("version", "", "Version to validate")
	duration := flag.Duration("wait", 0, "Max time to wait for server to report desired version")
	interval := flag.Duration("interval", time.Second*10, "Time to wait between requests")

	startTime := time.Now()

	flag.Parse()
	if *endpoint == "" {
		log.Fatalf("endpoint parameter required")
	}
	if *expectedVersion == "" {
		log.Fatalf("version parameter required")
	}
	u, err := url.Parse(*endpoint)
	for {
		checkVersion(u, expectedVersion)
		if duration.Seconds() == 0 {
			log.Fatalf("Correct version not returned")
		} else if time.Since(startTime).Seconds() > duration.Seconds() {
			log.Fatalf("Correct version not returned within %v", duration)
		} else {
			log.Printf("Correct version not yet returned, waiting for %v", interval)
			time.Sleep(*interval)
		}
	}
	if err != nil {
		log.Fatalf("Could not parse url %v", *endpoint)
	}



}

func checkVersion(u *url.URL, expectedVersion *string) {
	resp, err := http.Get(u.String())
	if err != nil {
		log.Printf("Encountered error while performing request : %v", err)
		return
	}
	status := new(gose4.Status)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Could not read http response : %v", err)
	}
	err = json.Unmarshal(bytes, status)
	if err != nil {
		log.Printf("Could not unmarshal json : %v", err)
	}
	if status.BuildNumber != *expectedVersion {
		log.Printf("Running build number '%v' does not match expected '%v'", status.BuildNumber, *expectedVersion)
	} else {
		os.Exit(0)
	}
}
