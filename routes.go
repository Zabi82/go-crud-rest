package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func addApproutes(route *mux.Router) {

	setStaticFolder(route)

	
	route.HandleFunc("/developer/{id}", getDeveloper).Methods("GET")

	route.HandleFunc("/developer", getDevelopers).Methods("GET")


	route.HandleFunc("/skill/{id}",getSkill).Methods("GET")

	route.HandleFunc("/skill",getSkills).Methods("GET")

	

}
