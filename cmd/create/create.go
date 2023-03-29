package create

import (
	"context"

	"github.com/spf13/cobra"
)

type cmdFlags struct {
	targetToken string
	HOST_NAME   string
}

func NewCmdCreate(ctx context.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create <target organization> [flags]",
		Short: "Create organization level webhooks",
		Long:  "Create organization level webhooks",
		Args:  cobra.ExactArgs(1),
	}

	return cmd
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
	cmd.PersistentFlags().StringVarP(&cmdFlags.targetToken, "target-token", "s", "", "GitHub personal access token for accessing target organization")

	return &cmd
}
