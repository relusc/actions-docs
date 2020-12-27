package markdown

import (
	"github.com/rescDev/actions-docs/pkg/tpl"
)

const (
	tableInputTpl = `
	## Inputs
	{{ if not .Inputs }}
		No inputs defined.
	{{- else }}
		| Name | Description | Required | Default |
		|------|-------------|----------|:-------:|
		{{- range $name, $input := .Inputs }}
			| {{ code $name }} | {{ $input.Description }} | {{ ternary $input.Required "yes" "no" }} | {{ isSet $input.Default "" "-"}} |
		{{- end }}
	{{- end }}
	`

	tableOutputTpl = `
	## Outputs
	{{ if not .Outputs }}
		No outputs defined.
	{{- else }}
		| Name | Description |
		|------|:-----------:|
		{{- range $name, $output := .Outputs }}
			| {{ $name }} | {{ $output.Description }} |
		{{- end }}
	{{- end -}}
	`
)

// Creates a markdown table for the inputs and outputs and returns their contents as string
func createInputOutputTable(meta *ActionMeta) (string, error) {
	// Create and render input table
	ttInput, err := tpl.CreateTemplate("inputs", tableInputTpl)
	if err != nil {
		return "", err
	}

	inputContent, err := tpl.RenderTemplate(ttInput, meta)
	if err != nil {
		return "", err
	}

	// Create and render output table
	ttOutput, err := tpl.CreateTemplate("outputs", tableOutputTpl)
	if err != nil {
		return "", err
	}

	outputContent, err := tpl.RenderTemplate(ttOutput, meta)
	if err != nil {
		return "", err
	}

	return inputContent + outputContent, nil
}
