package main

import (
	"log"

	"github.com/JuanAispuro/GoDatabase/pkg/product"
	"github.com/JuanAispuro/GoDatabase/storage"
)

func main() {
	storage.NewPostgresDB()
	//Crear la instancia de postgres que maneja el producto
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
