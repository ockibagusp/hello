package moder

import "time"

// User init
type User struct {
	ID        int
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
	City      uint
	Photo     string
}
