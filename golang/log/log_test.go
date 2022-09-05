package main

import (
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"testing"
)

var fakeMessage = "Test logging, but use a somewhat realistic message length."

func BenchmarkZerolog(b *testing.B) {

	logger := zerolog.New(ioutil.Discard).Level(zerolog.InfoLevel)
	//b.ResetTimer()
	//b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next() {
	//		logger.Info().Msg(fakeMessage)
	//	}
	//})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().Msg(fakeMessage)
	}
}

func BenchmarkZap(b *testing.B) {
	logger, _ := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.DebugLevel),
		//OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder, // INFO

			TimeKey:    "time",
			EncodeTime: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
	defer logger.Sync() // flushes buffer, if any
	for i := 0; i < b.N; i++ {
		logger.Info(fakeMessage)
	}

}
