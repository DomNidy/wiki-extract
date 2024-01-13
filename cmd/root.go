package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// The version of the program. This is set at compile time using: go build -ldflags "-X root.version=1.0.0"
var version string = "1.0.0"

func init() {
	// TODO: Add config file support
	// cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   "wkx",
	Short: "wiki-extract is a tool for extracting text from Wikipedia pages",
	Long:  `wiki-extract is a CLI program for extracting text from Wikipedia pages. Built using Golang & Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("======- Welcome to wiki-extract %s -======\n\n", version)
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
