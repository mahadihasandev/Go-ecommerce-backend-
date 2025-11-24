package main

import "net/http"

func handleProduct(write http.ResponseWriter, read *http.Request) {
	sendData(write, ProductList, 200)
}