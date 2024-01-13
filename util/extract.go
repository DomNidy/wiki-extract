// Extract the text from the URLs in the extraction list
package util

import (
	"io"
	"log"
	"net/http"
)

// Request a URL and return the raw text
func RequestURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
