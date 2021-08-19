package controllers

import (
	"fmt"
	"net/http"
)

func Health() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "{ site: up 🚀 , deployed_version: 🚦1.0.0🚦}")
	})
}
