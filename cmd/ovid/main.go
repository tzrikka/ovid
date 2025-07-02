package main

import (
	"context"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli/v3"

	"github.com/tzrikka/ovid/pkg/slack"
	"github.com/tzrikka/ovid/pkg/temporal"
	"github.com/tzrikka/ovid/pkg/thrippy"
	"github.com/tzrikka/xdg"
)

const (
	ConfigDirName  = "ovid"
	ConfigFileName = "config.toml"
)

func main() {
	buildInfo, _ := debug.ReadBuildInfo()
	configFilePath := configFile()

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "dev",
			Usage: "simple setup, but unsafe for production",
		},
	}

	flags = append(flags, temporal.Flags(configFilePath)...)
	flags = append(flags, thrippy.Flags(configFilePath)...)

	flags = append(flags, slack.LinkIDFlag(configFilePath))

	cmd := &cli.Command{
		Name:    "ovid",
		Usage:   "Temporal worker using Thrippy and Omdient",
		Version: buildInfo.Main.Version,
		Flags:   flags,
		Action:  temporal.Start,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Caller().Send()
	}
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
