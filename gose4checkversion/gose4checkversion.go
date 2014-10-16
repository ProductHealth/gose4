package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ProductHealth/gose4"
)

// Validate the version of a running se4 endpoint
func main() {
	endpoint := flag.String("endpoint", "", "Full se4 status url")
	expectedVersion := flag.String("version", "", "Version to validate")
	flag.Parse()
	if *endpoint == "" {
		log.Fatalf("endpoint parameter required")
	}
	if *expectedVersion == "" {
		log.Fatalf("version parameter required")
	}

	u, err := url.Parse(*endpoint)
	if err != nil {
		log.Fatalf("Could not parse url %v", *endpoint)
	}
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalf("Encountered error while performing request : %v", err)
	}
	status := new(gose4.Status)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Could not read http response : %v", err)
	}
	err = json.Unmarshal(bytes, status)
	if err != nil {
		log.Fatalf("Could not unmarshal json : %v", err)
	}
	if status.BuildNumber != *expectedVersion {
		log.Fatalf("Running build number '%v' does not match expected '%v'", status.BuildNumber, *expectedVersion)
	}
	os.Exit(0)
}
