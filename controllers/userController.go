package controllers

import (
	"encoding/json"
	"net/http"

	common "github.com/ankush263/e-commerce-api/common"
	model "github.com/ankush263/e-commerce-api/models"
	utils "github.com/ankush263/e-commerce-api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var reqUser model.UsersModel
	_ = json.NewDecoder(r.Body).Decode(&reqUser)

	userid, err := utils.CreateUserInDB(reqUser)

	common.CheckError(err)
	json.NewEncoder(w).Encode(userid)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	
	allUsers, err := utils.GetUsersFromDB()
	common.CheckError(err)

	json.NewEncoder(w).Encode(allUsers)
}
