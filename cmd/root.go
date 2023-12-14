/*
Copyright © 2023 Donovan Hubbard
*/
package cmd

import (
	"fmt"
	// "github.com/donovanhubbard/missle/backend"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/donovanhubbard/missile/models"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strconv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "missle <host:port>",
	Short: "An interactive CLI for interfacing with memcached",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		re := regexp.MustCompile(`.+:(\d+)$`)

		for _, arg := range args {
			matches := re.FindStringSubmatch(arg)
			if len(matches) < 2 {
				return fmt.Errorf("Must specify host and port delimited with ':' character. e.g. localhost:11211")
			}

			portString := matches[1]
			portNumber, err := strconv.Atoi(portString)

			if err != nil {
				return fmt.Errorf("Invalid port number %s", portString)
			}

			if portNumber < 1 || portNumber > 65535 {
				return fmt.Errorf("Invalid port number %d", portNumber)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running bare command")
		m := models.New(args)
		p := tea.NewProgram(m)
		_, err := p.Run()

		if err != nil {
			fmt.Println("There has been an error: %v", err)
			os.Exit(1)
		}
		// client.Connect(args...)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.missle.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
