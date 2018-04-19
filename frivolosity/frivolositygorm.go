package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ProductCategory struct {
	Id int64
	//CreatedAt time.Time
	//UpdatedAt time.Time

	Name string `gorm:"unique;not null"`
}

type Product struct {
	//gorm.Model
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	Code              string
	Price             int64
	ProductCategoryId int64
	ProductCategory   ProductCategory
}

//type User struct {
//	gorm.Model
//	Name string
//}
//
//// `Profile` belongs to `User`, `UserID` is the foreign key
//type Profile struct {
//	gorm.Model
//	UserID int
//	User   User `gorm:"foreignkey:UserID"`
//	Name   string
//}

func main() {
	db, err := gorm.Open("mysql", "root:@/gorm_testing?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.DropTableIfExists(&Product{}, &ProductCategory{})

	//db.AutoMigrate(&User{}, &Profile{})
	//db.AutoMigrate(&Profile{}, &User{})

	// Migrate the schema
	db.AutoMigrate(&Product{}, &ProductCategory{})

	// Create
	db.Save(&Product{Code: "L1212", Price: 1000, ProductCategory: ProductCategory{Name: "foo"}})
	db.Save(&Product{Code: "xyz", Price: 1000, ProductCategory: ProductCategory{Name: "foo"}})
	db.Create(&Product{Code: "xyz", Price: 1000, ProductCategory: ProductCategory{Name: "foo"}})

	// Read
	//var product Product
	//db.First(&product, 1)                   // find product with id 1
	//db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	//db.Delete(&product)
}
