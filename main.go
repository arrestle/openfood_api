package main

import (
	"fmt"
	"os"
	"log"
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/arrestle/openfood_api/pkg/models"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"

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

	apiKey := os.Getenv("GEOCODETOKEN")
	if apiKey == "" {
		fmt.Println("API_KEY environment variable not set")
		return
	}

	fmt.Println("API Key:", apiKey)
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	address := "403 West Whitaker Mill Road, Raleigh, NC 27608"

	r := &maps.GeocodingRequest{
		Address: address}

	resp, err := c.Geocode(context.Background(), r)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(resp[0].Geometry.Location);

}