package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/config"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/externals/datadog"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/logger"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/wire"
)

func main() {

	logger.Setting()

	dd := datadog.NewClient(config.NewDatadogConfig())
	dd.StartTrace()
	defer dd.StopTrace()

	di, c, err := wire.DI()
	if err != nil {
		log.Fatal("DI failed ", err)
		return
	}
	defer c()

	di.Router.Apply(di.Server.Echo)

	go func() {
		err = di.Server.Start()
		if err != nil {
			// logger.Error("Server Start failed", err)
			log.Fatal("Server Start failed", err)
			return
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := di.Server.Echo.Shutdown(ctx); err != nil {
		// logger.Error("Server Shutdown failed", err)
		log.Fatal("Server Shutdown failed", err)
	}
}
