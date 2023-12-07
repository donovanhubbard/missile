/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
  "fmt"
  "regexp"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "missle",
	Short: "An interactive CLI for interfacing with memcached",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Args: func(cmd *cobra.Command, args []string) error {
    // Optionally run one of the validators provided by cobra
    if err := cobra.ExactArgs(1)(cmd, args); err != nil {
        return err
    }

    match, _ := regexp.MatchString(".+:[0-9]+", args[0])

    if !match {
      return fmt.Errorf("Must specify host and port delimited with ':' character. e.g. localhost:11211")
    }
    return nil
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Running bare command")
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


