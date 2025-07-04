package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"

	"github.com/tzrikka/ovid/internal/temporal"
	"github.com/tzrikka/ovid/internal/thrippy"
	"github.com/tzrikka/ovid/pkg/slack"
	"github.com/tzrikka/xdg"
)

const (
	ConfigDirName  = "ovid"
	ConfigFileName = "config.toml"
)

func main() {
	bi, _ := debug.ReadBuildInfo()

	cmd := &cli.Command{
		Name:    "ovid",
		Usage:   "Temporal worker using Thrippy and Omdient",
		Version: bi.Main.Version,
		Flags:   flags(),
		Action:  temporal.Start,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func flags() []cli.Flag {
	fs := []cli.Flag{
		&cli.BoolFlag{
			Name:  "dev",
			Usage: "simple setup, but unsafe for production",
		},
	}

	// Core settings.
	path := configFile()
	fs = append(fs, temporal.Flags(path)...)
	fs = append(fs, thrippy.Flags(path)...)

	// Supported Thrippy Links IDs.
	fs = append(fs, slack.LinkIDFlag(path))

	return fs
}

// configFile returns the path to the app's configuration file.
// It also creates an empty file if it doesn't already exist.
func configFile() altsrc.StringSourcer {
	path, err := xdg.CreateFile(xdg.ConfigHome, ConfigDirName, ConfigFileName)
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
	}
	return altsrc.StringSourcer(path)
}
