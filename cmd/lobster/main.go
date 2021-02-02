package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	cmd, err := CreateCommand()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Panicf("create a command is panic")
	}
	if cmd == nil {
		return
	}

	if err = cmd.Execute(); err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error(err)
		os.Exit(1)
	}
}
