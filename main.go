package main

import (
	"fmt"
	"net/http"
	"ecommerce/handlers"
	"ecommerce/globalRouter"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /about", http.HandlerFunc(handlers.HandleAbout))
	mux.Handle("GET /product", http.HandlerFunc(handlers.HandleProduct))
	mux.Handle("POST /products", http.HandlerFunc(handlers.SetProduct))
	httpGlobalRouter := globalRouter.GlobalRoute(mux)
	fmt.Println("server is running in port:8000")
	err := http.ListenAndServe(":8000", httpGlobalRouter)
	if err != nil {
		fmt.Println(err)
	}
}