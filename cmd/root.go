package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The version of the program. This is set at compile time using: go build -ldflags "-X root.version=1.0.0"
var version string = "1.0.0"

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to initialize config")
		os.Exit(1)
	}

	viper.SetConfigName("wiki-extract-cfg")
	viper.SetConfigType("json")

	// Search for config file in the following directories
	configPath := filepath.Join(home, ".wiki-extract")
	viper.AddConfigPath(configPath)

	// Default config values
	viper.SetDefault("delay", 2500)                                         // Delay in milliseconds between requests to Wikipedia
	viper.SetDefault("outputDir", filepath.Join(configPath, "output"))      // Output directory for extracted text files
	viper.SetDefault("urlFilePath", filepath.Join(configPath, "urls.json")) // Path to the file containing the urls to extract text from

	if err := viper.ReadInConfig(); err != nil {
		// If no config file was found, create one
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, creating one...")

			// Ensure the config directory (.wiki-extract) exists
			if _, err := os.Stat(filepath.Join(home, ".wiki-extract")); os.IsNotExist(err) {
				os.Mkdir(filepath.Join(home, ".wiki-extract"), 0755)
			}

			viper.SafeWriteConfigAs(filepath.Join(configPath, "wiki-extract-cfg.json"))
		} else {
			// Config file was found, but another error was produced
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
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
