package markdown

import (
	"fmt"

	"github.com/rescDev/actions-docs/pkg/util"
	"github.com/urfave/cli/v2"
)

var (
	// Allowed and currently usable markdown formats
	allowedFormats = []string{"table", "list"}
)

// Command flags
var (
	// When using this flag, the tool will create a new file with the generated content
	// File name will be 'README.md'
	createFileFlag = &cli.BoolFlag{
		Name:     "create-file",
		Required: false,
		Value:    false,
	}

	// Define the format of the generated documentation about the Action inputs and outputs
	// Allowed values:
	// 'table' : creates a markdown table
	// 'list'  : creates a markdown list
	formatFlag = &cli.StringFlag{
		Name:     "format",
		Required: false,
		Value:    allowedFormats[0], // 'table' as default value
	}
)

// MarkdownCommand descibes the CLI command for creating the Markdown documentation
// Usage: 'actions-docs markdown ...' or 'actions-docs md ...'
func MarkdownCommand() *cli.Command {
	return &cli.Command{
		Name:            "markdown",
		Usage:           "creates Markdown documentation for the GitHub Action",
		HideHelpCommand: true,
		Action:          createMarkdownDocs,
		Aliases:         []string{"md"},
		Flags:           []cli.Flag{createFileFlag, formatFlag},
	}
}

// Called when markdown command is executed
func createMarkdownDocs(c *cli.Context) error {
	// Get flag values
	shouldCreateFile := c.Bool(createFileFlag.Name)
	format := c.String(formatFlag.Name)

	// Validate format flag
	if !util.Contains(allowedFormats, format) {
		return fmt.Errorf("format '%s' is not supported, use one of '%v'", format, allowedFormats)
	}

	// Parse GitHub Action metadata file
	meta, err := parseActionMetadata()
	if err != nil {
		return err
	}

	var content string
	var contentErr error
	if format == allowedFormats[0] {
		// table
		content, contentErr = createInputOutputTable(meta)
	} else {
		// list
		content, contentErr = createInputOutputList(meta)
	}
	if contentErr != nil {
		return contentErr
	}

	if shouldCreateFile {
		// create markdown file
		return createDocsFile(meta, content)
	}

	// Print generated content to stdout
	fmt.Println(util.Normalize(content))

	return nil
}
