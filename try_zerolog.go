package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = ""

	log.Print("hello world log")

	log.Info().Msg("hello world again")

	log.Warn().Msg("warning")

	log.Error().Msg("error")

	//	zerolog.SetGlobalLevel(zerolog.Disabled)

	//	log.Fatal().Msg("fatal")

	//	log.Panic().Msg("im panic")

	log.Debug().Str("game", "league of heros").
		Float64("score", 99.33).
		Msg("s8 will begin in 2018 se")

	log.Log().
		Str("foo", "bar l").
		Msg("")
	// different outputs
	f, err := os.Create("hahaha")
	if err != nil {
		log.Fatal().Err(err).Msg("shit happens")
	}
	defer f.Close()

	logger := zerolog.New(f).With().Timestamp().Logger()
	logger.Info().Str("foo", "bar").Msg("hello file")

	// Sub-loggers
	sublog := log.With().
		Str("component", "sub content").
		Logger()

	sublog.Info().Msg("HELLo world")

	// Sub dictionary
	log.Info().
		Str("foo key", "bar value").
		Dict("foo dic key", zerolog.Dict().
			Str("key bbb", "value aaa").
			Int("int key", 124),
		).Msg("hello dictionary")

	// sample log
	// sampled := log.Sample(&zerolog.BasicSampler{N: 10})
	// for i := 0; i < 20; i++ {
	// 	sampled.Info().Msg("willbe logged every 10 messges")
	// }
	// sampled := log.Sample(
	// 	zerolog.LevelSampler{
	// 		DebugSampler: &zerolog.BurstSampler{
	// 			Burst:       5,
	// 			Period:      time.Second,
	// 			NextSampler: &zerolog.BasicSampler{N: 500},
	// 		},
	// 	})
	// for i := 0; i < 2000; i++ {
	// 	sampled.Info().Msg("will not be logged ")
	// 	sampled.Debug().Msg("will be logged every 500 messges")
	// }

	// custromize automatic field names
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"

	log.Info().Msg("heelo customize field names")

	// Pretty logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Str("key foo", "value bar").Msg("pretty soup")

}
