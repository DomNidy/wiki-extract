package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Add urls to the extraction list json file
// Returns the amount of new urls that were added
// This function will automatically remove duplicates from the passed urls array, also removes them from the file
func WriteURLExtractionList(urls []string) int {
	var urlFilePath string = viper.GetString("urlFilePath")

	// Try to read the contents of the file if it exists, so that we can append to it
	var existingUrls []string = ReadURLExtractionList()

	// Append existing urls to the new urls
	urls = append(existingUrls, urls...)

	// Remove duplicate urls from the list
	urls = RemoveDuplicateStringsFromArray(urls)

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

// Removes duplicate strings from an array
func RemoveDuplicateStringsFromArray(array []string) []string {
	set := make(map[string]struct{}, len(array))

	for url := range array {
		set[array[url]] = struct{}{}
	}

	uniqueUrls := make([]string, 0, len(set))
	for url := range set {
		uniqueUrls = append(uniqueUrls, url)
	}

	return uniqueUrls
}

// Write the raw HTML of a wikipedia page to file
func WriteWikipediaRawHTML(filename string, text string) {
	outputDir := viper.GetString("outputDir")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	if err := os.WriteFile(filepath.Join(outputDir, fmt.Sprintf("%s_raw.html", filename)), []byte(text), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// Write contentful text to file
func WriteWikipediaContentfulText(filename string, text []string) {
	outputDir := viper.GetString("outputDir")
	text = []string{strings.Join(text, "\n")}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	if err := os.WriteFile(filepath.Join(outputDir, fmt.Sprintf("%s_contentful.txt", filename)), []byte(text[0]), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// Write the related links to file
func WriteWikipediaRelatedLinks(filename string, links []string) {
	outputDir := viper.GetString("outputDir")
	links = []string{strings.Join(links, "\n")}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	if err := os.WriteFile(filepath.Join(outputDir, fmt.Sprintf("%s_related.txt", filename)), []byte(links[0]), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// Write article dictionary to file
func WriteArticleDictionary(filename string, dictionary map[string]int) {
	outputDir := viper.GetString("outputDir")
	dictionaryJson, _ := json.MarshalIndent(dictionary, "", " ")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	if err := os.WriteFile(filepath.Join(outputDir, fmt.Sprintf("%s_dictionary.json", filename)), dictionaryJson, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
