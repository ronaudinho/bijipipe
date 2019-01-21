package http

/*
import (
	"github.com/ronaudinho/bijipipe/pkg/cen/service"
)
*/

// CenHttpHandler is http handler for cen
type CenHttpHandler struct {
	cenMessage string
	// cenService *service.CenService
}

// NewCenHttpHandler is new instance of CenHttpHandler
func NewCenHttpHandler(cenMessage string) *CenHttpHandler {
	// func NewCenHttpHandler(cenService *service.CenService) *CenHttpHandler {
	return &CenHttpHandler{cenMessage: cenMessage}
	// return &CenHttpHandler{cenService: cenService}
}
