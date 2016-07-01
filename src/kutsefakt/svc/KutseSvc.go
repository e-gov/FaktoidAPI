package main

import (
	"flag"
	"fmt"
	logging "github.com/op/go-logging"
	"net/http"
	"os"
	"faktoid"
	"kutsefakt"
)

var log = logging.MustGetLogger("KutseSvc")

func main() {
	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	var slog = flag.Bool("syslog", false, "If present, logs are sent to syslog")

	flag.Parse()

	setupLogging(slog)
	// Instantiate a new KutseFakt and set up the handlers with it
	faktoid.InitFakt(new(kutsefakt.KutseFakt))
	router := faktoid.NewRouter()
	log.Infof("Starting a server on localhost:%d", *port)
	log.Critical(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}


func setupLogging(slog *bool)  {
	var b logging.Backend

	format := logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	if *slog  {
	b, _ = logging.NewSyslogBackend("Login")
	}else{
	b = logging.NewLogBackend(os.Stdout, "", 0)
	}

	bFormatter := logging.NewBackendFormatter(b, format)
	logging.SetBackend(bFormatter)
}