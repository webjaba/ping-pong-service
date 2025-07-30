// @title Мой API
// @version 1.0
// @host localhost:9000
// @BasePath /
package main

import (
	"net/http"
	_ "ping-pong-service/docs"
	"ping-pong-service/internal/app"

	httpSwagger "github.com/swaggo/http-swagger"
)

// swagger: http://localhost:9000/docs/
func main() {
	s := app.New()

	http.HandleFunc("/docs/", httpSwagger.WrapHandler)
	http.HandleFunc("/about/", s.About)
	http.HandleFunc("/ping/", s.Ping)

	http.ListenAndServe(":9000", nil)
}
