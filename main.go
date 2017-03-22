package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	"github.com/octoblu/server-status-code/statusserver"
	"github.com/urfave/cli"
	De "github.com/visionmedia/go-debug"
)

var debug = De.Debug("server-status-code:main")

func main() {
	app := cli.NewApp()
	app.Name = "server-status-code"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "delay, d",
			EnvVar: "DELAY",
			Usage:  "Milliseconds to wait before sending the reponse",
			Value:  0,
		},
		cli.IntFlag{
			Name:   "port, p",
			EnvVar: "PORT",
			Usage:  "Port to listen on",
			Value:  80,
		},
		cli.IntFlag{
			Name:   "status-code, s",
			EnvVar: "STATUS_CODE",
			Usage:  "Status code to return to the user",
			Value:  200,
		},
	}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	delay, port, statusCode := getOpts(context)

	server := statusserver.New(delay, port, statusCode)
	err := server.Run()
	log.Fatalln(err.Error)
}

func getOpts(context *cli.Context) (int, int, int) {
	delay := context.Int("delay")
	port := context.Int("port")
	statusCode := context.Int("status-code")

	if statusCode == 0 {
		cli.ShowAppHelp(context)

		if statusCode == 0 {
			color.Red("  Missing required flag --status-code, -s or STATUS_CODE")
		}
		os.Exit(1)
	}

	return delay, port, statusCode
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
