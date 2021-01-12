package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var path = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()
}

func main() {
	injector, err := CreateInjector(*path)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Panicf("create an injector is panic")
	}
	if injector == nil {
		return
	}

	err = injector.Engine.Run(injector.C.HTTP.GetAddress())
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Panicf("run engine of app is panic")
	}
}
