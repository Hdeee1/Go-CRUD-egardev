package main

import (
	"go-crud-egardev/config"
	"go-crud-egardev/controllers/categorycontroller"
	"go-crud-egardev/controllers/homecontroller"
	"go-crud-egardev/controllers/productcontroller"
	"log"
	"net/http"
)

func main(){
	config.ConnectDB()

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)


	
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}