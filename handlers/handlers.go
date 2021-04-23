package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/emiliano080591/dragon/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	log.Println("Servidor corriendo en el puerto 8080")
	router := mux.NewRouter()

	router.HandleFunc("/getPersonajes", routers.ObtenerPersonajes).Methods("GET")
	router.HandleFunc("/personaje/{id}", routers.ObtenerPersonaje).Methods("GET")
	router.HandleFunc("/personaje", routers.CrearPersonaje).Methods("POST")
	router.HandleFunc("/personaje/{id}", routers.ActualizarPersonaje).Methods("PUT")
	router.HandleFunc("/personaje/{id}", routers.BorrarPersonaje).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
