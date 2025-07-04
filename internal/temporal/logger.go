package temporal

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// initLog initializes the logger for the Temporal worker,
// based on whether it's running in development mode or not.
func initLog(devMode bool) zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	if !devMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
		return log.Logger
	}

	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05.000",
	}).With().Caller().Logger()

	log.Warn().Msg("********** DEV MODE - UNSAFE IN PRODUCTION! **********")
	return log.Logger
}

// logAdapter implements the https://pkg.go.dev/go.temporal.io/sdk/log#Logger interface.
type logAdapter struct {
	zerolog zerolog.Logger
}

func (a logAdapter) Debug(msg string, keyvals ...any) {
	logMessageWithAttributes(a.zerolog.Debug(), msg, keyvals)
}

func (a logAdapter) Info(msg string, keyvals ...any) {
	logMessageWithAttributes(a.zerolog.Info(), msg, keyvals)
}

func (a logAdapter) Warn(msg string, keyvals ...any) {
	logMessageWithAttributes(a.zerolog.Warn(), msg, keyvals)
}

func (a logAdapter) Error(msg string, keyvals ...any) {
	logMessageWithAttributes(a.zerolog.Error().Stack(), msg, keyvals)
}

func logMessageWithAttributes(e *zerolog.Event, msg string, keyvals ...any) {
	for i, kv := range keyvals {
		as, ok := kv.([]any)
		if !ok {
			e = e.Any(fmt.Sprintf("attr_%d", i), kv)
			continue
		}
		for len(as) > 0 {
			e = e.Any(as[0].(string), as[1])
			as = as[2:]
		}
	}

	if msg == "" {
		e.Send()
	} else {
		e.Msg(msg)
	}
}
