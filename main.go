package main

import (
	"fmt"
	"log"

	"github.com/JuanAispuro/GoDatabase/pkg/product"
	"github.com/JuanAispuro/GoDatabase/storage"
)

func main() {
	//driver := storage.MySQL
	driver := storage.MySQL
	//Crear la instancia de postgres que maneja el producto
	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}
	serviceProduct := product.NewService(myStorage)

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println(ms)
}
