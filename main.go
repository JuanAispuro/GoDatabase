package main

import (
	"log"

	"github.com/JuanAispuro/GoDatabase/pkg/invoice"
	"github.com/JuanAispuro/GoDatabase/pkg/invoiceheader"
	"github.com/JuanAispuro/GoDatabase/pkg/invoiceitem"
	"github.com/JuanAispuro/GoDatabase/storage"
)

func main() {
	//Crear la instancia de postgres que maneja el producto
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvocieHeader(storage.Pool())
	storageItems := storage.NewPsqlInvocieItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)
	//Inserción de la base de datos
	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Alejandro",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 3},
			&invoiceitem.Model{ProductID: 1},
		},
	}
	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
	//Si todo sale bien en invoiceheaders y invoice items debe salir.
}
