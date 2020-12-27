package markdown

import (
	"github.com/rescDev/actions-docs/pkg/tpl"
)

const (
	listInputTpl = `
	## Inputs
	{{ if not .Inputs }}
		No inputs defined.
	{{- else }}
		{{- range $name, $input := .Inputs }}
			- {{ code $name }}: {{ $input.Description }}. {{ ternary $input.Required "Is required" "Is not required" }}. {{ isSet $input.Default "Defaults to " "" }}
		{{- end }}
	{{- end }}
	`

	listOutputTpl = `
	## Outputs
	{{ if not .Outputs }}
		No outputs defined.
	{{- else }}
		{{- range $name, $output := .Outputs }}
			- {{ code $name }}: {{ $output.Description }}
		{{- end }}
	{{- end -}}
	`
)

// Creates a markdown list for the inputs and outputs and returns their contents as string
func createInputOutputList(meta *ActionMeta) (string, error) {
	// Create and render input list
	ttInput, err := tpl.CreateTemplate("inputs", listInputTpl)
	if err != nil {
		return "", err
	}

	inputContent, err := tpl.RenderTemplate(ttInput, meta)
	if err != nil {
		return "", err
	}

	// Create and render output list
	ttOutput, err := tpl.CreateTemplate("outputs", listOutputTpl)
	if err != nil {
		return "", err
	}

	outputContent, err := tpl.RenderTemplate(ttOutput, meta)
	if err != nil {
		return "", err
	}

	return inputContent + outputContent, nil
}
