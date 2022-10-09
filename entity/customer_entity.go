package entity

type Customer struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
