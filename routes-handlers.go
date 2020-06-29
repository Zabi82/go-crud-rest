package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)



func getDeveloper(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusPreconditionFailed, Message: "",
	}
	developerID := mux.Vars(request)["id"]
	if developerID == "" {
		httpError.Message = "Developer id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
	

		var developer Developer
		db.Preload("Skills").First(&developer, developerID)
		json.NewEncoder(response).Encode(developer)
	}
}

func getDevelopers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var developers []Developer
	db.Preload("Skills").Find(&developers)
	json.NewEncoder(response).Encode(developers)


}

func getSkill(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusPreconditionFailed, Message: "",
	}
	skillID := mux.Vars(request)["id"]
	if skillID == "" {
		httpError.Message = "Developer id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
	

		var skill Skill
		db.Preload("Developers").First(&skill, skillID)
		json.NewEncoder(response).Encode(skill)
	}
}

func getSkills(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var skills []Skill
	db.Preload("Developers").Find(&skills)
	json.NewEncoder(response).Encode(skills)


}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}


