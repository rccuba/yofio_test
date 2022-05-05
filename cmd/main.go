package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"test_robert_yofio/internal/app"
	"test_robert_yofio/internal/config"
	"test_robert_yofio/internal/middleware"
	"test_robert_yofio/internal/static"
)

var httpHandler = app.App{}

func main() {
	if err, isConfigurable := config.ConfigEnv(); !isConfigurable {
		fmt.Printf(""+static.MsgResponseStartError+", %s", err)
		values := []interface{}{static.KeyType, static.ERROR, static.KeyMessage, static.MsgResponseStartError + ", " + err.Error()}
		middleware.LoggingOperation(httpHandler.Logg, values...)
	} else {
		fmt.Println(static.MsgResponseStartProcess)
		addr := ":" + viper.GetString(static.APP_PORT)
		_ = httpHandler.Initialize()
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		_ = httpHandler.Run(addr)
	}
}
