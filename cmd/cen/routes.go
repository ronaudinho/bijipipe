package main

import (
	"github.com/julienschmidt/httprouter"
)

// CenHttpRoute setup routes for this service
func CenHttpRoute() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/injured", cenHandler.GetInjured)

	return
}
