package update

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/urfave/cli/v2"
)

// Repository name
const repositoryName = "rescDev/actions-docs"

// UpdateCommand descibes the CLI command for updating the tool
// Usage: 'actions-docs update'
func UpdateCommand() *cli.Command {
	return &cli.Command{
		Name:            "update",
		Usage:           "perform a self update of the tool to its latest version",
		HideHelpCommand: true,
		Action:          doSelfUpdate,
	}
}

// Called when update command is executed
func doSelfUpdate(c *cli.Context) error {
	// Package can only parse semvers without "v" prefix (so e.g. 1.0.0 instead of v1.0.0)
	parsedVersion := semver.MustParse(c.App.Version[1:])

	updater, err := selfupdate.NewUpdater(selfupdate.Config{
		// Check only for release asstets starting with 'actions-docs_'
		Filters: []string{fmt.Sprintf("%s_", c.App.Name)},
	})
	if err != nil {
		return fmt.Errorf("could not configure self updater: %s", err)
	}

	latest, err := updater.UpdateSelf(parsedVersion, repositoryName)
	if err != nil {
		return fmt.Errorf("binary update failed: %s", err)
	}

	if latest.Version.Equals(parsedVersion) {
		// latest version is the same as current version; means current binary is up to date
		fmt.Printf("Current binary is the latest version: %s", c.App.Version)
	} else {
		fmt.Printf("Successfully updated to version %s", latest.Version)
		fmt.Printf("Release notes:\n %s", latest.ReleaseNotes)
	}

	return nil
}
