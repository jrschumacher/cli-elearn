package cmd

import (
	"github.com/jrschumacher/cli-learning/tui"
	"github.com/spf13/cobra"
)

var (
	designSystemCmd = &cobra.Command{
		Use:     "design-system",
		Aliases: []string{"ds"},
		Short:   "Start the design system mode",
		Long:    `Start the design system mode`,
		Run: func(cmd *cobra.Command, args []string) {
			tui.StartDesignSystem()
		},
	}
)

func init() {
	rootCmd.AddCommand(designSystemCmd)
}
