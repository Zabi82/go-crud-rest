package main

import (

	
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server will start at http://localhost:8585/")

	connectDatabase()

	route := mux.NewRouter()

	addApproutes(route)

	log.Fatal(http.ListenAndServe(":8585", route))
}


 