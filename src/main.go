package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var username string
var password string
var prefix string

// Example:
// APPDL_SOMEAPP=http://example.com/filename.zip url-downloader -username=<BasicAuthUsername> -password=<BasicAuthPassword>
func main() {
	flag.StringVar(&username, "username", "anonymous", "BasicAuth username.")
	flag.StringVar(&password, "password", "anonymous", "BasicAuth password.")
	flag.StringVar(&prefix, "prefix", "APPDL_", "environment variable prefix for downloading application(s).")
	flag.Parse()

	getAppByEnv(prefix)
}

// Scan environment (by prefix), and download files
func getAppByEnv(prefix string) {
	for _, s := range os.Environ() {
		kv := strings.SplitN(s, "=", 2)
		if strings.HasPrefix(kv[0], prefix) {
			getFile(kv[1])
		}
	}
}

// Download from url, using basic-auth
func getFile(url string) {
	fmt.Printf("[DOWNLOADING]\n- SOURCE: %q\n- TARGET: %s\n", url, filepath.Base(url))

	// Define http-request
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}
	req.Header.Set("User-Agent", "Golang / https://github.com/Kreditorforeningens-Driftssentral-DA/url-downloader")
	req.SetBasicAuth(username, password)

	// Define HTTP-Client
	c := http.Client{
		Timeout: time.Second * 5,
	}

	// Execute http-request
	resp, respErr := c.Do(req)
	if respErr != nil {
		log.Fatal(respErr)
	}

	// Check if any data in response
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Create output-file
	dlPath := os.Getenv(prefix + "DLPATH")
	if len(dlPath) == 0 {
		dlPath = "/tmp"
	}
	out, outErr := os.Create(dlPath + filepath.Base(url))
	if outErr != nil {
		log.Fatal(outErr)
	}
	defer out.Close()

	// Write response-body to output-file
	_, outErr = io.Copy(out, resp.Body)
}
