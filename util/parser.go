package util

import (
	"fmt"
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

// Parse out valid wikipedia URLs from a list of strings
func ParseURLS(args []string) []string {
	var urls []string

	// Regex to match valid wikipedia URLs
	regex := `^(http(s)?://)?([a-z]+\.)?wikipedia\.(org|com)(\/)?`

	for _, url := range args {

		if regexp.MustCompile(regex).MatchString(url) {
			urls = append(urls, url)
		} else {
			fmt.Println("Invalid URL:", url)
		}
	}

	return urls
}

// Parse text from raw html
func ParseTextFromHTML(html string) []string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Println("Error parsing html:", err)
		panic(err)
	}

	var parsedText []string = []string{}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		parsedText = append(parsedText, s.Text())
	})

	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		parsedText = append(parsedText, s.Text())
	})

	return parsedText
}
