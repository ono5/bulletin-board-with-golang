// routes.go
package main

import (
	"bulletin-board/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	// staticファイルの配信設定
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
