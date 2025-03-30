package utils

import (
	"fmt"

	common "github.com/ankush263/e-commerce-api/common"
	Models "github.com/ankush263/e-commerce-api/models"
	"golang.org/x/crypto/bcrypt"
)


type CreateUserResponse struct {
    Status int `json:"status"`
    Error error `json:"error"`
    Data Models.ResponseUsersModel `json:"data"`
}

type UsersResponse struct {
    Status int `json:"status"`
    Data *[]Models.ResponseUsersModel `json:"data"`
}


func CreateUserInDB(user Models.UsersModel) CreateUserResponse {
    db := common.SetupDB()
    var response Models.ResponseUsersModel

    // Hash the password
    hashedPassword, hasherror := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
    common.CheckError("Password Hashing Error: ", hasherror)

    err := db.QueryRow(`INSERT INTO users(username, email, password, role, phone) 
                        VALUES($1, $2, $3, $4, $5) 
                        RETURNING id, created_at, updated_at, username, email, password, role, phone`,
        user.UserName, user.Email, string(hashedPassword), user.Role, user.Phone).Scan(
            &response.ID, 
            &response.CreatedAt, 
            &response.UpdatedAt, 
            &response.UserName, 
            &response.Email, 
            &response.Password, 
            &response.Role, 
            &response.Phone)

    common.CheckError("Error: ", err)
    if err != nil {
        return CreateUserResponse{
            Status: 0,
            Error: err,
        }
    }

    return CreateUserResponse{
        Status: 1,
        Data: response,
    } 
}

func GetUsersFromDB() UsersResponse {
    db := common.SetupDB()
    
    rows, err := db.Query(`SELECT id, created_at, updated_at, username, email, password, phone FROM users`)
    if err != nil {
        fmt.Println("Error: ", err)
        return UsersResponse{
            Status: 0,
            Data: nil,
        }
    }
    defer rows.Close()

    var users []Models.ResponseUsersModel

    for rows.Next() {
        var user Models.ResponseUsersModel
        err := rows.Scan(
            &user.ID,
            &user.CreatedAt,
            &user.UpdatedAt,
            &user.UserName,
            &user.Email,
            &user.Password,
            &user.Phone,
        )
        common.CheckError("Error: ", err)
        users = append(users, user)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        fmt.Println("Error: ", err)
        return UsersResponse{
            Status: 0,
            Data: nil,
        }
    }

    return UsersResponse{
        Status: 1,
        Data: &users,
    }
}

func GetSingleUserInDB(id string) Models.ResponseDBModel {
    db := common.SetupDB()

    var response Models.ResponseUsersModel
    err := db.QueryRow(`SELECT id, created_at, updated_at, username, email, password, role, phone FROM users WHERE id = $1`, id).Scan(
        &response.ID,
        &response.CreatedAt,
        &response.UpdatedAt,
        &response.UserName,
        &response.Email,
        &response.Password,
        &response.Role,
        &response.Phone,
    )

    if err != nil {
        return Models.ResponseDBModel{
            Status: 0,
            Data: nil,
        }
    }

    return Models.ResponseDBModel{
        Status: 1,
        Data: &response,
    }
}

func GetUserByEmail(email string) Models.ResponseDBModel {
    db := common.SetupDB()

    var response Models.ResponseUsersModel
    err := db.QueryRow(`SELECT id, created_at, updated_at, username, email, password, phone FROM users WHERE email = $1`, email).Scan(
        &response.ID,
        &response.CreatedAt,
        &response.UpdatedAt,
        &response.UserName,
        &response.Email,
        &response.Password,
        &response.Phone,
    )

    if err != nil {
        return Models.ResponseDBModel{
            Status: 0,
            Data: nil,
        }
    }

    return Models.ResponseDBModel{
        Status: 1,
        Data: &response,
    }
}

func UpdateSingleUserById(user Models.UsersModel, id string) Models.ResponseDBModel {
    db := common.SetupDB()

    var response Models.ResponseUsersModel
    err := db.QueryRow(`
        UPDATE users 
        SET 
            username = COALESCE($1, username),
            email = COALESCE($2, email),
            phone = COALESCE($3, phone),
            updated_at = NOW()
        WHERE id = $4
        RETURNING id, created_at, updated_at, username, email, phone, password;
    `,user.UserName, user.Email, user.Phone, id).Scan(
        &response.ID,
        &response.CreatedAt,
        &response.UpdatedAt,
        &response.UserName,
        &response.Email,
        &response.Phone,
        &response.Password,
    )

    if err != nil {
        fmt.Println("Error: ", err)
        return Models.ResponseDBModel{
            Status: 0,
            Data: nil,
        }
    }

    return Models.ResponseDBModel{
        Status: 1,
        Data: &response,
    }
}

func DeleteUserById(id string) string {
    db := common.SetupDB()

    result, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
    if err != nil {
        fmt.Println("Error:", err)
        return "error"
    }

    rowsAffected, err := result.RowsAffected()

    if err != nil {
        fmt.Println("Error: ", err)
        return "error"
    }

    if rowsAffected == 0 {
        return "not found" 
    }

    return "success"
}
