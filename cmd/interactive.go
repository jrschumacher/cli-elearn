package cmd

import (
	"github.com/jrschumacher/cli-learning/tui"
	"github.com/spf13/cobra"
)

var (
	interactiveCmd = &cobra.Command{
		Use:     "interactive",
		Aliases: []string{"i"},
		Short:   "Start the interactive mode",
		Long:    `Start the interactive mode`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := tui.Start(); err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
