package cmd

import (
	"github.com/spf13/cobra"
	"gogo/backend/server"
)

var RootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		server.Start()
		return nil
	},
}
