package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type City struct {
	ID          int
	Description string
	Status      string
	Code        string
}

func main() {
	fmt.Println("Hello")
	dsn := "Data Source=DESKTOP-DHPI0AF;Initial Catalog=MediSoft;Persist Security Info=True;User ID=sa;Password=Desarrollo"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexion exitosa..")

	var city City

	// db.Table("City").First(&city)
	// fmt.Println(city.Code, city.Description)
	db.Table("City").Where("Id = ?", 9).Find(&city)
	fmt.Println(city.Code, city.Description)

}
