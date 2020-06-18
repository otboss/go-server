package controller

import (
	"fmt"
	"net/http"
)

// Root : Root Route
func Root(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `hello world`)
}
