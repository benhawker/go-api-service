package main

import (
	"net/http"
	"os"
	"os/signal"

	"context"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	log "github.com/sirupsen/logrus"
	routes "github.com/benhawker/go-api-service/internal/routes"
)

const port = ":1111"

var wg sync.WaitGroup

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)


	routes.Register(router)
	router.HandleFunc("/health", health)

	server := &http.Server{
		Addr: port, 
		Handler: loggedRouter,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	}
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Sever starting. Listening on port ", port)
		if err := server.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	<-stop
	log.Info("Shutting down...")
	server.Shutdown(context.Background())
	wg.Wait()
	log.Info("🏁 Server stopped")
}



func health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}
