package util

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Try to split the url by spaces, if it fails then it is a single url
// This is because the user may have entered in their urls with a quoted string
// e.g. wkx url add "https://en.wikipedia.org/wiki/Go_(programming_language)"
// This would result in the url being passed as a single string
func SplitURLInput(args []string) []string {
	var urls []string

	for _, url := range args {
		urls = append(urls, strings.Split(url, " ")...)
	}

	return urls
}

// Given an array of urls, return all elements of the array which are valid wikipedia urls
func ValidateWikipediaURLS(args []string) []string {
	var urls []string

	// Regex to match valid wikipedia URLs
	regex := `^(http(s)?://)?([a-z]+\.)?wikipedia\.(org|com)(\/)?`

	for _, url := range args {

		if regexp.MustCompile(regex).MatchString(url) {
			urls = append(urls, url)
		}
	}

	return urls
}

// Parse text from raw html
func ParseContentfulTextFromHTML(html string) []string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Println("Error parsing html:", err)
		panic(err)
	}

	// initialize an array to store the parsed text
	var parsedText []string = []string{}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		parsedText = append(parsedText, strings.TrimSpace(s.Text()))
	})

	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		parsedText = append(parsedText, strings.TrimSpace(s.Text()))
	})

	return parsedText
}

// Function which parses related links from the raw html
func ParseRelatedLinksFromHTML(html string, includeNonWikipediaLinks bool) []string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Println("Error parsing html:", err)
		panic(err)
	}

	var relatedLinks []string = []string{}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// Read href attribute from link
		href, exists := s.Attr("href")
		if exists {
			var url string
			// Check if the link is a relative link
			// We do this by checking if the link starts with a forward slash
			// Or if the gets returned from the ValidateWikipediaURLS function
			if strings.HasPrefix(href, "/") {
				url = path.Join("https://en.wikipedia.org", href)
			} else if len(ValidateWikipediaURLS([]string{href})) != 0 {
				url = href
			} else if !includeNonWikipediaLinks {
				// If the link is not a relative link and we don't want to include non-wikipedia links, then we skip it
				return
			}

			relatedLinks = append(relatedLinks, path.Join(url))
		}
	})

	return relatedLinks
}

// Returns a map of words and their frequency in the text
func CreateArticleDictionaryFromContentfulText(contentfulText []string, cleanLevel int) map[string]int {
	// join the text into a single string
	contentfulText = []string{strings.Join(contentfulText, " ")}
	var dictionary map[string]int = make(map[string]int)
	// split the text by spaces
	var splitText []string = strings.Split(contentfulText[0], " ")
	fmt.Println(splitText[0], "split text")
	fmt.Println(len(splitText[0]), "length of split text")
	for _, word := range splitText {
		// convert the word to lowercase
		word = _CleanString(strings.ToLower(word), cleanLevel)

		// if the word is not empty, then add it to the dictionary
		if word != "" {
			dictionary[word]++
		}
	}

	return dictionary
}

// Attempt to 'clean' a string (remove latex expressions/other non contentful text)
// cleanLevel 0 = no cleaning, cleanLevel 3 = maximum cleaning
// TODO: Improve the naming conventions here, regex, control flow, etc.
func _CleanString(text string, cleanLevel int) string {
	matchFullyNumericStrings := regexp.MustCompile(`^[0-9]+$`)           // Matches an entire string if it is fully numeric
	matchUploadWikimedia := regexp.MustCompile(`.*uploadwikimediaorg.*`) // Matches an entire string if it starts with the specified chars
	matchStringThatHasBackslash := regexp.MustCompile(`.*\\.*`)          // Matches an entire string if it has a backslash (removes a lot of latex expressions)

	level1Regex := regexp.MustCompile(`[',.;:!?"\s]`)                   // Level 1 : Matches all punctuation
	level2Regex := regexp.MustCompile(`[\(\[\{\|\)\]\}\(\[\{\|\)\]\}]`) // Level 2 Matches all opening and closing brackets
	level3Regex := regexp.MustCompile(`^[\\_=-^\/].*$`)                 // Level 3 Matches an entire string if it starts with the specified chars

	var cleanedText string = text

	switch {
	case cleanLevel >= 3:
		cleanedText = level3Regex.ReplaceAllString(cleanedText, "")
		fallthrough
	case cleanLevel >= 2:
		cleanedText = level2Regex.ReplaceAllString(cleanedText, "")
		cleanedText = matchFullyNumericStrings.ReplaceAllString(cleanedText, "")
		fallthrough
	case cleanLevel >= 1:
		cleanedText = level1Regex.ReplaceAllString(cleanedText, "")
		cleanedText = matchUploadWikimedia.ReplaceAllString(cleanedText, "")
		cleanedText = matchFullyNumericStrings.ReplaceAllString(cleanedText, "")
		cleanedText = matchStringThatHasBackslash.ReplaceAllString(cleanedText, "")
	}

	// Replace all matches with an empty string
	return cleanedText
}
