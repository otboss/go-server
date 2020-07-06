package middleware

import (
	"fmt"
	"net/http"
)

// PanicHandler :
func PanicHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			req.Body.Close()
			rec := recover()
			if rec != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Add("Content-Type", "application/json")
				fmt.Fprintf(w, `{"error": "%s"}`, rec)
			}
		}()
		next(w, req)
	}
}
