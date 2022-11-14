package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/JuanAispuro/GoDatabase/pkg/product"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// crearemos el singleton
// Variables Globales
var (
	db   *sql.DB //Tipo puntero
	once sync.Once
)

// Driver de storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// -------------------------------- DAO --------------------------------
// Crea la conexión con la base de datos.
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// -------------------------------- PostgresDB --------------------------------

func newPostgresDB() {
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

// -------------------------------- MySQL --------------------------------
func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:Destructor11@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("No pudimos abrir la base de datos: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("No pudimos hacer ping a la base de datos: %v", err)
		}
		fmt.Println("Conectado a la MySQL")
	}) //Lo que este aqui adentro solo se ejecutara una vez
}

// Pool retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}

// Nullstring
func stringToNull(s string) sql.NullString {
	//Si hay un valor entonces retorna un true.
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// NullTime
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

// DAO factory para product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("Driver no implementado")
	}

}
