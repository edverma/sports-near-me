package logger

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func InitLogger(logFileName, logPrefix string, ginWriter io.Writer) *log.Logger {
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logFile, err = os.Create(logFileName)
		if err != nil {
			panic(err)
		}
	}
	multiWriter := io.MultiWriter(logFile) // add os.StdErr to log to StdErr
	if ginWriter != nil {
		gin.DefaultWriter = multiWriter
		gin.DisableConsoleColor()
	}

	logPrefix = "\n" + logPrefix
	return log.New(multiWriter, logPrefix, log.LstdFlags|log.Lshortfile|log.LUTC)
}
