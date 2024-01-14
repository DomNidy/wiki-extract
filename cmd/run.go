package cmd

import (
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
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"extract"},
	Short:   "Extract the text from the URLs in the extraction list",
	Long:    "Extract the text from the URLs in the extraction list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wiki-extract written by Dominic Nidy https://github.com/DomNidy")

		urls := util.ReadURLExtractionList()

		if len(urls) == 0 {
			fmt.Println("No URLs in extraction list, cannot run extractor")
			fmt.Println("Try adding some with 'url add <url>'")
			return
		}

		queryDelay := viper.GetUint32("delay")
		for i, url := range urls {
			fmt.Print("Extracting text from URL ", i, " - ", url, "\n")
			wikipediaPageRawText := util.RequestURL(url)

			// Remove invalid characters from filename (so that it can be saved)
			filename := path.Base(url)
			filename = strings.ReplaceAll(filename, "/", "_")
			filename = strings.ReplaceAll(filename, ".", "_")

			// Parse text from the html
			parsedText := util.ParseTextFromHTML(wikipediaPageRawText)

			// Write the raw html to file
			util.WriteWikipediaRawText(filename, wikipediaPageRawText)
			// Write the parsed text to file
			util.WriteWikipediaParsedText(filename, parsedText)

			time.Sleep(time.Millisecond * time.Duration(queryDelay))
		}
	},
}
