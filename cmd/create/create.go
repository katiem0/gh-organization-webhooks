package create

import (
	"errors"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
	"github.com/cli/go-gh/pkg/auth"
	"github.com/katiem0/gh-organization-webhooks/internal/data"
	"github.com/katiem0/gh-organization-webhooks/internal/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type createCmdFlags struct {
	sourceToken string
	sourceOrg   string
	token       string
	hostname    string
	fileName    string
	debug       bool
}

func NewCreateCmd() *cobra.Command {
	createCmdFlags := createCmdFlags{}
	var authToken string

	createCmd := &cobra.Command{
		Use:   "create <target organization> [flags]",
		Short: "Create organization level webhooks",
		Long:  "Create organization level webhooks",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(createCmd *cobra.Command, args []string) error {
			if len(createCmdFlags.fileName) == 0 && len(createCmdFlags.sourceOrg) == 0 {
				return errors.New("A file or source organization must be specified where webhooks will be created from.")
			} else if len(createCmdFlags.sourceOrg) > 0 && len(createCmdFlags.sourceToken) == 0 {
				return errors.New("A Personal Access Token must be specified to access webhooks from the Source Organization.")
			}
			return nil
		},
		RunE: func(createCmd *cobra.Command, args []string) error {
			var err error
			var restClient api.RESTClient

			// Reinitialize logging if debugging was enabled
			if createCmdFlags.debug {
				logger, _ := log.NewLogger(createCmdFlags.debug)
				defer logger.Sync() // nolint:errcheck
				zap.ReplaceGlobals(logger)
			}

			if createCmdFlags.token != "" {
				authToken = createCmdFlags.token
			} else {
				t, _ := auth.TokenForHost(createCmdFlags.hostname)
				authToken = t
			}

			restClient, err = gh.RESTClient(&api.ClientOptions{
				Headers: map[string]string{
					"Accept": "application/vnd.github+json",
				},
				Host:      createCmdFlags.hostname,
				AuthToken: authToken,
			})

			if err != nil {
				zap.S().Errorf("Error arose retrieving rest client")
				return err
			}

			owner := args[0]

			return runCreateCmd(owner, &createCmdFlags, data.NewAPIGetter(restClient))
		},
	}
	// Configure flags for command
	createCmd.PersistentFlags().StringVarP(&createCmdFlags.token, "token", "t", "", `GitHub personal access token for organization to write to (default "gh auth token")`)
	createCmd.PersistentFlags().StringVarP(&createCmdFlags.sourceToken, "source-token", "s", "", `GitHub personal access token for Source Organization (Required for --source-organization)`)
	createCmd.PersistentFlags().StringVarP(&createCmdFlags.sourceOrg, "source-organization", "o", "", `Name of the Source Organization to copy webhooks from (Requires --source-token)`)
	createCmd.PersistentFlags().StringVarP(&createCmdFlags.hostname, "hostname", "", "github.com", "GitHub Enterprise Server hostname")
	createCmd.Flags().StringVarP(&createCmdFlags.fileName, "from-file", "f", "", "Name of file to create webhooks from")
	createCmd.PersistentFlags().BoolVarP(&createCmdFlags.debug, "debug", "d", false, "To debug logging")

	return createCmd
}

func runCreateCmd(owner string, createCmdFlags *createCmdFlags, g *data.APIGetter) error {
	return nil
}
