package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	// var newBook Book = Book{
	// 	Title:  b.Title,
	// 	Author: b.Author,
	// 	Copies: b.Copies - 1,
	// }

	// return newBook
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	b, ok := catalog[id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d not found in catalog.", id)
	}
	return b, nil
}
