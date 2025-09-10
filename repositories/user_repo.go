package repositories
import (
    "context"
    "cashmate-api/config"
    "cashmate-api/models"
    "log"
)

func CreateUser(user models.User) error{
    conn, err := config.ConnectDB()
    if err != nil {
        return err
    }

    defer conn.Close(context.Background())
    query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
    _, err = conn.Exec(context.Background(), query, user.Username, user.Email, user.Password)
    if err != nil {
        log.Println("Error inserting user:", err)
        return err
    }
    return nil
}