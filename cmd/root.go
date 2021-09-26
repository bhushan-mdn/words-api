package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "words-api",
	Short: "Words API",
	Long: `Commands:
server  - Run the server
migrate - Run database migrations`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
