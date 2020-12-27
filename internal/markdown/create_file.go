package markdown

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/rescDev/actions-docs/pkg/tpl"
	"github.com/rescDev/actions-docs/pkg/util"
)

const (
	// File name which is used when creating the documentation file
	createFileName = "README.md"

	// Template for header + introduction
	introTpl = `
	# {{ .Name }}
	{{ printf "\n" }}
	{{- .Description }}.
	`
)

// Creates a new file (README.md) with the generated content
// Overwrites an already existing README.md if wanted (asking for user confirmation via custom prompt)
func createDocsFile(meta *ActionMeta, content string) error {
	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to determine current directory, can not read metadata file: %s", err)
	}

	// Write to file
	// Appends the GitHub action name as top-level header and the description as introduction to the already generated content
	tt, err := tpl.CreateTemplate("intro", introTpl)
	if err != nil {
		return err
	}

	introContent, err := tpl.RenderTemplate(tt, meta)
	if err != nil {
		return err
	}

	// Check if file already exists
	filepath := fmt.Sprintf("%s/%s", dir, createFileName)
	if _, err := os.Stat(filepath); err == nil {
		fmt.Println()

		prompt := promptui.Prompt{
			Label: "A README file already exists. Do you wish to override the file ? (y|n)",
		}

		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("could not handle prompt input: %s", err)
		}

		// Abort if file should not be overwritten
		if result == "n" {
			fmt.Print("\nFile will not be overwritten, aborting due to user input\n\n")
			return nil
		}
	}

	// Create and write to file
	readmeContent := strings.TrimPrefix(util.Normalize(introContent+content), "\n")
	err = ioutil.WriteFile(filepath, []byte(readmeContent), 0644)
	if err != nil {
		return fmt.Errorf("unable to write README file: %s", err)
	}

	fmt.Print("\nSuccessfully created README.md :)\n\n")
	return nil
}
