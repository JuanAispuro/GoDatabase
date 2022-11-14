package main

import (
	"github.com/JuanAispuro/GoDatabase/storage"
)

func main() {
	//Crear la instancia de postgres que maneja el producto
	storage.NewMySQLDB()

}
