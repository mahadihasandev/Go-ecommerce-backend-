package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce/database"
	"ecommerce/util"
	
)

func SetProduct(write http.ResponseWriter, read *http.Request) {
	var setProductList database.Product
	decoder := json.NewDecoder(read.Body)
	err := decoder.Decode(&setProductList)

	if err != nil {
		http.Error(write, "Send a json object", 400)
		return
	}
	setProductList.Id = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, setProductList)
	util.SendData(write, setProductList, 201)
}