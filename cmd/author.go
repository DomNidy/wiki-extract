package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authorCmd)
}

var authorCmd = &cobra.Command{
	Use:     "author",
	Aliases: []string{"a"},
	Short:   "Print out the author of wiki-extract (that's me! the guy writing this!)",
	Long:    "Print out the author of wiki-extract (that's me! the guy writing this!)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wiki-extract written by Dominic Nidy https://github.com/DomNidy")
	},
}
