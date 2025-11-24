package main

import (
	"fmt"
	"net/http"
	"ecommerce/handlers"
	"ecommerce/globalRouter"
	"ecommerce/product"
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

func init() {
	prod1 := product.Product{
		Id:          1,
		Title:       "big mac",
		Description: "good quality food for alien",
		Price:       39.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-vqKcsZM2d2WxBPHhMD6IG5xDC-K-TDD59Q&s",
	}

	prod2 := product.Product{
		Id:          2,
		Title:       "Mango",
		Description: "good quality food for cyborg",
		Price:       29.99,
		ImgUrl:      "https://www.biovie.fr/img/cms/histoire-origine-mangue.png",
	}

	prod3 := product.Product{
		Id:          3,
		Title:       "Banana",
		Description: "good quality food for Meta-human",
		Price:       19.99,
		ImgUrl:      "https://nutritionsource.hsph.harvard.edu/wp-content/uploads/2018/08/bananas-1354785_1920.jpg",
	}
	product.ProductList = append(product.ProductList, prod1)
	product.ProductList = append(product.ProductList, prod2)
	product.ProductList = append(product.ProductList, prod3)
}