package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time
		start := time.Now()
	
		fmt.Printf("request entered | path: %v | method: %v\n", r.URL.Path, r.Method)
	
		next.ServeHTTP(w, r)
	
		fmt.Printf("request completed | path: %v | method: %v | time: %v\n",
		r.URL.Path, r.Method, time.Since(start))

	})
	


}

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Frame-Options", "Deny")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	

	})
	


}

