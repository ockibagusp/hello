package models

// City init
type City struct {
	ID   uint
	City string
}

// TableName name
func (City) TableName() string {
	return "cities"
}
