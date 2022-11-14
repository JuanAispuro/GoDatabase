package storage

import (
	"database/sql"
	"fmt"

	"github.com/JuanAispuro/GoDatabase/pkg/invoice"
	"github.com/JuanAispuro/GoDatabase/pkg/invoiceheader"
	"github.com/JuanAispuro/GoDatabase/pkg/invoiceitem"
)

// MySQLInvoice estructura para trabajar con postgres de la factura.
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// Constructor de la factura.
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Crea la interfaz invoice model
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("Header: %w", err)
	}
	fmt.Printf("Factura creada con id: %d \n", m.Header.ID)
	//Acceder al detalle
	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("Items: %w", err)
	}
	fmt.Printf("Items creados: %d \n", len(m.Items))

	return tx.Commit() //Confirmar y que se registren el encabezado y el detalle
}
