package middleware

import "net/http"

// GetRequest :
func GetRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			panic("only supports get requests")
		}
		next(w, req)
	}
}
