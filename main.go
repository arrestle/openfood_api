package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/arrestle/openfood_api/pkg/models"
)


func main() {
	db, err := gorm.Open("postgres", "user=postgres dbname=openfood password=postgres sslmode=disable")
	
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// user1 := User{Name: "Andrea Restle-Lay", Email: "arrestle@ncsu.edu"}
	// db.Create(&user1)
	// user1 = User{Name: "Mike Lay", Email: "cmikelay@yahoo.com"}
	// db.Create(&user1)


	user2 := db.Where("name = ?", "Andrea Restle-Lay").First(&models.User{}).Value.(*models.User)
	fmt.Printf("User 2 %+v\n\n",user2)

	var users []models.User
	db.Find(&users)

	for _, user := range users {
		fmt.Println(user)
		//db.Delete(&user)
	}

}