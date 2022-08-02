package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "ws agent",
}

func init() {
	RootCmd.AddCommand(ServerCmd)
}
