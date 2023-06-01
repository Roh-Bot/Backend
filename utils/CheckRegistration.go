package utils

import (
	"context"
	"fmt"
)

func CheckIfEmailExists(email string) bool {
	pool := PostgresConnectionPool()

	//pgCheckIfUserExists := `SELECT email FROM users WHERE email='Dhebug@God.com'`
	var emailScan string
	row := pool.QueryRow(context.Background(), `SELECT email FROM users where email=$1`, email)

	errScan := row.Scan(&emailScan)
	if errScan != nil {
		fmt.Println(errScan)
	}

	if emailScan != "" {
		return true
	}
	return false

}
