package main

import (
	"log"

	"github.com/JuanAispuro/GoDatabase/pkg/product"
	"github.com/JuanAispuro/GoDatabase/storage"
)

func main() {
	//Crear la instancia de postgres que maneja el producto
	storage.NewPostgresDB()

	//Create datos de product
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.GetByID: %v", err)
	}
}
