package cmd

import (
	"github.com/agastiya/tiyago/database/migrations"
	"github.com/spf13/cobra"
)

func (c *Cmd) MigrateMigrationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			migrations.Up(c.DB)
		},
	}
	return cmd
}
