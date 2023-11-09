package router

import (
	"mogoIn/controller"

	"github.com/gorilla/mux"
)

func Mains() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc(  "/api/getAllShows", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/updateShow/{id}", controller.UpdateMovie).Methods("POST")
	r.HandleFunc("/api/createShow", controller.CreateMovie).Methods("POST")


	return r
} 