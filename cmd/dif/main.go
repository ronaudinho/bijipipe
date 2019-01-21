package main

import (
	"fmt"
	"net/http"
	"os"

	difHttp "github.com/ronaudinho/bijipipe/pkg/dif/http"

	"github.com/gorilla/handlers"
)

var (
	difHandler *difHttp.DifHttpHandler
)

func main() {
	difHandler = difHttp.NewDifHttpHandler("pelirlo")
	router := handlers.CombinedLoggingHandler(os.Stdout, DifHttpRoute())
	router = handlers.RecoveryHandler()(router)

	fmt.Println("biji at 8082")
	fmt.Println(http.ListenAndServe(":8082", router))
}
