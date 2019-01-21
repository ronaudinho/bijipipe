package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetInjured return not injury status response
func (Dif *DifHttpHandler) GetInjured(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	// TODO create helper to return json
	w.Header().Set("Content-Type", "application/json")

	// dummy object
	type DummyJsonResponse struct {
		message string `json:"message"`
	}

	// turn object to json
	dummy := DummyJsonResponse{message: "pelir"}
	d, err := json.Marshal(dummy)
	fmt.Println(err)

	w.Write(d)
	return
}
