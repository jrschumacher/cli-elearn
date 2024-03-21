package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "app is a CLI application",
		Long:  `app is a CLI application`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
