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

func addDeveloper(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "",
	}

	var developer Developer
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&developer)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if developer.FirstName == "" {
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if developer.LastName == "" {
			httpError.Message = "Last Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if developer.ID != 0 {
			httpError.Message = "ID should be empty for new Developer"
			returnErrorResponse(response, request, httpError)
		} else {
			db.Create(&developer)

			//db.Set("gorm:save_associations", true).Create(&developer)
			//db.Set("gorm:save_associations", true).Save(&developer)

			json.NewEncoder(response).Encode(developer)
		}
	}

}

func updateDeveloper(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "",
	}

	var developer Developer
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&developer)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if developer.FirstName == "" {
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if developer.LastName == "" {
			httpError.Message = "Last Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if developer.ID == 0 {
			httpError.Message = "ID should be supplied for updating Developer"
			returnErrorResponse(response, request, httpError)
		} else {
			//TODO check if ID is present in DB

			db.Save(&developer)
			json.NewEncoder(response).Encode(developer)
		}
	}

}

func deleteDeveloper(response http.ResponseWriter, request *http.Request) {

	var httpError = ErrorResponse{
		Code: http.StatusPreconditionFailed, Message: "",
	}
	developerID := mux.Vars(request)["id"]
	if developerID == "" {
		httpError.Message = "Developer id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {

		db.Where("id = ?", developerID).Delete(&Developer{})
		response.WriteHeader(http.StatusNoContent)

	}


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

func addSkill(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "",
	}

	var skill Skill
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&skill)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if skill.Name == "" {
			httpError.Message = "Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if skill.Description == "" {
			httpError.Message = "Description can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if skill.ID != 0 {
			httpError.Message = "ID should be empty for new Skill"
			returnErrorResponse(response, request, httpError)
		} else {
			db.Create(&skill)

			//db.Set("gorm:save_associations", true).Create(&developer)
			//db.Set("gorm:save_associations", true).Save(&developer)

			json.NewEncoder(response).Encode(skill)
		}
	}

}

func updateSkill(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "",
	}

	var skill Skill
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&skill)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if skill.Name == "" {
			httpError.Message = "Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if skill.Description == "" {
			httpError.Message = "Description can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if skill.ID == 0 {
			httpError.Message = "ID should be supplied for updating Skill"
			returnErrorResponse(response, request, httpError)
		} else {
			//TODO check if ID is present in DB

			db.Save(&skill)
			json.NewEncoder(response).Encode(skill)
		}
	}

}

func deleteSkill(response http.ResponseWriter, request *http.Request) {

	var httpError = ErrorResponse{
		Code: http.StatusPreconditionFailed, Message: "",
	}
	skillID := mux.Vars(request)["id"]
	if skillID == "" {
		httpError.Message = "Skill id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {

		db.Where("id = ?", skillID).Delete(&Skill{})
		response.WriteHeader(http.StatusNoContent)

	}


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


