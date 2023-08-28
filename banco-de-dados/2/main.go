package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       string `gorm:"primaryKey`
	Name     string
	Products []Product
}

type Product struct {
	ID           string `gorm:"primaryKey`
	Name         string
	Price        float64
	CategoryID   string
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        string `gorm:"primaryKey`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{ID: uuid.New().String(), Name: "Eletrônicos"}
	db.Create(&category)

	db.Create(&Product{
		ID:         uuid.New().String(),
		Name:       "Geladeira",
		Price:      2700.00,
		CategoryID: category.ID,
	})

	db.Create(&SerialNumber{
		ID:        uuid.New().String(),
		Number:    uuid.New().String(),
		ProductID: 1,
	})

	pr := []Product{
		{ID: uuid.New().String(), Name: "Lixadeira", Price: 300.00},
		{ID: uuid.New().String(), Name: "Video-Game", Price: 5000.00},
		{ID: uuid.New().String(), Name: "Televisão", Price: 4000.00},
		{ID: uuid.New().String(), Name: "Notebook", Price: 2300.00},
	}

	db.Create(&pr)
	var product Product
	db.First(&product, "name = ?", "Lixadeira")
	fmt.Println(product)

	var produ []Product
	db.Find(&produ)
	for _, product := range produ {
		fmt.Println(product)
	}
	var p []Product
	db.Where("price > ?", 3000).Find(&p)
	for _, prod := range p {
		fmt.Println(prod)
	}
	var pro Product
	db.First(&pro)
	pro.Name = "Whatever"
	db.Save(&pro)

	var p2 Product
	db.First(&p2)
	fmt.Println(p2)
	db.Delete(&p2)

	var products []Product
	err = db.Preload("Products").Preload("Products.SerialNumber").Find(&products).Error
	if err != nil {
		panic(err)
	}
	for _, prod := range products {
		fmt.Println(prod.Name, prod.Category.Name, prod.SerialNumber.Number)
	}

}
