package main

import (
	"context"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"

	"github.com/tzrikka/ovid/internal/temporal"
	"github.com/tzrikka/ovid/pkg/slack"
)

// TODO: DELETE THIS BINARY!

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		fmt.Printf("client dial error: %v\n", err)
		os.Exit(1)
	}
	defer c.Close()

	opts := client.StartWorkflowOptions{
		TaskQueue: temporal.DefaultTaskQueue,
	}

	input := slack.ChatPostMessageRequest{
		Channel:      "U3TF86ZH7",
		MarkdownText: "Hello!",
	}

	run, err := c.ExecuteWorkflow(context.Background(), opts, "MiniWorkflow", input)
	if err != nil {
		fmt.Printf("execute workflow error: %v\n", err)
		return // os.Exit(1)
	}
	fmt.Printf("workflow ID %q, Run ID %q\n", run.GetID(), run.GetRunID())

	var result string
	if err := run.Get(context.Background(), &result); err != nil {
		fmt.Printf("unable to get workflow result: %v\n", err)
	}

	fmt.Println(result)
}
