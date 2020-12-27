package markdown

// ActionMeta describes the metadata syntax of a GitHub Action
// see https://docs.github.com/en/free-pro-team@latest/actions/creating-actions/metadata-syntax-for-github-actions
type ActionMeta struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Inputs      map[string]Input  `yaml:"inputs"`
	Outputs     map[string]Output `yaml:"outputs"`
}

// Input describes a single input variable of the action
type Input struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

// Output describes a single output variable of the action
type Output struct {
	Description string `yaml:"description"`
}
