package models


import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

)

type User struct {
	gorm.Model
	Name string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf("models.User{ID: %d, CreatedAt: %v, UpdatedAt: %v, DeletedAt: %v, Name: %s, Email: %s}",
		u.ID, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.Name, u.Email)
}

type Address struct {
	gorm.Model
	Street string
	City string
	State string
	Zip string
	Latitude float64
	Longitude float64
	UserID uint
}

type Producer struct {
	gorm.Model
	Name string
	UserID uint
	Address uint 
}

type Product struct {
	gorm.Model
	Name string
	ProducerID uint
	Price float64

}