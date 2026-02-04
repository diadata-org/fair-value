package utils

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetRequest performs a get request on @url and returns the response body
// as a slice of byte data.
func GetRequest(url string, requestHeaders map[string]string) ([]byte, int, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	for name, value := range requestHeaders {
		req.Header.Set(name, value)
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Close response body after function
	defer func() {
		cerr := response.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if response.StatusCode != 200 {
		return []byte{}, response.StatusCode, fmt.Errorf("Get HTTP Response Error %d for url %s", response.StatusCode, url)
	}

	// Read the response body
	XMLdata, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return []byte{}, response.StatusCode, err
	}

	return XMLdata, response.StatusCode, err
}
