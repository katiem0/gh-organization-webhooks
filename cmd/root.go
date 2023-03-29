package cmd

import (
	"github.com/spf13/cobra"

	createCmd "github.com/katiem0/gh-organization-webhooks/cmd/create"
	listCmd "github.com/katiem0/gh-organization-webhooks/cmd/list"
)

func NewCmd() *cobra.Command {

	cmd := cobra.Command{
		Use:   "organization-webhooks <command> [flags]",
		Short: "List and create organization webhooks.",
		Long:  "List and create organization level webhooks.",
	}

	cmd.AddCommand(listCmd.NewCmdList(ctx))
	cmd.AddCommand(createCmd.NewCmdCreate(ctx))
}
