package main

import (
	"github.com/julienschmidt/httprouter"
)

// AttHttpRoute setup routes for this service
func AttHttpRoute() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/injured", attHandler.GetInjured)

	return
}
