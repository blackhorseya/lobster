package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var path = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title Lobster API
// @version 0.0.1
// @description Lobster API

// @contact.name Sean Cheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space

// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html

// @BasePath /apis
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
