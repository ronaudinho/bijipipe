package main

import (
	"fmt"
	"os"

	"github.com/ronaudinho/bijipipe/pkg/cen/http"

	"github.com/gorilla/handlers"
)

var (
	cenHandler *http.CenHttpHandler
)

func main() {
	cenHandler = http.NewCenHttpHandler("pelirlo")
	router := handlers.CombinedLoggingHandler(os.Stdout, CenHttpRoute())
	router = handlers.RecoveryHandler()(router)

	fmt.Println("biji at 8083")
}
