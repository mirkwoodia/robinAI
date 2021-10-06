package main

import (
	//"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mirkwoodia/RobinAI/bot"
	"github.com/mirkwoodia/RobinAI/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
