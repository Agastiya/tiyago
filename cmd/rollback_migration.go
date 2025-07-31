package cmd

import (
	"github.com/agastiya/tiyago/database/migrations"
	"github.com/spf13/cobra"
)

func (c *Cmd) RollbackMigrationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollback",
		Short: "Rollback all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			migrations.Down(c.DB)
		},
	}
	return cmd
}
