package globalRouter

import "net/http"

func GlobalRoute(mux *http.ServeMux) http.Handler {
	AllRequest := func(write http.ResponseWriter, read *http.Request) {
		write.Header().Set("Access-Control-Allow-Origin", "*")
		write.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		write.Header().Set("Content-Type", "Application/json")

		if read.Method == "OPTIONS" {
			write.WriteHeader(200)
			return
		}
		mux.ServeHTTP(write, read)
	}
	return http.HandlerFunc(AllRequest)

}