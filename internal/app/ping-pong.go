package app

import (
	"fmt"
	"net/http"
)

type PingPong interface {
	About(w http.ResponseWriter, r *http.Request)
	Ping(w http.ResponseWriter, r *http.Request)
}

// About Выводит общую информацию о сервисе
// @Success 200 {string} string "sipmle ping-pong service, try to send request to /ping url"
// @Router /about [get]
func (s *Service) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "sipmle ping-pong service, try to send request to /ping url")
}

// Ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (s *Service) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
