package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Define the 'config' command and its subcommands
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
}

type ConfigUpdate struct {
	apiUrl      string
	accountsUrl string
}

// Subcommand for 'config show'
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show configuration settings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api:", viper.GetString("endpoint.api"))
		fmt.Println("accounts:", viper.GetString("endpoint.accounts"))
	},
}

// Subcommand for 'config update'
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a configuration setting",
	Run: func(cmd *cobra.Command, args []string) {
		recoverWithLog()
		var flags = &ConfigUpdate{}

		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Name == "apiUrl" {
				flags.apiUrl = f.Value.String()
			}
			if f.Name == "accountsUrl" {
				flags.accountsUrl = f.Value.String()
			}
		})

		if flags.apiUrl == "" && flags.accountsUrl == "" {
			fmt.Println("apiUrl or accountsUrl is required")
			fmt.Println(cmd.Flags().FlagUsages())
			return
		}

		if flags.apiUrl != "" {
			viper.Set("endpoint.api", flags.apiUrl)
			fmt.Println("Updating 'api' configuration:", flags.apiUrl)
		}
		if flags.accountsUrl != "" {
			viper.Set("endpoint.accounts", flags.accountsUrl)
			fmt.Println("Updating 'accounts' configuration:", flags.accountsUrl)
		}

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Error updating configuration:", err)
			return
		}
	},
}

func init() {
	// Set up Viper configuration
	// Set a default value for 'host' configuration
	viper.SafeWriteConfigAs("/cli-data/config.yaml")
	viper.SetDefault("endpoint.api", "https://api.ente.io")
	viper.SetDefault("endpoint.accounts", "https://accounts.ente.io")

	// Add 'config' subcommands to the root command
	rootCmd.AddCommand(configCmd)

	// Add flags to the 'config store' and 'config update' subcommands
	updateCmd.Flags().StringP("apiUrl", "a", "", "Update the 'api' configuration")
	updateCmd.Flags().StringP("accountsUrl", "A", "", "Update the 'accounts' configuration")
	// Mark 'host' flag as required for the 'update' command

	// Add 'config' subcommands to the 'config' command
	configCmd.AddCommand(showCmd, updateCmd)
}
