package routers

import (
	"mongodb/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET",)
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST",)
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatch).Methods("PUT",)
	router.HandleFunc("/api/movie/{id}", controllers.DeleteAMovie).Methods("DELETE",)
	router.HandleFunc("/api/deleteallmovie", controllers.DeleteAllMovie).Methods("DELETE",)

	return router
}