package bookstore

import "errors"

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

func GetBook(catalog []Book, id int) (Book, error) {
	for _, b := range catalog {
		if b.ID == id {
			return b, nil
		}
	}
	return Book{}, errors.New("Book not found")
}
