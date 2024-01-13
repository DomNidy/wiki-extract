package cmd

import (
	"fmt"
	"path"
	"strings"
	"wiki-extract/util"

	"github.com/spf13/cobra"
)

func init() {
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

		for i, url := range urls {
			fmt.Print("Extracting text from URL ", i, " - ", url, "\n")
			wikipediaPageRawText := util.RequestURL(url)

			// Remove invalid characters from filename (so that it can be saved)
			filename := path.Base(url)
			filename = strings.ReplaceAll(filename, "/", "_")
			filename = strings.ReplaceAll(filename, ".", "_")

			util.WriteWikipediaRawText(filename, wikipediaPageRawText)

		}
	},
}
