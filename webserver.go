package webserver

import (
	"code.google.com/p/tcgl/applog"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Environment struct {
	Listen   string
	LogLevel int
}

type WebServer struct {
	handler http.Handler
}

func New(handler http.Handler) *WebServer {
	return &WebServer{handler}
}

func (srv *WebServer) Run(env Environment) error {
	applog.Infof("Starting application")
	http.Handle("/", srv.handler)

	applog.Infof("Listening to incoming connection on %s", env.Listen)

	if err := http.ListenAndServe(env.Listen, nil); err != nil {
		applog.Criticalf(err.Error())
		return err
	}

	return nil
}

func RunCLI(srv WebServer) {
	var env Environment
	env.Listen = "0.0.0.0:9000"
	env.LogLevel = applog.LevelDebug

	cli := flag.NewFlagSet("WebServer", flag.ExitOnError)
	cli.StringVar(&env.Listen, "l", env.Listen, "The address to listen for incoming connections ([host]:port)")

	cli.Usage = func() {
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		fmt.Printf("Available options:\n")
		cli.PrintDefaults()
	}

	if err := cli.Parse(os.Args[1:]); err == flag.ErrHelp {
		cli.Usage()
		os.Exit(1)
	}

	applog.SetLevel(env.LogLevel)

	if err := srv.Run(env); err != nil {
		panic(err)
	}
}
