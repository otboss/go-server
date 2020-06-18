package middleware

import (
	"log"
	"model"
	"net/http"
	"strings"
)

// RequestLimiter : Limits the number and size of requests made by clients
func RequestLimiter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		IPAddress := req.RemoteAddr[:strings.LastIndex(req.RemoteAddr, ":")]
		requestCount, requestCountFetchError := model.GetRequestCount(IPAddress)
		if requestCountFetchError != nil {
			log.Println(requestCountFetchError)
			panic("internal server error")
		}
		if requestCount.Count >= 1200 {
			// Client has made too many requests within the hour
			panic("too many requests")
		}
		requestCountIncrementError := model.IncrementRequestCount(IPAddress)
		if requestCountIncrementError != nil {
			log.Println(requestCountIncrementError)
			panic("internal server error")
		}
		// Limit the request body size to 500 kb
		req.Body = http.MaxBytesReader(w, req.Body, 500000)
		next(w, req)
	}
}
