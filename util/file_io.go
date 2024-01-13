package util

import (
	"encoding/json"
	"fmt"
	"os"
)

// Add urls to the extraction list json file
// Returns the amount of new urls that were added
func WriteURLExtractionList(urls []string) int {
	// Try to read the contents of the file if it exists, so that we can append to it
	var existingUrls []string = ReadURLExtractionList()

	// Append existing urls to the new urls
	urls = append(existingUrls, urls...)

	// Remove duplicate urls from the list
	urls = RemoveDuplicates(urls)

	file, _ := json.MarshalIndent(urls, "", " ")
	_ = os.WriteFile("urls.json", file, 0644)

	return len(urls) - len(existingUrls)
}

// Read from the extraction list json file
func ReadURLExtractionList() []string {
	var urls []string
	file, _ := os.ReadFile("urls.json")
	_ = json.Unmarshal([]byte(file), &urls)
	return urls
}

// Delete the extraction list json file
func ClearURLExtractionList() {
	_ = os.Remove("urls.json")
}

func RemoveDuplicates(urls []string) []string {
	set := make(map[string]struct{}, len(urls))

	for url := range urls {
		set[urls[url]] = struct{}{}
	}

	uniqueUrls := make([]string, 0, len(set))
	for url := range set {
		uniqueUrls = append(uniqueUrls, url)
	}

	return uniqueUrls
}

// TODO: Implement reading output path from config file
// Write the raw text to from a wikipedia url to file
func WriteWikipediaRawText(filename string, text string) {
	_ = os.MkdirAll("output", 0755)
	_ = os.WriteFile(fmt.Sprintf("output/%s_raw.txt", filename), []byte(text), 0644)
}
