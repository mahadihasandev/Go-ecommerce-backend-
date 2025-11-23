package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleAbout(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintln(write, "this is About")
}

type Product struct {
	Id          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var ProductList []Product

func handleProduct(write http.ResponseWriter, read *http.Request) {
	sendData(write, ProductList, 200)
}

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

func sendData(write http.ResponseWriter, data interface{}, statusCode int) {
	write.WriteHeader(statusCode)
	encoder := json.NewEncoder(write)
	encoder.Encode(data)
}

func corsMiddleware(next http.Handler) http.Handler {
	handleCors := func(write http.ResponseWriter, read *http.Request) {
		write.Header().Set("Access-Control-Allow-Origin", "*")
		write.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		write.Header().Set("Content-Type", "Application/json")

		next.ServeHTTP(write, read)
	}
	handler := http.HandlerFunc(handleCors)
	return handler
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /about", http.HandlerFunc(handleAbout))

	mux.Handle("GET /product", corsMiddleware(http.HandlerFunc(handleProduct)))

	mux.Handle("POST /products", corsMiddleware(http.HandlerFunc(setProduct)))

	httpGlobalRouter := GlobalRoute(mux)

	fmt.Println("server is running in port:8000")

	err := http.ListenAndServe(":8000", httpGlobalRouter)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	prod1 := Product{
		Id:          1,
		Title:       "big mac",
		Description: "good quality food for alien",
		Price:       39.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-vqKcsZM2d2WxBPHhMD6IG5xDC-K-TDD59Q&s",
	}

	prod2 := Product{
		Id:          2,
		Title:       "Mango",
		Description: "good quality food for cyborg",
		Price:       29.99,
		ImgUrl:      "https://www.biovie.fr/img/cms/histoire-origine-mangue.png",
	}

	prod3 := Product{
		Id:          3,
		Title:       "Banana",
		Description: "good quality food for Meta-human",
		Price:       19.99,
		ImgUrl:      "https://nutritionsource.hsph.harvard.edu/wp-content/uploads/2018/08/bananas-1354785_1920.jpg",
	}

	ProductList = append(ProductList, prod1)
	ProductList = append(ProductList, prod2)
	ProductList = append(ProductList, prod3)
}

func GlobalRoute(mux *http.ServeMux) http.Handler {
	GlobalRoutes := func(write http.ResponseWriter, read *http.Request) {
		if read.Method == "OPTIONS" {
			write.Header().Set("Access-Control-Allow-Origin", "*")
			write.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			write.Header().Set("Content-Type", "Application/json")
			write.WriteHeader(200)
		} else {
			mux.ServeHTTP(write, read)
		}
	}
	return http.HandlerFunc(GlobalRoutes)

}
