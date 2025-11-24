package main

import (
	"encoding/json"
	"net/http"
)

func setProduct(write http.ResponseWriter, read *http.Request) {
	var setProductList Product
	decoder := json.NewDecoder(read.Body)
	err := decoder.Decode(&setProductList)

	if err != nil {
		http.Error(write, "Send a json object", 400)
		return
	}
	setProductList.Id = len(ProductList) + 1
	ProductList = append(ProductList, setProductList)
	sendData(write, setProductList, 201)
}