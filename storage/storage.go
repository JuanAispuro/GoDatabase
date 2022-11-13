package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

// crearemos el singleton
// Variables Globales
var (
	db   *sql.DB //Tipo puntero
	once sync.Once
)

/*
"postgres","postgres://

	postgres:Destructor11@localhost:5432/Godb?sslmode=disable"
*/
func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:Destructor11@localhost:5432/Godb?sslmode=disable")
		if err != nil {
			log.Fatalf("No pudimos abrir la base de datos: %v", err)
		}
		// defer db.Close()
		if err := db.Ping(); err != nil {
			log.Fatalf("No pudimos hacer ping a la base de datos: %v", err)
		}
		fmt.Println("Conectado a la base de datos")
	}) //Lo que este aqui adentro solo se ejecutara una vez
}

// Pool retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}
