package models

// City init
type City struct {
	ID   uint   `json:"id" form:"id"`
	City string `json:"city" form:"city"`
}

// TableName name
func (City) TableName() string {
	return "cities"
}
