package logger

<<<<<<< HEAD
import (
	"go.uber.org/zap"
)
=======
import "go.uber.org/zap"
>>>>>>> master

var Log *zap.Logger

func Initializing(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	cfg := zap.NewProductionConfig()

	cfg.Level = lvl
<<<<<<< HEAD
=======

>>>>>>> master
	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	Log = zl
	return nil
}
<<<<<<< HEAD

func Mserror(msg string, err error, fields []zap.Field) {
	fields = append(fields, zap.Error(err))
	Log.Error(msg, fields...)
}

func Msinfo(msg string, err error, fields []zap.Field) {
	fields = append(fields, zap.Error(err))
	Log.Info(msg, fields...)
}
=======
>>>>>>> master
