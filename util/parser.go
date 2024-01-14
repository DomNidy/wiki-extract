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
		parsedText = append(parsedText, strings.TrimSpace(s.Text()))
	})

	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		parsedText = append(parsedText, strings.TrimSpace(s.Text()))
	})

	return parsedText
}

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
			// Or if the gets returned from the ParseURLS function
			if strings.HasPrefix(href, "/") {
				url = path.Join("https://en.wikipedia.org", href)
			} else if len(ParseURLS([]string{href})) != 0 {
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
