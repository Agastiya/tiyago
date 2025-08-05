package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/agastiya/tiyago/pkg/helper/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func (c *Cmd) MakeFeatureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make:feature [name]",
		Short: "Create new feature",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			lowerName := strings.ToLower(name)
			funcName := utils.ToFuncName(name)

			// Create Model
			modelFileName := fmt.Sprintf("models/%s.go", lowerName)
			shouldWrite := true

			if FileExists(modelFileName) {
				answer := AskForConfirmation("This model file already exists. Do you want to override it? (y/N): ")
				shouldWrite = answer
			}

			if shouldWrite {
				if err := RenderAndWriteFile(lowerName, funcName, "model", modelFileName); err != nil {
					log.Fatal().Msgf("❌ Error writing file: %v", err)
				}
			}

			// Create Dto
			dtoFileName := fmt.Sprintf("dto/%s.go", lowerName)
			if err := RenderAndWriteFile(lowerName, funcName, "dto", dtoFileName); err != nil {
				log.Fatal().Msgf("%v", err)
			}

			// Create Repository
			repoPath := fmt.Sprintf("repository/%s", lowerName)
			if !FileExists(repoPath) {
				if err := os.Mkdir(repoPath, 0755); err != nil {
					log.Fatal().Msgf("❌ Error creating repository feature folder: %v", err)
				}
			}

			repoName := fmt.Sprintf("%s/%s.go", repoPath, lowerName)
			if err := RenderAndWriteFile(lowerName, funcName, "repository", repoName); err != nil {
				log.Fatal().Msgf("%v", err)
			}

			repoModelName := fmt.Sprintf("%s/model.go", repoPath)
			if err := RenderAndWriteFile(lowerName, funcName, "repository_model", repoModelName); err != nil {
				log.Fatal().Msgf("%v", err)
			}

			// Create Service
			servicePath := fmt.Sprintf("service/%s", lowerName)
			if !FileExists(servicePath) {
				if err := os.Mkdir(servicePath, 0755); err != nil {
					log.Fatal().Msgf("❌ Error creating repository feature folder: %v", err)
				}
			}

			serviceName := fmt.Sprintf("%s/%s.go", servicePath, lowerName)
			if err := RenderAndWriteFile(lowerName, funcName, "service", serviceName); err != nil {
				log.Fatal().Msgf("%v", err)
			}

			serviceModelName := fmt.Sprintf("%s/model.go", servicePath)
			if err := RenderAndWriteFile(lowerName, funcName, "service_model", serviceModelName); err != nil {
				log.Fatal().Msgf("%v", err)
			}

			// Create Controller
			controllerName := fmt.Sprintf("controller/%s.go", lowerName)
			if err := RenderAndWriteFile(lowerName, funcName, "controller", controllerName); err != nil {
				log.Fatal().Msgf("%v", err)
			}

			log.Info().Msgf("✅ Feature %s Created", funcName)
		},
	}
	return cmd
}

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

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func AskForConfirmation(prompt string) bool {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.ToLower(strings.TrimSpace(answer))

	return answer == "y" || answer == "yes"
}
