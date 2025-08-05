package cmd

import (
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type Cmd struct {
	DB *gorm.DB
}

func (c *Cmd) RegisterCommands(root *cobra.Command) {
	root.AddCommand(
		c.MakeMigrationCmd(),
		c.MigrateMigrationCmd(),
		c.RollbackMigrationCmd(),
		c.MakeFeatureCmd(),
	)
}

func RootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "tiyago",
		Short: "Tiyago CLI Tool",
	}
	return rootCmd
}
