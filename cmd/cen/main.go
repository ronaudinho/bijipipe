package main

import (
	"fmt"
	"net/http"
	"os"

	cenHttp "github.com/ronaudinho/bijipipe/pkg/cen/http"

	"github.com/gorilla/handlers"
)

var (
	cenHandler *cenHttp.CenHttpHandler
)

func main() {
	cenHandler = cenHttp.NewCenHttpHandler("pelirlo")
	router := handlers.CombinedLoggingHandler(os.Stdout, CenHttpRoute())
	router = handlers.RecoveryHandler()(router)

	fmt.Println("biji at 8083")
	fmt.Println(http.ListenAndServe(":8083", router))
}
