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