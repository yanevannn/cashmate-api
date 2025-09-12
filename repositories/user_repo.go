package repositories
import (
    "context"
    "cashmate-api/config"
    "cashmate-api/models"
    "log"
)

func CreateUser(user *models.User) error{
    conn, err := config.ConnectDB()
    if err != nil {
        return err
    }

    defer conn.Close(context.Background())
    query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
    // _, err = conn.Exec(context.Background(), query, user.Username, user.Email, user.Password)
    err = conn.QueryRow(context.Background(), query, user.Username, user.Email, user.Password).Scan(&user.ID)
    if err != nil {
        log.Println("Error inserting user:", err)
        return err
    }
    return nil
}

func GetuserByID(id int) (*models.User, error) {
    conn, err := config.ConnectDB()
    if err != nil {
        return nil, err
    }
    defer conn.Close(context.Background())

    // query := `SELECT * FROM users WHERE id = $1`
    query := `SELECT id, username, email, password FROM users WHERE id = $1` // Specify columns explicitly makes it clearer and avoids issues if table schema changes
    row := conn.QueryRow(context.Background(), query, id)
    var user models.User
    err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

    if err != nil {
        log.Println("Error fetching user by ID:", err)
        return nil, err
    }
    return &user, nil
}