package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/isavinof/pricer/app"
)

func main() {
	host, _ := os.Hostname()
	logrus.Infof("Hostname:%v", host)
	app, err := app.NewApp(os.Environ())
	if err != nil {
		logrus.WithError(err).Fatalf("init app")
		return
	}

	err = app.Run()
	if err != nil {
		logrus.WithError(err).Fatalf("init app")
		return
	}

	return
}
