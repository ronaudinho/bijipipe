package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetInjured return not injury status response
func (Cen *CenHttpHandler) GetInjured(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	// TODO create helper to return json
	w.Header().Set("Content-Type", "application/json")
	// TODO investigate why order of status header matters in firefox response page
	// w.WriteHeader(http.StatusOK)

	// dummy object
	type DummyJsonResponse struct {
		Message string `json:"message"`
	}

	// turn object to json
	dummy := DummyJsonResponse{Message: "pelir"}
	d, err := json.Marshal(dummy)
	fmt.Println(err)

	w.Write(d)
}
