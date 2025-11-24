package handlers

import (
	"net/http"
	"ecommerce/database"
	"ecommerce/util"
)

func HandleProduct(write http.ResponseWriter, read *http.Request) {
	util.SendData(write, database.ProductList, 200)
}