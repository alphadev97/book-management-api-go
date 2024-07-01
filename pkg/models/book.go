package models

import (
	"log"

	"github.com/alphadev97/bookstore-api-go/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	// db.Create(&b)
	// return b

	result := db.Create(&b)
	if result.Error != nil {
		log.Fatal("Error in create book")
		return nil
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// func GetBookById(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	db := db.Where("ID=?", Id).Find(&getBook)
// 	return &getBook, db
// }

func DeleteBook(ID int64) *Book {
	var book Book
	result := db.Where("ID = ?", ID).First(&book)
	if result.Error != nil {
		return nil
	}
	db.Delete(&book)
	return &book
}

func UpdateBook(book *Book) error {
	result := db.Save(book)
	return result.Error
}

func GetBookById(Id int64) (*Book, error) {
	var getBook Book
	result := db.Where("ID = ?", Id).First(&getBook)
	if result.Error != nil {
		return nil, result.Error
	}
	return &getBook, nil
}
