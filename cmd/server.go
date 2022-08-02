package cmd

import (
	"context"
	"wsagent/logic"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		SignalObserve(cancel)
		s := logic.NewWebSocketServer(ctx)
		_ = s.Serve()
	},
}
