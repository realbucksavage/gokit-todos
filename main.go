package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"gokit-todos/database"
	"gokit-todos/todos"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP listen address")
	flag.Parse()

	db := database.InitDb()
	ctx := context.WithValue(context.Background(), database.ContextKey, db)

	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		handler := mux.NewRouter()
		handler.Use(commonMiddleware)

		todos.AddRoutes(ctx, handler)
		errChan <- http.ListenAndServe(*addr, handler)
	}()

	log.Fatalln(<-errChan)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
