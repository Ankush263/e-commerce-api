package utils

import (
	"log"

	common "github.com/ankush263/e-commerce-api/common"
	Models "github.com/ankush263/e-commerce-api/models"
)


func CreateUserInDB(user Models.UsersModel) (Models.ResponseUsersModel, error) {
    db := common.SetupDB()
    var response Models.ResponseUsersModel
    err := db.QueryRow(`INSERT INTO users(username, email, password, phone) 
                        VALUES($1, $2, $3, $4) 
                        RETURNING id, created_at, updated_at, username, email, password, phone`,
        user.UserName, user.Email, user.Password, user.Phone).Scan(
            &response.ID, 
            &response.CreatedAt, 
            &response.UpdatedAt, 
            &response.UserName, 
            &response.Email, 
            &response.Password, 
            &response.Phone)

    common.CheckError(err)
    return response, nil
}

func GetUsersFromDB() ([]Models.ResponseUsersModel, error) {
    db := common.SetupDB()
    
    rows, err := db.Query(`SELECT id, created_at, updated_at, username, email, password, phone FROM users`)
    if err != nil {
        log.Printf("Error querying users: %v", err)
        return nil, err
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
        common.CheckError(err)
        users = append(users, user)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        log.Printf("Error iterating rows: %v", err)
        return nil, err
    }

    return users, nil
}