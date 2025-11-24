package handlers

import (
	"net/http"
	"ecommerce/product"
	"ecommerce/util"
)

func HandleProduct(write http.ResponseWriter, read *http.Request) {
	util.SendData(write, product.ProductList, 200)
}