package main

import (
	"fmt"
	"os"

	"github.com/ronaudinho/bijipipe/pkg/dif/http"

	"github.com/gorilla/handlers"
)

var (
	difHandler *http.DifHttpHandler
)

func main() {
	difHandler = http.NewDifHttpHandler("pelirlo")
	router := handlers.CombinedLoggingHandler(os.Stdout, DifHttpRoute())
	router = handlers.RecoveryHandler()(router)

	fmt.Println("biji at 8082")
}
