package temporal

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/tzrikka/ovid/pkg/slack"
)

// Start initializes logging and the Temporal worker.
func Start(_ context.Context, cmd *cli.Command) error {
	logger := initLog(cmd.Bool("dev"))

	c, err := client.Dial(client.Options{
		HostPort:  cmd.String("temporal-host-port"),
		Namespace: cmd.String("temporal-namespace"),
		Logger:    logAdapter{zerolog: logger},
	})
	if err != nil {
		return fmt.Errorf("client dial error: %w", err)
	}
	defer c.Close()

	w := worker.New(c, cmd.String("temporal-task-queue"), worker.Options{})

	slack.Register(cmd, w)

	return w.Run(worker.InterruptCh())
}
