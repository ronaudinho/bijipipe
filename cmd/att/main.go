package main

import (
	"fmt"
	"net/http"
	"os"

	attHttp "github.com/ronaudinho/bijipipe/pkg/att/http"

	"github.com/gorilla/handlers"
)

var (
	attHandler *attHttp.AttHttpHandler
)

func main() {
	attHandler = attHttp.NewAttHttpHandler("pelirlo")
	router := handlers.CombinedLoggingHandler(os.Stdout, AttHttpRoute())
	router = handlers.RecoveryHandler()(router)

	fmt.Println("biji at 8084")
	fmt.Println(http.ListenAndServe(":8084", router))
}
