package tpl

import (
	"bytes"
	"fmt"
	"text/template"
)

// CreateTemplate creates a Golang text template with the given name and the submitted template string
func CreateTemplate(name, tplString string) (*template.Template, error) {
	tt, err := template.New(name).Funcs(templateFuncs()).Parse(tplString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %s", err)
	}

	return tt, nil
}

// RenderTemplates fills the submitted template with the submitted content and returns the template content as string
func RenderTemplate(tt *template.Template, content interface{}) (string, error) {
	var buffer bytes.Buffer
	err := tt.Execute(&buffer, content)
	if err != nil {
		return "", fmt.Errorf("error while rendering template: %s", err)
	}

	return buffer.String(), nil
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		// Simulates the ternary operator, checks if cond is true and returns values accordingly
		"ternary": func(cond bool, trueValue, falseValue string) string {
			if cond {
				return trueValue
			}
			return falseValue
		},
		// Checks if submittedVal is non-emtpy and returns values accordingly
		"isSet": func(submittedVal, trueValue, falseValue string) string {
			if submittedVal != "" {
				return trueValue + fmt.Sprintf("`%s`", submittedVal)
			}
			return falseValue
		},
		// Surrounds submittedVal with backticks (displayed as code in markdown)
		"code": func(submittedVal string) string {
			return fmt.Sprintf("`%s`", submittedVal)
		},
	}
}
