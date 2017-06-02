package beegoutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// ReadDockerSecret reads content of docker secret's file
func ReadDockerSecret(filename string) string {
	secret, err := ioutil.ReadFile(fmt.Sprintf("/run/secrets/%s", filename))
	if err != nil {
		log.Fatalf("Unable to read Docker Secret file %s", filename)
	}
	return strings.TrimSpace(string(secret))
}
