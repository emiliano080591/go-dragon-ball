package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emiliano080591/go-dragon-ball/bd"
	"github.com/emiliano080591/go-dragon-ball/models"
	"github.com/gorilla/mux"
)

var resFailure models.ResponseFailure

/*CREATE OPERATIONS*/
func CrearPersonaje(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var personaje models.Personaje
	err := json.NewDecoder(r.Body).Decode(&personaje)

	if err != nil {
		resFailure := models.ResponseFailure{}
		resFailure.Status = false
		resFailure.Message = "ERROR: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	bd.DB.Create(&personaje)
	json.NewEncoder(w).Encode(personaje)
}

/*READ OPERATIONS*/
func ObtenerPersonajes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var personajes []models.Personaje
	bd.DB.Find(&personajes)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(personajes)
}

func ObtenerPersonaje(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var personaje models.Personaje
	params := mux.Vars(r)

	if len(params["id"]) < 1 {
		resFailure.Status = false
		resFailure.Message = "Debe de enviarse el parametro id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	_, err := strconv.Atoi(params["id"])
	if err != nil {
		resFailure.Status = false
		resFailure.Message = "Debe enviar el parámetro ID con un valor mayor a 0"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	er := bd.DB.First(&personaje, params["id"]).Error
	if er != nil {
		resFailure.Status = false
		resFailure.Message = "Registro no encontrado"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(personaje)
}

/*UPDATE OPERATIONS*/
func ActualizarPersonaje(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var personaje models.Personaje
	params := mux.Vars(r)

	if len(params["id"]) < 1 {
		resFailure.Status = false
		resFailure.Message = "Debe de enviarse el parametro id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	_, err := strconv.Atoi(params["id"])
	if err != nil {
		resFailure.Status = false
		resFailure.Message = "Debe enviar el parámetro ID con un valor mayor a 0"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	bd.DB.First(&personaje, params["id"])
	err = json.NewDecoder(r.Body).Decode(&personaje)

	if err != nil {
		resFailure := models.ResponseFailure{}
		resFailure.Status = false
		resFailure.Message = "ERROR: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}
	bd.DB.Save(&personaje)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(personaje)
}

/*DELETE OPERATIONS*/
func BorrarPersonaje(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var personaje models.Personaje
	if len(params["id"]) < 1 {
		resFailure.Status = false
		resFailure.Message = "Debe de enviarse el parametro id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	_, err := strconv.Atoi(params["id"])
	if err != nil {
		resFailure.Status = false
		resFailure.Message = "Debe enviar el parámetro ID con un valor mayor a 0"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resFailure)
		return
	}

	bd.DB.Delete(&personaje, params["id"])
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("The User is Deleted Successfully!")
}
