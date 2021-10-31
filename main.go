package main

import (
	"fiber-101/utils"
	"flag"

	"github.com/sirupsen/logrus"
)

var port = flag.String("port", ":5050", "Port to listen on")

func main() {
	app := utils.SetupApp()

	logrus.Fatal(app.Listen(*port))
}
