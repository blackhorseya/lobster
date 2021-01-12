package main

import (
	"flag"
	"strings"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/sirupsen/logrus"
)

const timeFormat = "2006-01-02 15:04:05"

var path = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()
}

func initLogger(cfg *config.Config) {
	if level, err := logrus.ParseLevel(cfg.Log.Level); err != nil {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warn(err)
	} else {
		logrus.SetLevel(level)
	}

	switch strings.ToLower(cfg.Log.Format) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeFormat,
			DisableQuote:    true,
		})
	}
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

	initLogger(injector.C)

	err = injector.Engine.Run(injector.C.HTTP.GetAddress())
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Panicf("run engine of app is panic")
	}
}
