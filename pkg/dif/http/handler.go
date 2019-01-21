package http

/*
import (
	"github.com/ronaudinho/bijipipe/pkg/dif/service"
)
*/

// DifHttpHandler is http handler for dif
type DifHttpHandler struct {
	difMessage string
	// difService *service.DifService
}

// NewDifHttpHandler is new instance of DifHttpHandler
func NewDifHttpHandler(difMessage string) *DifHttpHandler {
	// func NewDifHttpHandler(difService *service.DifService) *DifHttpHandler {
	return &DifHttpHandler{difMessage: difMessage}
	// return &DifHttpHandler{difService: difService}
}
