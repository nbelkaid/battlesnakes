package game

import (
	mux "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB
)

// ConfigRoute configure toutes les routes pour gérer la récupération d'image depuis S3.
func ConfigRoute(dB *gorm.DB, r *mux.Router) error {
	DB = dB

	s := r.PathPrefix("/game").Subrouter()

	s.HandleFunc("/", HandleIndex).Methods("GET")
	s.HandleFunc("/start", HandleStart).Methods("POST")
	s.HandleFunc("/move", HandleMove).Methods("POST")
	s.HandleFunc("/end", HandleEnd).Methods("POST")

	return nil
}
