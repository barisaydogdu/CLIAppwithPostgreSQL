package repository

import (
	"time"
)

type User struct {
	Id         int
	First_name string
	Last_name  string
	Number     int
	Balance    int
	Created_at time.Time
}
