package utils

import (
	"context"
	"fmt"
	"github.com/Roh-Bot/Backend/models/Registration"
)

func GetEmailFromDB() {
	pool := PostgresConnectionPool()
	var Register = &Registration.Register
	fmt.Println(Register.Email)
	rows := pool.QueryRow(context.Background(), "SELECT user_id from users WHERE email=$1", Register.Email)
	errScan := rows.Scan(&Register.User_Id)
	if errScan != nil {
		fmt.Println("Scanning Error ", errScan)
	}
	fmt.Println(Register.User_Id)
}
