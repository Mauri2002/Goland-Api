package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Definición de la estructura Proyecto
type Proyecto struct {
	NumeroControl  string `json:"numero_control"`
	NombreAlumno   string `json:"nombre_alumno"`
	Carrera        string `json:"carrera"`
	NombreProyecto string `json:"nombre_proyecto"`
}

var db *sql.DB

// Función principal
func main() {
	var err error
	// Conectar a la base de datos MySQL
	dsn := "root:290200Carlos$@tcp(localhost:3306)/residencias" // Reemplaza con tus credenciales y base de datos
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable()

	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/proyectos", GetAllProyectos).Methods("GET")
	router.HandleFunc("/proyectos/{numero_control}", GetProyectoByID).Methods("GET")
	router.HandleFunc("/proyectos", CreateProyecto).Methods("POST")
	router.HandleFunc("/proyectos/{numero_control}", DeleteProyecto).Methods("DELETE")

	fmt.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
