package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/agastiya/tiyago/pkg/helper/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func RenderAndWriteFile(lowerName, funcName, templateName, fileName string) error {
	content, err := utils.RenderTemplate(templateName, map[string]string{
		"LowerFeatureName": lowerName,
		"FuncName":         funcName,
	})
	if err != nil {
		return fmt.Errorf("[❌ Failed to render %s template] %v", templateName, err)
	}

	if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
		return fmt.Errorf("[❌ Failed creating %s file] %v", fileName, err)
	}

	return err
}

func (c *Cmd) MakeFeatureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make:feature [name]",
		Short: "Create new feature",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			lowerName := strings.ToLower(name)
			funcName := utils.ToFuncName(name)

			// Create Dto
			dtoFileName := fmt.Sprintf("dto/%s.go", lowerName)
			err := RenderAndWriteFile(lowerName, funcName, "dto", dtoFileName)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}

			// Create Repository
			repoPath := fmt.Sprintf("repository/%s", lowerName)
			_, err = os.Stat(repoPath)
			if os.IsNotExist(err) {
				if err := os.Mkdir(repoPath, 0755); err != nil {
					log.Fatal().Msgf("❌ Error creating repository feature folder: %v", err)
				}
			}

			repoName := fmt.Sprintf("%s/%s.go", repoPath, lowerName)
			err = RenderAndWriteFile(lowerName, funcName, "repository", repoName)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}

			repoModelName := fmt.Sprintf("%s/model.go", repoPath)
			err = RenderAndWriteFile(lowerName, funcName, "repository_model", repoModelName)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}

			// Create Service
			servicePath := fmt.Sprintf("service/%s", lowerName)
			_, err = os.Stat(servicePath)
			if os.IsNotExist(err) {
				if err := os.Mkdir(servicePath, 0755); err != nil {
					log.Fatal().Msgf("❌ Error creating service feature folder: %v", err)
				}
			}

			serviceName := fmt.Sprintf("%s/%s.go", servicePath, lowerName)
			err = RenderAndWriteFile(lowerName, funcName, "service", serviceName)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}

			serviceModelName := fmt.Sprintf("%s/model.go", servicePath)
			err = RenderAndWriteFile(lowerName, funcName, "service_model", serviceModelName)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}

			// Create Controller
			controllerName := fmt.Sprintf("controller/%s.go", lowerName)
			err = RenderAndWriteFile(lowerName, funcName, "controller", controllerName)
			if err != nil {
				log.Fatal().Msgf("%v", err)
			}

			log.Info().Msgf("✅ Feature %s Created", funcName)
		},
	}
	return cmd
}
