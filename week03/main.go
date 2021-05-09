// Week03, 本次作业主要借鉴了go-kratos的app实现的
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	app := New()

	srv1ServMutex := http.NewServeMux()
	srv1ServMutex.HandleFunc("/", srv1Hello)
	srv1 := &http.Server{
		Addr:           ":8000",
		Handler:        srv1ServMutex,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srv2ServMutex := http.NewServeMux()
	srv2ServMutex.HandleFunc("/", srv2Hello)
	srv2 := &http.Server{
		Addr:           ":9000",
		Handler:        srv2ServMutex,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// srv1.ListenAndServe()
	// srv2.ListenAndServe()
	srvs := []*http.Server{srv1, srv2}
	app.SetServers(srvs...)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func srv1Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this is srv1\n")
}

func srv2Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this is srv2\n")
}

type App struct {
	ctx     context.Context
	cancel  func()
	sigs    []os.Signal
	servers []*http.Server
}

func New() *App {
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:    ctx,
		cancel: cancel,
		sigs:   []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT},
	}
}

func (a *App) Run() error {
	g, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done() // wait for stop signal
			logrus.Info("http server shutting down")
			return srv.Shutdown(context.Background())
		})
		g.Go(func() error {
			return srv.ListenAndServe()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

func (a *App) SetOSSigs(sigs ...os.Signal) {
	a.sigs = sigs
}

func (a *App) SetServers(srvs ...*http.Server) {
	a.servers = srvs
}
