package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Add urls to the extraction list json file
// Returns the amount of new urls that were added
func WriteURLExtractionList(urls []string) int {
	var urlFilePath string = viper.GetString("urlFilePath")

	// Try to read the contents of the file if it exists, so that we can append to it
	var existingUrls []string = ReadURLExtractionList()

	// Append existing urls to the new urls
	urls = append(existingUrls, urls...)

	// Remove duplicate urls from the list
	urls = RemoveDuplicates(urls)

	file, _ := json.MarshalIndent(urls, "", " ")
	_ = os.WriteFile(urlFilePath, file, 0644)

	return len(urls) - len(existingUrls)
}

// Read from the extraction list json file
func ReadURLExtractionList() []string {
	var urlFilePath string = viper.GetString("urlFilePath")
	var urls []string
	file, _ := os.ReadFile(urlFilePath)
	_ = json.Unmarshal([]byte(file), &urls)
	return urls
}

// Delete the extraction list json file
func ClearURLExtractionList() {
	var urlFilePath string = viper.GetString("urlFilePath")
	_ = os.Remove(urlFilePath)
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

// Write the raw text to from a wikipedia url to file
func WriteWikipediaRawText(filename string, text string) {
	outputDir := viper.GetString("outputDir")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	if err := os.WriteFile(filepath.Join(outputDir, fmt.Sprintf("%s_raw.txt", filename)), []byte(text), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
