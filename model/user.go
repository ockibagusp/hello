package model

import "time"

// User init
type User struct {
	ID        int
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	City      uint
	Photo     string
}

var (
	// Users init
	Users = map[int]*User{}
	// Seq init
	Seq = 1
)
