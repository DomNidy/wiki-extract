package cmd

import (
	"fmt"
	"wiki-extract/util"

	"github.com/spf13/cobra"
)

func init() {
	// Add subcommands to url command
	urlCmd.AddCommand(urlAddCmd)
	urlCmd.AddCommand(urlListCmd)
	urlCmd.AddCommand(urlClearCmd)

	rootCmd.AddCommand(urlCmd)
}

var urlCmd = &cobra.Command{
	Use:     "url",
	Aliases: []string{"urls", "u"},
	Short:   "Manage the URLs to extract text from",
	Long:    `Manage the Wikipedia URLs to extract text from.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("You are running wiki-extract %s\n", version)
		cmd.Help()
	},
}

var urlAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires at least one URL argument")
		}
		return nil
	}, Example: "  wkx url add https://en.wikipedia.org/wiki/Go_(programming_language) https://en.wikipedia.org/wiki/Python_(programming_language)",
	Short: "Add URL(s) to the extraction list",
	Long:  `Add URL(s) to the extraction list, you can add multiple urls separated by spaces.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Split the args by spaces (in case the user entered multiple urls in one string)
		var splitUrls []string = util.SplitURLInput(args)
		var validUrls []string = util.RemoveDuplicateStringsFromArray(util.ValidateWikipediaURLS(splitUrls))
		fmt.Printf("Trying to add %d URL(s) to extraction list...\n", len(splitUrls))

		if !(len(validUrls) == len(splitUrls)) {
			fmt.Printf("%d invalid or duplicate URL(s) not added to the extraction list\n", len(splitUrls)-len(validUrls))
		}

		fmt.Printf("%d URL(s) parsed successfully\n", len(validUrls))
		for i, url := range validUrls {
			fmt.Println(i, "-", url)
		}

		fmt.Println("\nAdding URL(s) to extraction list...")
		addedURLCount := util.WriteURLExtractionList(validUrls)

		if !(addedURLCount == len(validUrls)) {
			fmt.Printf("%d URL(s) already present in the extraction list and were not added (duplicates)\n", len(validUrls)-addedURLCount)
		}

		if addedURLCount < 0 {
			fmt.Printf("Removed %d duplicate URL(s) from the extraction list\n", addedURLCount*-1)
		}

		fmt.Printf("%d URL(s) added to extraction list", addedURLCount)
	},
}

var urlListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all URLs in the extraction list",
	Long:    `List all URLs in the extraction list`,
	Run: func(cmd *cobra.Command, args []string) {
		var urls []string = util.ReadURLExtractionList()

		for i, url := range urls {
			fmt.Println(i, "-", url)
		}

		fmt.Printf("Total of %d URL(s) in extraction list\n", len(urls))
	},
}

var urlClearCmd = &cobra.Command{
	Use:     "clear",
	Aliases: []string{"delete", "del"},
	Short:   "Delete the extraction list",
	Long:    `Delete the extraction list json file`,
	Run: func(cmd *cobra.Command, args []string) {
		util.ClearURLExtractionList()
		fmt.Println("Extraction list cleared")
	},
}
