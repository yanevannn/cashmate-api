package repositories

import (
	"cashmate-api/config"
	"cashmate-api/models"
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

func GetUserByEmail(email string) (*models.User, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	query := `SELECT id, username, email, password, role FROM users WHERE email = $1` // Specify columns explicitly makes it clearer and avoids issues if table schema changes
	row := conn.QueryRow(context.Background(), query, email)
	var user models.User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("No user found with email:", email)
			return nil, nil // No user found with this email
		}
		log.Println("Error fetching user by email:", err)
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.RegisterUser, OTP string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}

	// Start a transaction
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background()) // Rollback in case 1 of error

	// 1 Transaction: Insert user
	queryUser := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	var userID int
	err = tx.QueryRow(context.Background(), queryUser, user.Username, user.Email, user.Password).Scan(&userID)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	// 2 Transaction: Insert OTP
	queryOTP := `
                    INSERT INTO user_verifications (user_id, code, expires_at) 
                    VALUES ($1, $2, NOW() + INTERVAL '15 minutes')
                `
	_, err = tx.Exec(context.Background(), queryOTP, userID, OTP)
	if err != nil {
		log.Println("Error inserting OTP:", err)
		return err
	}

	// Commit the transaction
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("Error committing transaction:", err)
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

func GetAllUsers() ([]models.User, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	query := `SELECT id, username, email, password FROM users`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func DeleteUser(id int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := `DELETE FROM users WHERE id = $1`
	_, err = conn.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}
