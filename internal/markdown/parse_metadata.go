package markdown

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	// Possible names of GitHub Action metadata files
	metaFileNames = []string{"action.yaml", "action.yml"}
)

// Parses the GitHub Action metadata file and returns its unmarshalled content
// Looks for a file with the possible file names and fails if none can be found
// If both file names are be found, the content will be overwritten
func parseActionMetadata() (*ActionMeta, error) {
	var meta *ActionMeta

	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		return meta, fmt.Errorf("unable to determine current directory, can not read metadata file: %s", err)
	}

	var input []byte
	var readErr error

	// Check if a file called 'action.yaml' is provided
	input, readErr = ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, metaFileNames[0]))
	if readErr != nil {
		// Check 'action.yml'
		input, readErr = ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, metaFileNames[1]))
		if readErr != nil {
			return meta, fmt.Errorf("unable to parse action metadata file (action.yaml/yml): %s", err)
		}
	}

	// Unmarshal file content
	err = yaml.Unmarshal(input, &meta)
	if err != nil {
		return meta, err
	}

	return meta, nil
}
