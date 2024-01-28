package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Catalog map[int]Book
type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
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

func (c Catalog) GetAllBooks() []Book {
	var books []Book
	for _, b := range c {
		books = append(books, b)
	}
	return books
}

func (c Catalog) GetBook(id int) (Book, error) {
	b, ok := c[id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d not found in catalog.", id)
	}
	return b, nil
}

// func NetPriceCents(b Book) int {
// 	actual_price := b.PriceCents * (100 - b.DiscountPercent) / 100

// 	return actual_price
// }

// Method of the Book type
// that returns discounted price
func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(newPrice int) error {
	if newPrice < 0 {
		return fmt.Errorf("%d is invalid value for book price.", newPrice)
	}
	// automatic dereferencing
	b.PriceCents = newPrice
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func (b *Book) SetCategory(category Category) error {
	if validCategory[category] {
		b.category = category
		return nil
	} else {
		return fmt.Errorf("Unknown category %v", category)
	}
}
