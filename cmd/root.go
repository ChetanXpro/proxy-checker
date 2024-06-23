package cmd

import (
	"fmt"
	"os"
	"proxxy/checker"

	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFile string
	threads    int
)

var rootCmd = &cobra.Command{
	Use:   "proxxy",
	Short: "Proxxy is a simple proxy checker",
	Long:  `Proxxy is a CLI tool that checks the validity of proxy from a given input file and outputs the results to a specified file.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := checker.CheckProxies(inputFile, outputFile, threads)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file containing proxies (required)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for results (required)")
	rootCmd.Flags().IntVarP(&threads, "thread", "t", 10, "Number of threads to use")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
