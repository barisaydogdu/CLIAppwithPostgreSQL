package user

import "time"

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Number    int
	Balance   int
	Method    string
	Typ       string
	CreatedAt time.Time
}
