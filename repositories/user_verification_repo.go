package repositories

import (
	"cashmate-api/config"
	"context"
	"fmt"
)

func VerificationOtpIsValid(userID int, code string) (bool, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return false, err
	}
	defer conn.Close(context.Background())

	query := `
				SELECT COUNT(1) 
				FROM user_verifications 
				WHERE user_id = $1 
				AND code = $2 
				AND expires_at > NOW() 
				AND is_used = FALSE
			`

	var count int
	err = conn.QueryRow(context.Background(), query, userID, code).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func ValidateOTP(userID int, code string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := `UPDATE user_verifications SET is_used = TRUE WHERE user_id = $1 AND code = $2`
	updateOTP, err := conn.Exec(context.Background(), query, userID, code)
	if err != nil {
		return err
	}

	if updateOTP.RowsAffected() == 0 {
		return fmt.Errorf("otp not found or already used")
	}

	return nil
}

func StoreNewOTP(userID int, code string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := `
	INSERT INTO user_verifications (user_id, code, expires_at, is_used)
	VALUES ($1, $2, NOW() + INTERVAL '15 minutes', FALSE)
	ON CONFLICT (user_id) DO UPDATE SET 
		code = EXCLUDED.code, 
		expires_at = EXCLUDED.expires_at, 
		is_used = FALSE
	`
	_, err = conn.Exec(context.Background(), query, userID, code)
	if err != nil {
		return err
	}

	return nil
}
