package cmd

import (
	"github.com/bhushan-mdn/words-api/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate data to new database",
	// Long:  `Something longer`,
	Run: func(cmd *cobra.Command, args []string) {
		database.Migrate()
	},
}
