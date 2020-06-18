package middleware

import "net/http"

// PostRequest :
func PostRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			panic("only supports post requests")
		}
		next(w, req)
	}
}
