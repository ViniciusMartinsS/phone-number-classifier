package api

import "net/http"

func SetHeaders(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().
			Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
