package commands

import (
	"fmt"
	"math"
	"strings"

	"github.com/AfshinJalili/gorandom/internal/articles"
	"github.com/AfshinJalili/gorandom/internal/history"
	"github.com/AfshinJalili/gorandom/internal/ui"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show progress",
	Run:   runStats,
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func runStats(cmd *cobra.Command, args []string) {
	readUrls, err := historyStore.GetReadUrls()
	if err != nil {
		printHistoryLoadError(cmd, err)
		return
	}

	streak, err := historyStore.CalculateStreak()
	if err != nil {
		printError(cmd, fmt.Sprintf("Could not calculate streak: %v", err), "check your history file")
		return
	}
	streakIcon := "🔥"
	if streak == 0 {
		streakIcon = "🧊"
	} else if streak > 7 {
		streakIcon = "⚡"
	}

	cmd.Println()
	cmd.Println(ui.TitleStyle.Render("📊 Your Go Learning Progress"))
	cmd.Println()

	streakMsg := fmt.Sprintf("%s Current Streak: %d day(s)", streakIcon, streak)
	cmd.Println(ui.StatsStyle.Render(streakMsg))
	cmd.Println()

	totalRead := 0
	totalArticles := 0

	for _, source := range articles.Sources {
		// Filter articles by source
		sourceArticles := articles.FilterBySource(articles.Data, source)

		sourceRead := 0
		for _, a := range sourceArticles {
			if readUrls[a.URL] {
				sourceRead++
			}
		}

		totalRead += sourceRead
		totalArticles += len(sourceArticles)

		percent := 0.0
		if len(sourceArticles) > 0 {
			percent = float64(sourceRead) / float64(len(sourceArticles)) * 100
		}

		barLength := ui.ProgressBarWidth
		filled := int(math.Floor(percent / (100.0 / float64(barLength))))
		if filled > barLength {
			filled = barLength
		}

		barStr := strings.Repeat("█", filled) + strings.Repeat("░", barLength-filled)
		barRendered := ui.StatBarStyle.Render(barStr)

		sourceName := articles.FormatSource(source)

		line := fmt.Sprintf("%s %s %s",
			ui.StatLabelStyle.Render(sourceName),
			barRendered,
			ui.StatValueStyle.Render(fmt.Sprintf("%d/%d (%d%%)", sourceRead, len(sourceArticles), int(math.Round(percent)))),
		)
		cmd.Println(line)
	}

	cmd.Println(ui.SubtleStyle.Render(strings.Repeat("─", 50)))

	totalPercent := 0.0
	if totalArticles > 0 {
		totalPercent = float64(totalRead) / float64(totalArticles) * 100
	}

	cmd.Printf("%s %s %s\n",
		ui.StatLabelStyle.Render("Total"),
		strings.Repeat(" ", ui.ProgressBarWidth), // spacer for bar alignment
		ui.StatValueStyle.Render(fmt.Sprintf("%d/%d (%d%%)", totalRead, totalArticles, int(math.Round(totalPercent)))),
	)

	cmd.Println()
	path, err := history.GetHistoryFilePath()
	if err != nil {
		printError(cmd, fmt.Sprintf("Could not resolve history path: %v", err), "set GORANDOM_CONFIG_DIR to a writable folder")
		return
	}
	cmd.Println(ui.SubtleStyle.Render(fmt.Sprintf("History stored at: %s", path)))
	cmd.Println()
}
