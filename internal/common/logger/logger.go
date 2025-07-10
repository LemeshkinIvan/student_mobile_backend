package custom_logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

var Logg = InitLogger()

func InitLogger() *slog.Logger {
	settings := &lumberjack.Logger{
		Filename:   "../logs/app.log",
		MaxSize:    5,    // megabytes
		MaxBackups: 12,   //count of files
		MaxAge:     30,   //days
		Compress:   true, //gzip
	}

	multiWriter := io.MultiWriter(os.Stdout, settings)
	logger := slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		AddSource: true,
	}))

	gin.DefaultWriter = multiWriter
	gin.DefaultErrorWriter = multiWriter

	return logger
}
