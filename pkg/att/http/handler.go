package http

/*
import (
	"github.com/ronaudinho/bijipipe/pkg/att/service"
)
*/

// AttHttpHandler is http handler for att
type AttHttpHandler struct {
	attMessage string
	// attService *service.AttService
}

// NewAttHttpHandler is new instance of AttHttpHandler
func NewAttHttpHandler(attMessage string) *AttHttpHandler {
	// func NewAttHttpHandler(attService *service.AttService) *AttHttpHandler {
	return &AttHttpHandler{attMessage: attMessage}
	// return &AttHttpHandler{attService: attService}
}
