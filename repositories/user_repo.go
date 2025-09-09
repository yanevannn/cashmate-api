package repositories

import "cashmate-api/models"

var Users []models.User

func CreateUser(user models.User) {
    user.ID = len(Users) + 1
    Users = append(Users, user)
}

func GetAllUsers() []models.User {
    return Users
}
