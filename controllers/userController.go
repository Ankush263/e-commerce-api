package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	common "github.com/ankush263/e-commerce-api/common"
	"github.com/ankush263/e-commerce-api/models"
	model "github.com/ankush263/e-commerce-api/models"
	utils "github.com/ankush263/e-commerce-api/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type MyCustomClaims struct {
	Userid int `json:"userid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// This function will generate the JWT token with userid and email.
func generateJWT(userid int, email string) (string, error) {
    claims := MyCustomClaims{
        userid,
        email,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(jwtSecret)
}

// This function will set the JWT token to the header.
func setJWTToken(w http.ResponseWriter, user model.ResponseUsersModel, reqUser model.UsersModel) {
	token, jwtErr := generateJWT(user.ID, *reqUser.Email)

    common.CheckError("Generate JWT Error: ", jwtErr)

	// Set the cookies to the header
    http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token,
		HttpOnly: true,
		Secure: true,
		Path: "/",
		Expires:  time.Now().Add(24 * time.Hour),
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var reqUser model.UsersModel
	_ = json.NewDecoder(r.Body).Decode(&reqUser)

	user, err := utils.CreateUserInDB(reqUser)

	common.CheckError("Create User Error: ", err)

	setJWTToken(w, user, reqUser)

	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqUser model.UsersModel
	_ = json.NewDecoder(r.Body).Decode(&reqUser)

	response := utils.GetUserByEmail(*reqUser.Email)

	if response.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Invalid Email or Password",
			Data: nil,
		})
		return 
	}

	if err := bcrypt.CompareHashAndPassword([]byte(response.Data.Password), []byte(*reqUser.Password)); err != nil {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Invalid Email or Password",
			Data: nil,
		})
		return 
	}

	setJWTToken(w, *response.Data, reqUser)

	json.NewEncoder(w).Encode(model.ResponseModel{
		Status: 1,
		Message: "success",
		Data: &response.Data,
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	
	allUsers := utils.GetUsersFromDB()

	if allUsers.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Users not found",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: allUsers.Data,
		})
	}

}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	params := mux.Vars(r)
	id := params["id"]

	userdata := utils.GetSingleUserInDB(id)

	if userdata.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "User with that id doesn't exists",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: userdata.Data,
		})
	}
}

func UpdateSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	var reqUser model.UsersModel
	_ = json.NewDecoder(r.Body).Decode(&reqUser)

	params := mux.Vars(r)
	id := params["id"]

	updatedData := utils.UpdateSingleUserById(reqUser, id)

	if updatedData.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "User with that id doesn't exists",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: updatedData.Data,
		})
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	
	params := mux.Vars(r)
	id := params["id"]

	deleteUser := utils.DeleteUserById(id)

	if deleteUser == "success" {
		json.NewEncoder(w).Encode("Deleted")
	} else {
		json.NewEncoder(w).Encode("Failed to delete the user")
	}
}
