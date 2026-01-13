package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type GitHubContent struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

func ReadFile(filename string, remoteConfig bool, branchConfig string) ([]byte, error) {
	if remoteConfig {
		return readFromRemote(filename, branchConfig)
	}
	return readFromFile(filename)
}

func readFromRemote(filename string, branchConfig string) (data []byte, err error) {

	URL := "https://api.github.com/repos/diadata-org/fair-value/contents/config/" + filename
	if branchConfig != "" {
		URL += "?ref=" + url.QueryEscape(branchConfig)
	}

	req, _ := http.NewRequest("GET", URL, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("GitHub API error: %s\n%s", resp.Status, string(body))
		return
	}

	var gh GitHubContent
	if err = json.NewDecoder(resp.Body).Decode(&gh); err != nil {
		return
	}
	data, err = base64.StdEncoding.DecodeString(gh.Content)
	return

}

// readFile reads a gitcoin submission json file and returns the slice of items.
func readFromFile(filename string) (data []byte, err error) {
	var (
		jsonFile *os.File
	)
	path := os.Getenv("GOPATH") + "/src/github.com/diadata-org/fair-value/config/" + filename
	jsonFile, err = os.Open(path)
	if err != nil {
		return
	}
	defer func() {
		cerr := jsonFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	return io.ReadAll(jsonFile)

}
