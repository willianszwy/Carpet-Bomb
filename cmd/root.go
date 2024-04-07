/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"FC-stress-test/internal/loadtest"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var (
	target      string
	requests    int
	concurrency int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loadtest-test",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		url, err := url.Parse(target)
		if err != nil {
			return
		}
		test := loadtest.NewTest(url, requests, concurrency)
		test.Run()
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.FC-loadtest-test.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&target, "url", "u", "", "URL to be tested")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 1000, "Number of requests to run the test")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Concurrent users")
}
