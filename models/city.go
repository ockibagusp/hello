package models

// City init
type City struct {
	ID   int
	City string
}

// TableName name
func (City) TableName() string {
	return "cities"
}
