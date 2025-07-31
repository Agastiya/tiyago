package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/agastiya/tiyago/pkg/helper/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func (c *Cmd) MakeMigrationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make:migration [name]",
		Short: "Create a new migration file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			timestamp := time.Now().Format("2006_01_02_150405")
			filename := fmt.Sprintf("%s_%s.go", timestamp, name)
			path := fmt.Sprintf("database/migrations/%s", filename)

			content, err := utils.RenderTemplate("migration", map[string]string{
				"FuncName": utils.ToFuncName(name),
			})
			if err != nil {
				log.Fatal().Msgf("❌ Failed to render migration: %v", err)
			}

			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				log.Fatal().Msgf("❌ Error creating migration: %v", err)
			}

			log.Info().Msgf("✅ Migration created: %s", path)
		},
	}
	return cmd
}
