package database

var ProductList []Product
type Product struct {
	Id          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
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