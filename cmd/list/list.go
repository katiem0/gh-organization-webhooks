package list

import (
	"context"

	"github.com/spf13/cobra"
)

type cmdFlags struct {
	sourceToken string
	HOST_NAME   string
}

func NewCmdList(ctx context.Context) *cobra.Command {
	cmdFlags := cmdFlags{}

	cmd := &cobra.Command{
		Use:   "list <source organization> [flags]",
		Short: "List organization level webhooks",
		Long:  "List organization level webhooks",
		Args:  cobra.ExactArgs(1),
	}
	// Configure flags for command
	cmd.PersistentFlags().StringVarP(&cmdFlags.sourceToken, "source-token", "s", "", "GitHub personal access token for reading source organization")

	return &cmd
}
