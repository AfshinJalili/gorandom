package commands

import (
	"bytes"
	"testing"
)

func TestRootCommand(t *testing.T) {
	// We can test the root command by redirecting its output
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"--help"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("rootCmd.Execute() failed: %v", err)
	}

	if b.Len() == 0 {
		t.Error("Expected help output, got nothing")
	}
}

func TestStatsCommandHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"stats", "--help"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("stats command help failed: %v", err)
	}

	if b.Len() == 0 {
		t.Error("Expected help output for stats, got nothing")
	}
}
