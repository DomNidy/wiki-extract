package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wiki-extract",
	Long:  `All software has versions. This is wiki-extract's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("You are running wiki-extract %s\n", version)
	},
}
