package worker

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)


type Logger struct{}


func NewLogger() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return &Logger{}
}

func (logger *Logger) Print(v ...interface{}) {
	log.Info().Msg(fmt.Sprint(v...))
}

func (logger *Logger) Printf(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func (logger *Logger) Println(v ...interface{}) {
	log.Info().Msg(fmt.Sprint(v...))
}


// Custom log levels
func (logger *Logger) Debug(format string, v ...interface{}) { log.Debug().Msgf(format, v...) }
func (logger *Logger) Info(format string, v ...interface{})  { log.Info().Msgf(format, v...) }
func (logger *Logger) Warn(format string, v ...interface{})  { log.Warn().Msgf(format, v...) }
func (logger *Logger) Error(format string, v ...interface{}) { log.Error().Msgf(format, v...) }
func (logger *Logger) Fatal(format string, v ...interface{}) { log.Fatal().Msgf(format, v...) }
