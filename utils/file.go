package utils

import (
	"io"
	"os"
	"os/user"
)

// readFile reads a gitcoin submission json file and returns the slice of items.
func ReadFile(filename string) (data []byte, err error) {
	var (
		jsonFile *os.File
	)
	jsonFile, err = os.Open(configFileConnectors(filename))
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

func configFileConnectors(filename string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if dir == "/root" || dir == "/home" {
		return "https://github.com/diadata-org/fair-value/config/" + filename
	}

	return os.Getenv("GOPATH") + "/src/github.com/diadata-org/fair-value/config/" + filename
}
