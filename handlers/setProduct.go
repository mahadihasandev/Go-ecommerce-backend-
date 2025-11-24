package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce/product"
	"ecommerce/util"
	
)

func SetProduct(write http.ResponseWriter, read *http.Request) {
	var setProductList product.Product
	decoder := json.NewDecoder(read.Body)
	err := decoder.Decode(&setProductList)

	if err != nil {
		http.Error(write, "Send a json object", 400)
		return
	}
	setProductList.Id = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, setProductList)
	util.SendData(write, setProductList, 201)
}