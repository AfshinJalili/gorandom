package commands

import (
	"os"

	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/spf13/cobra"
)

var historyStore history.Store = history.DefaultStore

var rootCmd = &cobra.Command{
	Use:   "gorandom",
	Short: "Get a random Go article",
	Long:  `gorandom is a CLI tool to discover random Go articles, tutorials, and documentation.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default to random command if no subcommand provided
		runRandom(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
	rootCmd.Flags().BoolP("any", "a", false, "Include already-read articles in random pick")
	rootCmd.Flags().StringP("source", "s", "", "Filter by source (docs, tour, gobyexample, pkg, blog)")
}
