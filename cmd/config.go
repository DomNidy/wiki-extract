package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	configCmd.AddCommand(configReadCmd)
	configCmd.AddCommand(configResetCmd)

	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the configuration file",
	Long:  `Manage the configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var configReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Print the configuration file settings",
	Long:  `Print the configuration file settings`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuration file contents:")
		fmt.Println("Reading config file from:", viper.ConfigFileUsed())
		for setting := range viper.GetViper().AllSettings() {
			fmt.Println(setting+":", viper.Get(setting))
		}
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Deletes the configuration file (will be recreated on next run)",
	Long:  `Deletes the configuration file (will be recreated on next run)`,
	Run: func(cmd *cobra.Command, args []string) {
		// Delete the config file
		if err := os.Remove(viper.ConfigFileUsed()); err != nil {
			fmt.Println("Error deleting config file:\n", err)
			return
		} else {
			fmt.Println("Successfully deleted config file")
		}
	},
}
