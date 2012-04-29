package webserver

import (
	"code.google.com/p/tcgl/applog"
	"flag"
	"fmt"
	"net/http"
	"os"
  "time"
  "path"
)

type Environment struct {
	Listen       string
  PublicFolder string
	LogLevel     int
}

type WebServer struct {
	appHandler http.Handler
}

func New(handler http.Handler) *WebServer {
	return &WebServer{handler}
}

func (srv *WebServer) Run(env *Environment) error {
	applog.Infof("Starting application")

  disp := newDispatcher(srv, env)
	http.Handle("/", disp)

	applog.Infof("Listening to incoming connection on %s", env.Listen)

	if err := http.ListenAndServe(env.Listen, nil); err != nil {
		applog.Criticalf(err.Error())
		return err
	}

	return nil
}

func (srv *WebServer) RunCLI() {
  env := &Environment{
	  Listen: "0.0.0.0:9000",
	  LogLevel: applog.LevelDebug,
    PublicFolder: "public",
  }

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

type dispatcher struct {
  env *Environment
  appHandler, publicHandler http.Handler
}

func newDispatcher(srv *WebServer, env *Environment) *dispatcher {
  return &dispatcher{env, srv.appHandler, http.FileServer(http.Dir(env.PublicFolder))}
}

func (disp *dispatcher) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now()

	applog.Infof("Started %s %s", req.Method, req.URL)

	localPath := path.Join(disp.env.PublicFolder, req.URL.Path)

	applog.Debugf("Checking for existance of %s", localPath)

	if info, err := os.Stat(localPath); err == nil {
		if info.IsDir() {
			indexPath := path.Join(localPath, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				applog.Debugf("Serving static file %s", indexPath)
				disp.publicHandler.ServeHTTP(w, req)
			}

		} else {
			applog.Debugf("Serving static file %s", localPath)
			disp.publicHandler.ServeHTTP(w, req)
		}
	} else {
		applog.Debugf("Dispatching %s", req.URL)
    disp.appHandler.ServeHTTP(w, req)
	}

	duration := time.Now().Sub(startTime)
	applog.Infof("Finished processing %s %s (%s)", req.Method, req.URL, duration.String())
}
