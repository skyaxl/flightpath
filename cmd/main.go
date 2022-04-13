package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/skyaxl/flightpath/api"
	"github.com/skyaxl/flightpath/flightpathservice"
)

func main() {
	portStr := os.Getenv("APP_PORT")
	port, _ := strconv.ParseInt(portStr, 10, 64)
	if port == 0 {
		port = 8080
	}
	r := mux.NewRouter()
	service := flightpathservice.New()
	r.Handle("/calculate", api.NewHandler(service))

	errs := make(chan error, 2)
	go func() {
		golog.Infof("[Flight Path Api] Has started with address localhost:%d\n", port)
		errs <- http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	golog.Infof("[Flight Path Api] Has stoped. \n%s\n", <-errs)
}
