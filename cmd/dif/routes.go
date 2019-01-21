package main

import (
	"github.com/julienschmidt/httprouter"
)

// DifHttpRoute setup routes for this service
func DifHttpRoute() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/injured", difHandler.GetInjured)

	return
}
