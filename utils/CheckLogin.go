package utils

import (
	"context"
	"fmt"
)

func CheckLoginCredentials(email string, mobile int, password string) (bool, string, int) {
	pool := PostgresConnectionPool()

	//pgCheckIfUserExists := `SELECT email FROM users WHERE email='Dhebug@God.com'`
	var emailScan int
	row := pool.QueryRow(context.Background(), `select user_id from users WHERE (email=$1 OR mobile=$2) AND password=$3`, email, mobile, password)
	fmt.Println("Before Scan: ", emailScan)
	errScan := row.Scan(&emailScan)
	if errScan != nil {
		fmt.Println(errScan)
	}
	fmt.Println("After Scan: ", emailScan)
	if emailScan != 0 {
		return true, "Logged in", emailScan
	}
	return false, "Credentials do not match", emailScan

}
