package cmd

import (
	"errors"
	"fmt"
	"path"
	"strings"
	"time"
	"wiki-extract/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	runCmd.Flags().Uint32P("delay", "d", 2000, "Delay in milliseconds between requests to Wikipedia")
	runCmd.Flags().Uint8P("cleanLevel", "c", 1, "An integer in the closed interval [0,3] that determines the cleaning 'strength' to use when parsing text")

	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"extract"},
	Short:   "Extract the text from the URLs in the extraction list",
	Long:    "Extract the text from the URLs in the extraction list",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Ensure that the clean level is between 0 and 3
		cleanLevel, err := cmd.Flags().GetUint8("cleanLevel")
		if err != nil || cleanLevel > 3 {
			return errors.New("cleanLevel must be between 0 and 3")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wiki-extract written by Dominic Nidy https://github.com/DomNidy")

		urls := util.ReadURLExtractionList()

		if len(urls) == 0 {
			fmt.Println("No URLs in extraction list, cannot run extractor")
			fmt.Println("Try adding some with 'url add <url>'")
			return
		}

		cleanLevel, err := cmd.Flags().GetUint8("cleanLevel")
		if err != nil {
			fmt.Println("Error parsing clean level, defaulting to 1", err)
			cleanLevel = 1
		}

		// Override the delay in config file if the user has passed a delay flag
		queryDelay, err := cmd.Flags().GetUint32("delay")
		if err != nil {
			queryDelay = viper.GetUint32("delay")
		}

		fmt.Println("Running with clean level", cleanLevel)

		for i, url := range urls {
			fmt.Print("Extracting text from URL ", i, " - ", url, "\n")
			wikipediaPageRawText := util.RequestURL(url)

			// Remove invalid characters from filename (so that it can be saved)
			filename := path.Base(url)
			filename = strings.ReplaceAll(filename, "/", "_")
			filename = strings.ReplaceAll(filename, ".", "_")

			// Parse text from the html
			contentfulText := util.ParseContentfulTextFromHTML(wikipediaPageRawText)
			relatedLinks := util.ParseRelatedLinksFromHTML(wikipediaPageRawText, false)
			articleDictionary := util.CreateArticleDictionaryFromContentfulText(contentfulText, int(cleanLevel))

			// Write the raw html to file
			util.WriteWikipediaRawHTML(filename, wikipediaPageRawText)
			// Write the parsed text to file
			util.WriteWikipediaContentfulText(filename, contentfulText)
			// Find related links and write them to file
			util.WriteWikipediaRelatedLinks(filename, relatedLinks)
			// Write article dictionary to file
			util.WriteArticleDictionary(filename, articleDictionary)

			time.Sleep(time.Millisecond * time.Duration(queryDelay))
		}
	},
}
