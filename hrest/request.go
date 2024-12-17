package hrest

import "net/http"

// Get
func Get(next http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(next, http.MethodGet)
}

// Post
func Post(next http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(next, http.MethodPost)
}

// Put
func Put(next http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(next, http.MethodPut)
}

// Delete
func Delete(next http.HandlerFunc) http.HandlerFunc {
	return allowedMethod(next, http.MethodDelete)
}

func allowedMethod(next http.HandlerFunc, method string) http.HandlerFunc {
	withCors(next)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
		next(w, r)
	}
}

func withCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
